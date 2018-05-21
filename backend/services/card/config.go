package services

import (
	models "identify/backend/models/card"

	"github.com/go-xorm/xorm"
)

type Config struct {
	engine *xorm.Engine
}

func NewConfig(engine *xorm.Engine) *Config {
	return &Config{
		engine: engine,
	}
}

func (s *Config) Add(m *models.Config) (err error) {
	_, err = s.engine.Insert(m)
	return err
}

func (s *Config) Del(m *models.Config, hard bool) (err error) {
	if hard {
		_, err = s.engine.Unscoped().Delete(m)
	} else {
		_, err = s.engine.Delete(m)
	}
	return err
}
