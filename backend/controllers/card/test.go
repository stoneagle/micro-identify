package card

import (
	"identify/backend/common"
	"identify/backend/controllers"
	models "identify/backend/models/card"
	services "identify/backend/services/card"

	"github.com/gin-gonic/gin"
)

type Test struct {
	controllers.Base
}

func NewTest() *Test {
	test := &Test{}
	test.Prepare(common.ProjectCard)
	return test
}

func (t *Test) Router(router *gin.RouterGroup) {
	tests := router.Group("")
	tests.PUT("album", t.AlbumAdd)
	tests.DELETE("album", t.AlbumDel)
	tests.PUT("release", t.ReleaseAdd)
	tests.DELETE("release", t.ReleaseDel)
	tests.PUT("", t.CardAdd)
	tests.DELETE("", t.CardDel)
	tests.PUT("tag", t.TagAdd)
	tests.DELETE("tag", t.TagDel)
	tests.PUT("config", t.ConfigAdd)
	tests.DELETE("config", t.ConfigDel)
	tests.PUT("message", t.MessageAdd)
	tests.DELETE("message", t.MessageDel)
}

func (c *Test) AlbumAdd(ctx *gin.Context) {
	var album models.Album
	err := ctx.ShouldBindJSON(&album)
	if err != nil {
		common.ResponseErrorBusiness(ctx, common.ErrorParams, "album params bind failed", err)
		return
	}

	svc := services.NewAlbum(c.Engine, c.Cache)
	err = svc.Add(&album)
	if err != nil {
		common.ResponseErrorBusiness(ctx, common.ErrorMysql, "album save failed", err)
		return
	}

	common.ResponseSuccess(ctx, album)
}

func (c *Test) AlbumDel(ctx *gin.Context) {
	var album models.Album
	err := ctx.ShouldBindJSON(&album)
	if err != nil {
		common.ResponseErrorBusiness(ctx, common.ErrorParams, "album params bind failed", err)
		return
	}

	svc := services.NewAlbum(c.Engine, c.Cache)
	err = svc.Del(&album, true)
	if err != nil {
		common.ResponseErrorBusiness(ctx, common.ErrorMysql, "album delete failed", err)
		return
	}

	common.ResponseSuccess(ctx, album.Name)
}

func (c *Test) ReleaseAdd(ctx *gin.Context) {
	var release models.Release
	err := ctx.ShouldBindJSON(&release)
	if err != nil {
		common.ResponseErrorBusiness(ctx, common.ErrorParams, "release params bind failed", err)
		return
	}

	svc := services.NewRelease(c.Engine, c.Cache)
	err = svc.Add(&release)
	if err != nil {
		common.ResponseErrorBusiness(ctx, common.ErrorMysql, "release save failed", err)
		return
	}

	common.ResponseSuccess(ctx, release)
}

func (c *Test) ReleaseDel(ctx *gin.Context) {
	var release models.Release
	err := ctx.ShouldBindJSON(&release)
	if err != nil {
		common.ResponseErrorBusiness(ctx, common.ErrorParams, "release params bind failed", err)
		return
	}

	svc := services.NewRelease(c.Engine, c.Cache)
	err = svc.Del(&release, true)
	if err != nil {
		common.ResponseErrorBusiness(ctx, common.ErrorMysql, "release delete failed", err)
		return
	}

	common.ResponseSuccess(ctx, release.AgentId)
}

func (c *Test) CardAdd(ctx *gin.Context) {
	var card models.Card
	err := ctx.ShouldBindJSON(&card)
	if err != nil {
		common.ResponseErrorBusiness(ctx, common.ErrorParams, "card params bind failed", err)
		return
	}

	svc := services.NewCard(c.Engine, c.Cache)
	err = svc.Add(&card)
	if err != nil {
		common.ResponseErrorBusiness(ctx, common.ErrorMysql, "card save failed", err)
		return
	}

	common.ResponseSuccess(ctx, card)
}

