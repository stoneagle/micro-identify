package services

import (
	"encoding/json"
	models "identify/backend/models/card"
	"time"

	"github.com/go-redis/redis"
	"github.com/go-xorm/xorm"
)

type Config struct {
	engine *xorm.Engine
	cache  *redis.Client
}

var (
	configsPrefix = "Card-Unique-Configs-"
)

func NewConfig(engine *xorm.Engine, cache *redis.Client) *Config {
	return &Config{
		engine: engine,
		cache:  cache,
	}
}

func (s *Config) GetByCardId(cardId uint) (configs []models.Config, err error) {
	configsModelSlice := make([]models.ConfigModel, 0)
	err = s.engine.Sql("select config.*, message.* from config, message where config.id = message.config_id").Where("config.card_id = ?", cardId).Find(&configsModelSlice)
	if err != nil {
		return configs, err
	}

	configsMap := make(map[uint]models.Config)
	for _, one := range configsModelSlice {
		if target, ok := configsMap[one.Config.Id]; !ok {
			config := one.Config
			config.Messages = make([]models.Message, 0)
			msg := one.Message
			if msg.Id != 0 {
				config.Messages = append(config.Messages, msg)
			}
			configsMap[one.Config.Id] = config
		} else if one.Message.Id != 0 {
			msg := one.Message
			target.Messages = append(target.Messages, msg)
			configsMap[one.Config.Id] = target
		}
	}

	configs = make([]models.Config, 0)
	for _, one := range configsMap {
		configs = append(configs, one)
	}
	return configs, err
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

func (s *Config) SetCache(uniqueId string, data []models.Config) (err error) {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	key := configsPrefix + uniqueId
	err = s.cache.Set(key, string(b), 60*60*time.Second).Err()
	return err
}

func (s *Config) GetCache(uniqueId string) (configs []models.Config, err error) {
	key := configsPrefix + uniqueId
	val, err := s.cache.Get(key).Result()
	if err == redis.Nil {
		return configs, nil
	} else if err != nil {
		return
	}
	err = json.Unmarshal([]byte(val), &configs)
	return
}

func (s *Config) DelCache(uniqueId string) (err error) {
	key := configsPrefix + uniqueId
	_, err = s.cache.Del(key).Result()
	return
}
