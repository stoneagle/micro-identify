package services

import (
	"errors"
	models "identify/backend/models/card"

	"github.com/go-redis/redis"
	"github.com/go-xorm/xorm"
)

type Album struct {
	engine *xorm.Engine
	cache  *redis.Client
}

func NewAlbum(engine *xorm.Engine, cache *redis.Client) *Album {
	return &Album{
		engine: engine,
	}
}

func (s *Album) Get(a *models.Album) (err error) {
	has, err := s.engine.Get(a)
	if err != nil {
		return err
	}
	if !has {
		err = errors.New("album data not exist")
	}
	return err
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
