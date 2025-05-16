package service

import (
	"encoding/json"
	"github.com/bugoutianzhen123/TruthOrDare/domain"
	"github.com/bugoutianzhen123/TruthOrDare/repository"

	"github.com/gorilla/websocket"
	"log"
	"sync"
	"time"
)

type User struct {
	Conn   *websocket.Conn
	UserId uint64
}

type GroupManager struct {
	GroupId     uint64
	users       map[*websocket.Conn]*User
	messageChan chan ChatMessage
	closeChan   chan struct{}
	mutex       sync.Mutex
	r           repository.GroupChat
}

type ChatMessage struct {
	GroupId  uint64    `json:"group_id"`
	UserId   uint64    `json:"user_id"`
	Content  string    `json:"content"`
	SendTime time.Time `json:"send_time"`
}

func NewGroupManager(repo repository.GroupChat) *GroupManager {
	gm := &GroupManager{
		users:       make(map[*websocket.Conn]*User),
		messageChan: make(chan ChatMessage, 100),
		closeChan:   make(chan struct{}),
		r:           repo,
	}
	go gm.messageDispatcher()
	return gm
}

// 核心消息调度器
func (gm *GroupManager) messageDispatcher() {

	for {
		select {
		case msg := <-gm.messageChan:
			gm.broadcast(msg)
		case <-gm.closeChan:
			return
		}
	}
}

// 添加客户端
func (gm *GroupManager) AddClient(userId uint64, conn *websocket.Conn) {
	gm.mutex.Lock()
	defer gm.mutex.Unlock()

	gm.users[conn] = &User{
		Conn:   conn,
		UserId: userId,
	}
	go gm.handleClient(userId, conn)
}

// 处理单个客户端消息
func (gm *GroupManager) handleClient(userId uint64, conn *websocket.Conn) {
	defer gm.RemoveClient(conn)

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err) {
				log.Printf("客户端异常断开: %v", err)
			}
			return
		}
		msg := ChatMessage{
			GroupId:  gm.GroupId,
			UserId:   userId,
			Content:  string(message),
			SendTime: time.Now(),
		}

		if err := gm.r.SaveGroupMessage(domain.GroupChatHistory{
			Created:  msg.SendTime,
			SenderId: msg.UserId,
			GroupId:  gm.GroupId,
			Text:     msg.Content,
		}); err != nil {
			log.Printf("消息保存失败: %v", err)
		}

		gm.messageChan <- msg
	}
}

// 移除客户端
func (gm *GroupManager) RemoveClient(conn *websocket.Conn) {
	gm.mutex.Lock()
	defer gm.mutex.Unlock()

	delete(gm.users, conn)
	conn.Close()

	if len(gm.users) == 0 {
		gm.closeChan <- struct{}{}
	}
}

// 广播消息
func (gm *GroupManager) broadcast(msg ChatMessage) {
	gm.mutex.Lock()
	defer gm.mutex.Unlock()

	jsonData, _ := json.Marshal(msg)
	for _, client := range gm.users {
		if err := client.Conn.WriteMessage(websocket.TextMessage, jsonData); err != nil {
			log.Printf("消息发送失败: %v", err)
			client.Conn.Close()
			delete(gm.users, client.Conn)
		}
	}
}
