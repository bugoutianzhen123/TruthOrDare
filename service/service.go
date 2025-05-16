package service

import (
	"github.com/bugoutianzhen123/TruthOrDare/domain"
	"github.com/bugoutianzhen123/TruthOrDare/repository"
)

type Service interface {
	UserService
	GroupService
	GroupChatService
	GameService
}

type ser struct {
	r  repository.Repository
	cm *ClientManager
	gm *domain.GameClientManager
}

func NewService(r repository.Repository) Service {
	cm := NewClientManager(r)
	gm := domain.NewGameClientManager()
	return &ser{r, cm, gm}
}
