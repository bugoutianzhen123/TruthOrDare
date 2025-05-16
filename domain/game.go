package domain

import (
	"encoding/json"
	"time"
)

const (
	normal  = 1 << 0
	couple  = 1 << 1
	friends = 1 << 2
	party   = 1 << 3
)

const (
	mild     = 1 << 0
	exciting = 1 << 1
	funny    = 1 << 2
)

const (
	truth = 1
	dare  = 2
)

// cardInRepository
type Card struct {
	ID        uint64    `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Mode      int8      `json:"mode"`  // normal, couple, friends, party
	Style     int8      `json:"style"` // mild, exciting, funny
	Type      int8      `json:"type"`  // truth or dare
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type CardResponse struct {
	ID      uint64 `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Content string `json:"content"`
}

type GameRecord struct {
	CardID    int
	FlippedBy uint64
	Content   string
	FlipTime  time.Time
}

type GameMessage struct {
	Action  string          `json:"action"` // join_room, ready, start_game, flip_card 等
	UserID  uint64          `json:"user_id"`
	RoomID  uint64          `json:"room_id"`
	Payload json.RawMessage `json:"payload"` // 具体操作的内容
}

type EmptyPayload struct{}

type FlipCardPayload struct {
	CardIndex int `json:"card_index"`
}

type ChangeModPayload struct {
}
