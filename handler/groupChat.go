package handler

import (
	"github.com/bugoutianzhen123/TruthOrDare/service"

	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type GroupChatHandler interface {
	HandleWebSocket(c *gin.Context)
}

type groupChatHandler struct {
	clientManager *service.ClientManager
}

func NewGroupChatHandler(cm *service.ClientManager) *groupChatHandler {
	return &groupChatHandler{clientManager: cm}
}

func (h *hand) HandleWebSocket(c *gin.Context) {
	groupID, ok1 := parseId(c, "group_id")
	userID, ok2 := parseId(c, "user_id")
	if !ok1 || !ok2 {
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("WebSocket升级失败: %v", err)
		return
	}

	gm := h.ser.FindGroup(groupID)
	gm.AddClient(userID, conn)
}

func (h *hand) GetGroupChatHistory(c *gin.Context) {
	groupID, ok := parseId(c, "group_id")
	if !ok {
		return
	}

	limitStr := c.DefaultQuery("limit", "50")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit > 100 {
		limit = 50
	}

	messages, err := h.ser.GetGroupChatHistory(groupID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取历史失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": messages})
}
