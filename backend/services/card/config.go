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

func (s *Config) GetByCardId(cardId uint) (configs []models.Config, err error) {
	configsSlice := make([]models.Config, 0)
	err = s.engine.Join("LEFT", "message", "config.id = message.config_id").Where("config.card_id = ?", cardId).Find(&configsSlice)
	if err != nil {
		return configs, err
	}

	configsMap := make(map[uint]models.Config)
	for _, one := range configsSlice {
		if target, ok := configsMap[one.Id]; !ok {
			one.Messages = make([]models.Message, 0)
			msg := one.Message
			if msg.Id != 0 {
				one.Messages = append(one.Messages, msg)
			}
			configsMap[one.Id] = one
		} else if one.Message.Id != 0 {
			msg := one.Message
			target.Messages = append(target.Messages, msg)
			configsMap[one.Id] = target
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
