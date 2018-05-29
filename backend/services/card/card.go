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
	cardModelSlice := make([]models.CardModel, 0)
	err = s.engine.Sql("select card.*, album.* from card, album where card.album_id = album.id and card.unique_id = ?", uniqueId).Find(&cardModelSlice)
	if err != nil {
		return card, err
	}
	if len(cardModelSlice) == 0 {
		err = errors.New("card data not exist with uniqueId: " + uniqueId)
		return
	}
	card = cardModelSlice[0].Card
	card.Album = cardModelSlice[0].Album
	return card, nil
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
