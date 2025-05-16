package service

import "github.com/bugoutianzhen123/TruthOrDare/domain"

type GroupService interface {
	CreateGroup(group domain.Group) error
}

func (s *ser) CreateGroup(group domain.Group) error {
	return s.r.CreateGroup(group)
}
