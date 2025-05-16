package service

import (
	"github.com/bugoutianzhen123/TruthOrDare/domain"
	"time"
)

type UserService interface {
	CreateUser(user domain.User) error
}

func (s *ser) CreateUser(user domain.User) error {
	user.Created = time.Now()
	user.Updated = user.Created
	return s.r.CreateUser(user)
}
