package repository

import "github.com/bugoutianzhen123/TruthOrDare/domain"

type Group interface {
	CreateGroup(group domain.Group) error
}

func (r *repo) CreateGroup(group domain.Group) error {
	return r.dao.CreateGroup(group)
}
