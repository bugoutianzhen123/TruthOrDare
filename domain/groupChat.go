package domain

import (
	"github.com/gorilla/websocket"
	"time"
)

// 用户连接
type Client struct {
	ID        string
	Conn      *websocket.Conn
	SendChan  chan []byte
	RoomID    string
	CreatedAt time.Time
}

// 房间实体
type Room struct {
	ID           string
	Clients      map[string]*Client // key: clientID
	CreatedAt    time.Time
	LastActiveAt time.Time
}

// 消息实体
type Message struct {
	RoomID    string
	Content   []byte
	SenderID  string
	Timestamp time.Time
}
