package handler

import (
	"github.com/bugoutianzhen123/TruthOrDare/domain"
	"github.com/gin-gonic/gin"
	"log"
)

type GameHandler interface {
	StartGame(c *gin.Context)
	CreateCard(c *gin.Context)
	RemoveCard(c *gin.Context)
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
	cards := h.ser.GetCards(mode, ty, style)
	c.JSON(200, gin.H{
		"data": cards,
	})
}

func (h *hand) CreateCard(c *gin.Context) {
	card := domain.Card{}
	err := c.ShouldBind(&card)
	if err != nil {
		c.JSON(400, gin.H{})
	}
	if err := h.ser.CreateCard(card); err != nil {
		c.JSON(400, gin.H{})
	}
	c.JSON(200, gin.H{})
}

func (h *hand) RemoveCard(c *gin.Context) {
	id, _ := parseId(c, "card_id")
	if err := h.ser.DeleteCard(id); err != nil {
		c.JSON(400, gin.H{})
	}
	c.JSON(200, gin.H{})
}
