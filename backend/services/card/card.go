package services

import (
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
