package dao

import "github.com/bugoutianzhen123/TruthOrDare/domain"

type UserDao interface {
	CreateUser(u domain.User) error
}

func (d *dao) CreateUser(u domain.User) error {
	return d.db.Create(&u).Error
}
