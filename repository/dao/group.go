package dao

import "github.com/bugoutianzhen123/TruthOrDare/domain"

type GroupDao interface {
	CreateGroup(group domain.Group) error
}

func (d *dao) CreateGroup(group domain.Group) error {
	return d.db.Create(&group).Error
}
