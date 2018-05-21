package services

import (
	models "identify/backend/models/card"

	"github.com/go-xorm/xorm"
)

type Album struct {
	engine *xorm.Engine
}

func NewAlbum(engine *xorm.Engine) *Album {
	return &Album{
		engine: engine,
	}
}

func (s *Album) Add(a *models.Album) (err error) {
	_, err = s.engine.Insert(a)
	return err
}

func (s *Album) Del(a *models.Album, hard bool) (err error) {
	if hard {
		_, err = s.engine.Unscoped().Delete(a)
	} else {
		_, err = s.engine.Delete(a)
	}
	return err
}
