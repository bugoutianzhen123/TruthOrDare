package repository

import (
	"time"

	"github.com/bugoutianzhen123/TruthOrDare/domain"
)

type GameRepository interface {
	CreatedCard(c domain.Card) error
	DeletedCard(c domain.Card) error
	GetCards(mode, ty, style, num int8) *[]domain.Card
	BatchCreatedCards(cards []domain.Card) error
	SaveGameHistory(h domain.GameHistory) error
	GetAllGameHistories() ([]domain.GameHistory, error)
	GetGameHistoriesByUserID(userID uint64) ([]domain.GameHistory, error)
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

func (r *repo) BatchCreatedCards(cards []domain.Card) error {
	for i := range cards {
		cards[i].CreatedAt = time.Now()
	}
	return r.dao.BatchCreatedCards(cards)
}

func (r *repo) SaveGameHistory(h domain.GameHistory) error {
	h.CreatedAt = time.Now()
	return r.dao.SaveGameHistory(h)
}

func (r *repo) GetAllGameHistories() ([]domain.GameHistory, error) {
	return r.dao.GetAllGameHistories()
}

func (r *repo) GetGameHistoriesByUserID(userID uint64) ([]domain.GameHistory, error) {
	return r.dao.GetGameHistoriesByUserID(userID)
}
