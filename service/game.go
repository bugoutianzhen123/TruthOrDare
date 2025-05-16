package service

import (
	"github.com/bugoutianzhen123/TruthOrDare/domain"
)

type GameService interface {
	GetGameRoom(roomId, hostId uint64) *domain.GameRoom
	CreateCard(c domain.Card) error
	DeleteCard(id uint64) error
	GetCards(mode, ty, style int8) *[]domain.CardResponse
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

func (s *ser) GetCards(mode, ty, style int8) *[]domain.CardResponse {
	cards := s.r.GetCards(mode, ty, style)
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
