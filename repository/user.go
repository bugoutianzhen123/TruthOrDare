package repository

import "github.com/bugoutianzhen123/TruthOrDare/domain"

type User interface {
	CreateUser(user domain.User) error
}

func (r *repo) CreateUser(user domain.User) error {
	return r.dao.CreateUser(user)
}
