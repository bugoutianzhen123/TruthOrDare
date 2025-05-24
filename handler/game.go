package handler

import (
	"fmt"
	"log"

	"github.com/bugoutianzhen123/TruthOrDare/domain"
	"github.com/gin-gonic/gin"
)

type GameHandler interface {
	StartGame(c *gin.Context)
	CreateCard(c *gin.Context)
	RemoveCard(c *gin.Context)
	BatchCreateCards(c *gin.Context)
	SaveGameHistory(c *gin.Context)
	GetAllGameHistories(c *gin.Context)
	GetGameHistoriesByUserID(c *gin.Context)
}

// user id     room id
func (h *hand) GameWebSocket(c *gin.Context) {
	userId, ok1 := parseId(c, "user_id")
	roomId, ok2 := parseId(c, "room_id")
	if !ok1 || !ok2 {
		return
	}

	userName := c.Query("username")

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("WebSocket升级失败: %v", err)
		return
	}

	gm := h.ser.GetGameRoom(userId, roomId)
	gm.AddUser(&domain.Player{
		Conn:     conn,
		UserID:   userId,
		UserName: userName,
		IsReady:  userId == roomId,
	})

}

func (h *hand) StartGame(c *gin.Context) {
	mode, _ := parseInt8(c, "mode")
	ty, _ := parseInt8(c, "type")
	style, _ := parseInt8(c, "style")
	num, _ := parseInt8(c, "num")
	cards := h.ser.GetCards(mode, ty, style, num)
	c.JSON(200, gin.H{
		"data": cards,
	})
}

func (h *hand) CreateCard(c *gin.Context) {
	card := domain.Card{}
	err := c.ShouldBind(&card)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
	}
	if err := h.ser.CreateCard(card); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
	}
	c.JSON(200, gin.H{})
}

func (h *hand) RemoveCard(c *gin.Context) {
	id, _ := parseId(c, "card_id")
	if err := h.ser.DeleteCard(id); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
	}
	c.JSON(200, gin.H{})
}

func (h *hand) BatchCreateCards(c *gin.Context) {
	var cards []domain.Card
	if err := c.ShouldBindJSON(&cards); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	if err := h.ser.BatchCreateCards(cards); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{})
}

func (h *hand) SaveGameHistory(c *gin.Context) {
	var req struct {
		Mode       int8    `json:"mode"`
		Type       int8    `json:"type"`
		Style      int8    `json:"style"`
		CardNumber int     `json:"card_number"`
		CardIDs    []uint64 `json:"card_ids"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	// 将 card_ids 转为逗号分隔字符串
	ids := ""
	for i, id := range req.CardIDs {
		if i > 0 {
			ids += ","
		}
		ids += fmt.Sprintf("%d", id)
	}
	history := domain.GameHistory{
		Mode: req.Mode,
		Type: req.Type,
		Style: req.Style,
		CardNumber: req.CardNumber,
		CardIDs: ids,
	}
	if err := h.ser.SaveGameHistory(history); err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "ok"})
}

func (h *hand) GetAllGameHistories(c *gin.Context) {
	histories, err := h.ser.GetAllGameHistories()
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": histories})
}

func (h *hand) GetGameHistoriesByUserID(c *gin.Context) {
	userID, ok := parseId(c, "user_id")
	if !ok {
		return
	}
	histories, err := h.ser.GetGameHistoriesByUserID(userID)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": histories})
}
