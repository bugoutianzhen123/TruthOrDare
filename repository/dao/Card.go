package dao

import "github.com/bugoutianzhen123/TruthOrDare/domain"

type Card interface {
	CreatedCard(c domain.Card) error
	DeletedCard(c domain.Card) error
	GetCard(mode, ty, style, num int8) *[]domain.Card
}

func (d *dao) CreatedCard(c domain.Card) error {
	return d.db.Create(&c).Error
}

func (d *dao) DeletedCard(c domain.Card) error {
	return d.db.Model(&c).Delete(&c).Error
}

func (d *dao) GetCard(mode, ty, style, num int8) *[]domain.Card {
	var cards []domain.Card
	d.db.Model(&cards).Where("mode = ? and type = ? and style = ?", mode, ty, style).Limit(int(num)).Find(&cards)
	return &cards
}
