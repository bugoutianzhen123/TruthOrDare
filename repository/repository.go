package repository

import "github.com/bugoutianzhen123/TruthOrDare/repository/dao"

type Repository interface {
	User
	Group
	GroupChat
	GameRepository
}

type repo struct {
	dao dao.Dao
}

func NewRepository(dao dao.Dao) Repository {
	return &repo{dao}
}
