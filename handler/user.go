package handler

import (
	"net/http"

	"github.com/bugoutianzhen123/TruthOrDare/domain"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler interface {
	CreateUser(c *gin.Context)
	Login(c *gin.Context)
}

func (h *hand) CreateUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.ser.CreateUser(user); err != nil {
		if err == gorm.ErrDuplicatedKey {
			c.JSON(http.StatusBadRequest, gin.H{"error": "该邮箱已被注册"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"user": user})
}

func (h *hand) Login(c *gin.Context) {
	var req struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := h.ser.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "邮箱或密码错误"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}
