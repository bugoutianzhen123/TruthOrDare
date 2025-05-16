package repository

// 房间持久化接口
type GroupChat interface {
	SaveGroupMessage(message domain.GroupChatHistory) error
	GetGroupChatHistory(groupId uint64) ([]domain.GroupChatHistory, error)
}

func (r *repo) SaveGroupMessage(message domain.GroupChatHistory) error {
	return r.dao.CreateGroupChatHistory(message)
}

func (r *repo) GetGroupChatHistory(groupId uint64) ([]domain.GroupChatHistory, error) {
	return r.dao.GetGroupHistory(groupId)
}
