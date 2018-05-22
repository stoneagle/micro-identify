package services

import (
	"encoding/json"
	models "identify/backend/models/card"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/go-xorm/xorm"
)

type Release struct {
	engine *xorm.Engine
	cache  *redis.Client
}

var (
	agentsPrefix = "Card-Album-Release-Agents-"
)

func NewRelease(engine *xorm.Engine, cache *redis.Client) *Release {
	return &Release{
		engine: engine,
		cache:  cache,
	}
}

func (s *Release) GetByAlbumId(albumId uint) (releases []models.Release, err error) {
	err = s.engine.Where("album_id = ?", albumId).Find(&releases)
	return releases, err
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

func (s *Release) SetCache(albumId uint, data map[string]int) (err error) {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	key := agentsPrefix + strconv.FormatUint(uint64(albumId), 10)
	err = s.cache.Set(key, string(b), 60*60*time.Second).Err()
	return err
}

func (s *Release) GetCache(albumId uint) (agentsMap map[string]int, err error) {
	key := agentsPrefix + strconv.FormatUint(uint64(albumId), 10)
	val, err := s.cache.Get(key).Result()
	if err == redis.Nil {
		return agentsMap, nil
	} else if err != nil {
		return
	}
	err = json.Unmarshal([]byte(val), &agentsMap)
	return
}

func (s *Release) DelCache(albumId string) (err error) {
	key := agentsPrefix + albumId
	_, err = s.cache.Del(key).Result()
	return
}
