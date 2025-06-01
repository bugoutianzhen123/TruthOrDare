package dao

import "github.com/bugoutianzhen123/TruthOrDare/domain"

type Card interface {
	CreatedCard(c domain.Card) error
	DeletedCard(c domain.Card) error
	GetCard(mode, ty, style, num int8) *[]domain.Card
	BatchCreatedCards(cards []domain.Card) error
	SaveGameHistory(h domain.GameHistory) error
	GetAllGameHistories() ([]domain.GameHistory, error)
	GetGameHistoriesByUserID(userID uint64) ([]domain.GameHistory, error)
}

func (d *dao) CreatedCard(c domain.Card) error {
	return d.db.Create(&c).Error
}

func (d *dao) DeletedCard(c domain.Card) error {
	return d.db.Model(&c).Delete(&c).Error
}

func (d *dao) GetCard(mode, ty, style, num int8) *[]domain.Card {
	var cards []domain.Card
	d.db.Model(&cards).
		Where("mode & ? > 0 and type & ? > 0 and style = ?", mode, ty, style).
		Order("RAND()").
		Limit(int(num)).
		Find(&cards)
	return &cards
}

func (d *dao) BatchCreatedCards(cards []domain.Card) error {
	return d.db.Create(&cards).Error
}

func (d *dao) SaveGameHistory(h domain.GameHistory) error {
	return d.db.Create(&h).Error
}

func (d *dao) GetAllGameHistories() ([]domain.GameHistory, error) {
	var histories []domain.GameHistory
	err := d.db.Order("created_at desc").Find(&histories).Error
	return histories, err
}

func (d *dao) GetGameHistoriesByUserID(userID uint64) ([]domain.GameHistory, error) {
	var histories []domain.GameHistory
	err := d.db.Where("user_id = ?", userID).Order("created_at desc").Find(&histories).Error
	return histories, err
}
