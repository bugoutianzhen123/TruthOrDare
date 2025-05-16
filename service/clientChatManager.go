package service

import (
	"github.com/bugoutianzhen123/TruthOrDare/repository"
	"log"
	"sync"
)

type ClientManager struct {
	r      repository.GroupChat
	groups map[uint64]*GroupManager
	mutex  sync.Mutex
}

func NewClientManager(repo repository.GroupChat) *ClientManager {
	return &ClientManager{
		groups: make(map[uint64]*GroupManager),
		r:      repo,
	}
}

func (cm *ClientManager) GetGroup(groupID uint64) *GroupManager {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	if gm, exists := cm.groups[groupID]; exists {
		return gm
	}

	gm := NewGroupManager(cm.r)
	cm.groups[groupID] = gm
	log.Printf("创建新分组: %d", groupID)
	return gm
}

func (cm *ClientManager) RemoveGroup(groupID uint64) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	if gm, exists := cm.groups[groupID]; exists {
		gm.closeChan <- struct{}{}
		delete(cm.groups, groupID)
		log.Printf("移除分组: %d", groupID)
	}
}
