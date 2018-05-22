package services

import (
	"errors"
	models "identify/backend/models/card"

	"github.com/go-redis/redis"
	"github.com/go-xorm/xorm"
)

type Card struct {
	engine *xorm.Engine
	cache  *redis.Client
}

func NewCard(engine *xorm.Engine, cache *redis.Client) *Card {
	return &Card{
		engine: engine,
		cache:  cache,
	}
}

func (s *Card) GetByUniqueId(uniqueId string) (card models.Card, err error) {
	cardSlice := make([]models.Card, 0)
	err = s.engine.Join("LEFT", "album", "album.id = card.album_id").Where("card.unique_id = ?", uniqueId).Find(&cardSlice)
	if err != nil {
		return card, err
	}
	if len(cardSlice) == 0 {
		err = errors.New("card data not exist with uniqueId: " + uniqueId)
		return
	}
	return cardSlice[0], nil
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
