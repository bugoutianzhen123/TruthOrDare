package service

import (
	"github.com/bugoutianzhen123/TruthOrDare/domain"
)

type GameService interface {
	GetGameRoom(roomId, hostId uint64) *domain.GameRoom
	CreateCard(c domain.Card) error
	DeleteCard(id uint64) error
	GetCards(mode, ty, style, num int8) *[]domain.CardResponse
	BatchCreateCards(cards []domain.Card) error
	SaveGameHistory(h domain.GameHistory) error
	GetAllGameHistories() ([]domain.GameHistory, error)
	GetGameHistoriesByUserID(userID uint64) ([]domain.GameHistory, error)
}

func (s *ser) GetGameRoom(roomId, hostId uint64) *domain.GameRoom {
	return s.gm.GetRoom(roomId, hostId)
}

func (s *ser) CreateCard(c domain.Card) error {
	return s.r.CreatedCard(c)
}

func (s *ser) DeleteCard(id uint64) error {
	return s.r.DeletedCard(domain.Card{ID: id})
}

func (s *ser) GetCards(mode, ty, style, num int8) *[]domain.CardResponse {
	cards := s.r.GetCards(mode, ty, style, num)
	cardsResponse := []domain.CardResponse{}
	for _, c := range *cards {
		c := domain.CardResponse{
			ID:      c.ID,
			Content: c.Content,
		}
		cardsResponse = append(cardsResponse, c)
	}
	return &cardsResponse
}

func (s *ser) BatchCreateCards(cards []domain.Card) error {
	return s.r.BatchCreatedCards(cards)
}

func (s *ser) SaveGameHistory(h domain.GameHistory) error {
	return s.r.SaveGameHistory(h)
}

func (s *ser) GetAllGameHistories() ([]domain.GameHistory, error) {
	return s.r.GetAllGameHistories()
}

func (s *ser) GetGameHistoriesByUserID(userID uint64) ([]domain.GameHistory, error) {
	return s.r.GetGameHistoriesByUserID(userID)
}
