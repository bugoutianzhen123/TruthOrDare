package dao

import (
	"github.com/bugoutianzhen123/TruthOrDare/domain"
	"github.com/bugoutianzhen123/TruthOrDare/pkg/logger"
	"gorm.io/gorm"
)

type Dao interface {
	UserDao
	GroupDao
	groupChatDao
	Card
}

type dao struct {
	db *gorm.DB
	l  logger.Logger
}

func NewDao(db *gorm.DB, l logger.Logger) Dao {
	return &dao{db: db, l: l}
}

func InitTables(db *gorm.DB) error {
	db.AutoMigrate(&domain.UserPermission{})
	db.AutoMigrate(&domain.User{})

	db.AutoMigrate(&domain.Friend{})
	db.AutoMigrate(&domain.FriendChatHistory{})
	db.AutoMigrate(&domain.FriendApplication{})

	db.AutoMigrate(&domain.RequestStatus{})

	db.AutoMigrate(&domain.Group{})
	db.AutoMigrate(&domain.GroupIdentity{})
	db.AutoMigrate(&domain.GroupChatHistory{})
	db.AutoMigrate(&domain.GroupMember{})

	db.AutoMigrate(&domain.Card{})

	return db.Error
}
