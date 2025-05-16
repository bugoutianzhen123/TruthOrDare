package domain

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"sync"
	"time"
)

type GameClientManager struct {
	Rooms map[uint64]*GameRoom
	mutex sync.Mutex
}

type GameRoom struct {
	ID          uint64
	HostId      uint64
	Clients     map[uint64]*Player
	GameMode    int
	maxPlayers  int
	messageChan chan GameMessage
	Mutex       sync.Mutex
	closeChan   chan struct{}
}

type Player struct {
	Conn     *websocket.Conn
	UserID   uint64
	UserName string
	IsReady  bool // 是否点击准备
}

type BroadcastMessage struct {
	Type    string      `json:"type"`    // 类型，例如 "player_ready", "card_flipped"
	Payload interface{} `json:"payload"` // 数据载荷
}

func NewGameRoom(id, hostId uint64) *GameRoom {
	return &GameRoom{
		ID:          id,
		HostId:      hostId,
		Clients:     make(map[uint64]*Player),
		closeChan:   make(chan struct{}),
		messageChan: make(chan GameMessage),
		maxPlayers:  6,
	}
}

func (gm *GameRoom) AddUser(user *Player) error {
	gm.Mutex.Lock()
	defer gm.Mutex.Unlock()

	if len(gm.Clients) >= gm.maxPlayers {
		return fmt.Errorf("\"房间已满\"")
	}

	gm.Clients[user.UserID] = user

	go gm.HandleClient(user.Conn)
	gm.BroadcastRoomState()
	return nil
}

func (gm *GameRoom) RemoveUser(user *Player) {
	gm.Mutex.Lock()
	defer gm.Mutex.Unlock()

	delete(gm.Clients, user.UserID)
	if len(gm.Clients) == 0 {
		close(gm.messageChan)
	}

	gm.BroadcastRoomState()
}

func (gm *GameRoom) Start() {
	for {
		select {
		case msg := <-gm.messageChan:
			gm.HandleMessage(msg)
		case <-gm.closeChan:
			log.Println("GameRoom shutting down")
			return
		}
	}
}

func (gm *GameRoom) HandleClient(conn *websocket.Conn) {
	defer conn.Close()

	for {
		_, data, err := conn.ReadMessage()
		if err != nil {
			log.Printf("用户连接中断: %v", err)
			return
		}

		msg := GameMessage{}
		if err := json.Unmarshal(data, &msg); err != nil {
			log.Printf("GameMessage 解析失败: %v", err)
			continue
		}
		gm.messageChan <- msg
	}
}

func (gm *GameRoom) HandleMessage(msg GameMessage) {
	switch msg.Action {
	case "ready":
		gm.setReady(msg.UserID, true)
	case "cancel_ready":
		gm.setReady(msg.UserID, false)

	case "start_game":
		if msg.UserID != gm.HostId {
			//不是房主
			return
		}
		if !gm.AllReady() {
			gm.Broadcast(BroadcastMessage{
				Type: "not_all_ready",
				Payload: map[string]interface{}{
					"message": "不是所有玩家都已准备",
				},
			})
			return
		}
		gm.StartGame()

	case "flip_card":
		var payload FlipCardPayload
		if err := json.Unmarshal(msg.Payload, &payload); err != nil {
			log.Printf("flip_card payload 解析失败: %v", err)
			return
		}
		gm.handleFlipCard(msg.UserID, payload.CardIndex)
	default:
		log.Printf("未知操作: %s", msg.Action)
	}
}

func (gm *GameRoom) setReady(userID uint64, isReady bool) {
	gm.Mutex.Lock()
	defer gm.Mutex.Unlock()
	if player, ok := gm.Clients[userID]; ok {
		player.IsReady = isReady
		log.Printf("用户 %d 设置准备状态: %v", userID, isReady)
	}

	gm.BroadcastRoomState()
}

func (gm *GameRoom) AllReady() bool {
	gm.Mutex.Lock()
	defer gm.Mutex.Unlock()

	for _, player := range gm.Clients {
		if !player.IsReady {
			return false
		}
	}
	return true
}

func (gm *GameRoom) StartGame() {
	// TODO: 初始化游戏状态，如卡牌、分数等（根据具体游戏逻辑）

	// 广播游戏开始
	gm.Broadcast(BroadcastMessage{
		Type: "game_started",
		Payload: map[string]interface{}{
			"room_id":    gm.ID,
			"start_time": time.Now(),
		},
	})
}

func (gm *GameRoom) handleFlipCard(userID uint64, cardIndex int) {
	log.Printf("用户 %d 翻了第 %d 张牌", userID, cardIndex)
	// TODO: 翻牌逻辑
}

func (gm *GameRoom) BroadcastRoomState() {
	players := []map[string]interface{}{}
	for _, p := range gm.Clients {
		players = append(players, map[string]interface{}{
			"user_id":   p.UserID,
			"user_name": p.UserName,
			"is_ready":  p.IsReady,
		})
	}

	gm.Broadcast(BroadcastMessage{
		Type: "room_state",
		Payload: map[string]interface{}{
			"room_id": gm.ID,
			"players": players,
		},
	})
}

func (gm *GameRoom) Broadcast(msg BroadcastMessage) {
	gm.Mutex.Lock()
	defer gm.Mutex.Unlock()

	for _, player := range gm.Clients {
		if player.Conn == nil {
			continue
		}
		err := player.Conn.WriteJSON(msg)
		if err != nil {
			log.Printf("发送广播失败，用户 %d: %v", player.UserID, err)
		}
	}
}

func NewGameClientManager() *GameClientManager {
	return &GameClientManager{
		Rooms: make(map[uint64]*GameRoom),
	}
}

func (c *GameClientManager) GetRoom(roomID, hostID uint64) *GameRoom {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	room, ok := c.Rooms[roomID]
	if !ok {
		room = NewGameRoom(roomID, hostID)
		c.Rooms[roomID] = room
		go room.Start()
	}
	return room
}

func (c *GameClientManager) RemoveRoom(roomID uint64) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if room, ok := c.Rooms[roomID]; ok {
		close(room.closeChan)
		delete(c.Rooms, roomID)
		log.Printf("房间 %d 已移除", roomID)
	}
}
