package service

import (
	"time"

	"github.com/bugoutianzhen123/TruthOrDare/domain"
)

type UserService interface {
	CreateUser(user domain.User) error
	Login(email, password string) (*domain.User, error)
}

func (s *ser) CreateUser(user domain.User) error {
	user.Created = time.Now()
	user.Updated = user.Created
	return s.r.CreateUser(user)
}

func (s *ser) Login(email, password string) (*domain.User, error) {
	return s.r.FindUserByEmailAndPassword(email, password)
}
