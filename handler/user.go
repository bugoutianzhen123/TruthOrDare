package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler interface {
	CreateUser(c *gin.Context)
}

func (h *hand) CreateUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := h.ser.CreateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusCreated, gin.H{"user": user})
}
