package controllers

import (
	"identify/backend/common"

	"github.com/go-redis/redis"
	"github.com/go-xorm/xorm"
)

type Base struct {
	Config common.Conf
	Type   common.ProjectType
	Engine *xorm.Engine
	Cache  *redis.Client
}

func (b *Base) Prepare(ptype common.ProjectType) {
	b.Config = *common.GetConfig()
	b.Cache = common.GetRedis()
	b.Type = ptype
	switch b.Type {
	case common.ProjectCard:
		b.Engine = common.GetEngine(b.Config.Card.Database.Name)
	default:
		panic("project type relate database engine not exist")
	}
}
