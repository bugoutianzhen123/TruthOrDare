package repository

import "github.com/bugoutianzhen123/TruthOrDare/domain"

type User interface {
	CreateUser(user domain.User) error
	FindUserByEmailAndPassword(email, password string) (*domain.User, error)
}

func (r *repo) CreateUser(user domain.User) error {
	user.Permission.Id = 1
	return r.dao.CreateUser(user)
}

func (r *repo) FindUserByEmailAndPassword(email, password string) (*domain.User, error) {
	return r.dao.FindUserByEmailAndPassword(email, password)
}