func (c *Test) CardDel(ctx *gin.Context) {
	var card models.Card
	err := ctx.ShouldBindJSON(&card)
	if err != nil {
		common.ResponseErrorBusiness(ctx, common.ErrorParams, "card params bind failed", err)
		return
	}

	svc := services.NewCard(c.Engine, c.Cache)
	err = svc.Del(&card, true)
	if err != nil {
		common.ResponseErrorBusiness(ctx, common.ErrorMysql, "card delete failed", err)
		return
	}

	common.ResponseSuccess(ctx, card.Name)
}

func (c *Test) TagAdd(ctx *gin.Context) {
	var tag models.Tag
	err := ctx.ShouldBindJSON(&tag)
	if err != nil {
		common.ResponseErrorBusiness(ctx, common.ErrorParams, "tag params bind failed", err)
		return
	}

	svc := services.NewTag(c.Engine)
	err = svc.Add(&tag)
	if err != nil {
		common.ResponseErrorBusiness(ctx, common.ErrorMysql, "tag save failed", err)
		return
	}

	common.ResponseSuccess(ctx, tag)
}

func (c *Test) TagDel(ctx *gin.Context) {
	var tag models.Tag
	err := ctx.ShouldBindJSON(&tag)
	if err != nil {
		common.ResponseErrorBusiness(ctx, common.ErrorParams, "tag params bind failed", err)
		return
	}

	svc := services.NewTag(c.Engine)
	err = svc.Del(&tag, true)
	if err != nil {
		common.ResponseErrorBusiness(ctx, common.ErrorMysql, "tag delete failed", err)
		return
	}

	common.ResponseSuccess(ctx, tag.Name)
}

func (c *Test) ConfigAdd(ctx *gin.Context) {
	var config models.Config
	err := ctx.ShouldBindJSON(&config)
	if err != nil {
		common.ResponseErrorBusiness(ctx, common.ErrorParams, "config params bind failed", err)
		return
	}

	svc := services.NewConfig(c.Engine, c.Cache)
	err = svc.Add(&config)
	if err != nil {
		common.ResponseErrorBusiness(ctx, common.ErrorMysql, "config save failed", err)
		return
	}

	common.ResponseSuccess(ctx, config)
}

func (c *Test) ConfigDel(ctx *gin.Context) {
	var config models.Config
	err := ctx.ShouldBindJSON(&config)
	if err != nil {
		common.ResponseErrorBusiness(ctx, common.ErrorParams, "config params bind failed", err)
		return
	}

	svc := services.NewConfig(c.Engine, c.Cache)
	err = svc.Del(&config, true)
	if err != nil {
		common.ResponseErrorBusiness(ctx, common.ErrorMysql, "config delete failed", err)
		return
	}

	common.ResponseSuccess(ctx, config.Name)
}

func (c *Test) MessageAdd(ctx *gin.Context) {
	var message models.Message
	err := ctx.ShouldBindJSON(&message)
	if err != nil {
		common.ResponseErrorBusiness(ctx, common.ErrorParams, "message params bind failed", err)
		return
	}

	svc := services.NewMessage(c.Engine)
	err = svc.Add(&message)
	if err != nil {
		common.ResponseErrorBusiness(ctx, common.ErrorMysql, "message save failed", err)
		return
	}

	common.ResponseSuccess(ctx, message)
}

func (c *Test) MessageDel(ctx *gin.Context) {
	var message models.Message
	err := ctx.ShouldBindJSON(&message)
	if err != nil {
		common.ResponseErrorBusiness(ctx, common.ErrorParams, "message params bind failed", err)
		return
	}

	svc := services.NewMessage(c.Engine)
	err = svc.Del(&message, true)
	if err != nil {
		common.ResponseErrorBusiness(ctx, common.ErrorMysql, "message delete failed", err)
		return
	}

	common.ResponseSuccess(ctx, message.Detail)
}
