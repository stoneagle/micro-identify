package services

import (
	models "identify/backend/models/card"

	"github.com/go-xorm/xorm"
)

type Release struct {
	engine *xorm.Engine
}

func NewRelease(engine *xorm.Engine) *Release {
	return &Release{
		engine: engine,
	}
}

func (s *Release) Add(m *models.Release) (err error) {
	_, err = s.engine.Insert(m)
	return err
}

func (s *Release) Del(m *models.Release, hard bool) (err error) {
	if hard {
		_, err = s.engine.Unscoped().Delete(m)
	} else {
		_, err = s.engine.Delete(m)
	}
	return err
}
