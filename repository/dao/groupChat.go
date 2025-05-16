package dao

import "github.com/bugoutianzhen123/TruthOrDare/domain"

type groupChatDao interface {
	CreateGroupChatHistory(message domain.GroupChatHistory) error
	GetGroupHistory(groupId uint64) ([]domain.GroupChatHistory, error)
}

func (d *dao) CreateGroupChatHistory(message domain.GroupChatHistory) error {
	return d.db.Create(&message).Error
}

func (d *dao) GetGroupHistory(groupId uint64) ([]domain.GroupChatHistory, error) {
	var historys []domain.GroupChatHistory
	err := d.db.Where("group_id = ?", groupId).Order("created desc").Find(&historys).Error
	return historys, err
}
