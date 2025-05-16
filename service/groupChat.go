package service

import "github.com/bugoutianzhen123/TruthOrDare/domain"

type GroupChatService interface {
	FindGroup(groupID uint64) *GroupManager
	RemoveClient(groupID uint64) error
	GetGroupChatHistory(groupID uint64) ([]domain.GroupChatHistory, error)
}

func (s *ser) FindGroup(groupID uint64) *GroupManager {
	return s.cm.GetGroup(groupID)
}

func (s *ser) RemoveClient(groupID uint64) error {
	s.cm.RemoveGroup(groupID)
	return nil
}

func (s *ser) GetGroupChatHistory(groupID uint64) ([]domain.GroupChatHistory, error) {
	return s.r.GetGroupChatHistory(groupID)
}
