package controllers

import (
	"identify/backend/common"
	"net/http"

	"github.com/gin-gonic/gin"
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

func (b *Base) Redirect(ctx *gin.Context, uri string) {
	ctx.Redirect(http.StatusFound, uri)
}

func (b *Base) Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": common.ErrorOk,
		"data": data,
		"desc": "",
	})
}

func (b *Base) ErrorBusiness(ctx *gin.Context, code common.ErrorCode, desc string, err error) {
	if err != nil {
		desc += ":" + err.Error()
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": struct{}{},
		"desc": desc,
	})
}

func (b *Base) ErrorServer(ctx *gin.Context, desc string) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": common.ErrorServer,
		"data": struct{}{},
		"desc": desc,
	})
}
