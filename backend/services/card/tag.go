package services

import (
	models "identify/backend/models/card"

	"github.com/go-xorm/xorm"
)

type Tag struct {
	engine *xorm.Engine
}

func NewTag(engine *xorm.Engine) *Tag {
	return &Tag{
		engine: engine,
	}
}

func (s *Tag) Add(m *models.Tag) (err error) {
	_, err = s.engine.Insert(m)
	return err
}

func (s *Tag) Del(m *models.Tag, hard bool) (err error) {
	if hard {
		_, err = s.engine.Unscoped().Delete(m)
	} else {
		_, err = s.engine.Delete(m)
	}
	return err
}
