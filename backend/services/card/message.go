package services

import (
	models "identify/backend/models/card"

	"github.com/go-xorm/xorm"
)

type Message struct {
	engine *xorm.Engine
}

func NewMessage(engine *xorm.Engine) *Message {
	return &Message{
		engine: engine,
	}
}

func (s *Message) Add(m *models.Message) (err error) {
	_, err = s.engine.Insert(m)
	return err
}

func (s *Message) Del(m *models.Message, hard bool) (err error) {
	if hard {
		_, err = s.engine.Unscoped().Delete(m)
	} else {
		_, err = s.engine.Delete(m)
	}
	return err
}
