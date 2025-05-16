package handler

import (
	"fmt"
	"github.com/bugoutianzhen123/TruthOrDare/pkg/logger"
	"github.com/bugoutianzhen123/TruthOrDare/service"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
)

type Handler interface {
	UserHandler
	GroupHandler
	GroupChatHandler
	GameHandler
}

type hand struct {
	ser service.Service
	l   logger.Logger
}

func NewHandler(ser service.Service, l logger.Logger) *hand {
	return &hand{ser, l}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func parseId(c *gin.Context, parse string) (uint64, bool) {
	str := c.Query(parse)
	if str == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "缺少 userid 参数",
		})
		return 0, false
	}

	id, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("无效的 userid: %s", str),
		})
		return 0, false
	}
	return id, true
}

func parseInt8(c *gin.Context, parse string) (int8, bool) {
	num := c.Query(parse)
	if num == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "缺少 " + parse + " 参数",
		})
		return 0, false
	}

	n, err := strconv.Atoi(num)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("无效的 roomid: %s", num),
		})
		return 0, false
	}
	return int8(n), true
}
