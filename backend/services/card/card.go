package services

import (
	"errors"
	models "identify/backend/models/card"

	"github.com/go-xorm/xorm"
)

type Card struct {
	engine *xorm.Engine
}

func NewCard(engine *xorm.Engine) *Card {
	return &Card{
		engine: engine,
	}
}

func (s *Card) GetByUniqueId(uniqueId string) (card models.Card, err error) {
	has, err := s.engine.Where("unique_id=?", uniqueId).Get(&card)
	if err != nil {
		return card, err
	}
	if !has {
		err = errors.New("card data not exist with uniqueId: " + uniqueId)
	}
	return card, err
}

func (s *Card) Add(m *models.Card) (err error) {
	_, err = s.engine.Insert(m)
	return err
}

func (s *Card) Del(m *models.Card, hard bool) (err error) {
	if hard {
		_, err = s.engine.Unscoped().Delete(m)
	} else {
		_, err = s.engine.Delete(m)
	}
	return err
}
