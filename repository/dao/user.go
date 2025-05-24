package dao

import (
	"github.com/bugoutianzhen123/TruthOrDare/domain"
	"gorm.io/gorm"
)

type UserDao interface {
	CreateUser(u domain.User) error
	FindUserByEmailAndPassword(email, password string) (*domain.User, error)
}

func (d *dao) CreateUser(u domain.User) error {
	var count int64
	err := d.db.Model(&domain.User{}).Where("email = ?", u.Email).Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return gorm.ErrDuplicatedKey
	}
	return d.db.Create(&u).Error
}

func (d *dao) FindUserByEmailAndPassword(email, password string) (*domain.User, error) {
	var user domain.User
	err := d.db.Where("email = ? AND password = ?", email, password).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
