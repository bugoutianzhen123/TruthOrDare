package repository

import (
	"github.com/bugoutianzhen123/TruthOrDare/domain"
	"time"
)

type GameRepository interface {
	CreatedCard(c domain.Card) error
	DeletedCard(c domain.Card) error
	GetCards(mode, ty, style, num int8) *[]domain.Card
}

func (r *repo) CreatedCard(c domain.Card) error {
	c.CreatedAt = time.Now()
	return r.dao.CreatedCard(c)
}

func (r *repo) DeletedCard(c domain.Card) error {
	return r.dao.DeletedCard(c)
}

func (r *repo) GetCards(mode, ty, style, num int8) *[]domain.Card {
	return r.dao.GetCard(mode, ty, style, num)
}
