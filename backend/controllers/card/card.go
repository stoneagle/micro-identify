package card

import (
	"identify/backend/common"
	"identify/backend/controllers"
	"identify/backend/ipc"
	models "identify/backend/models/card"
	services "identify/backend/services/card"

	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
)

type Card struct {
	controllers.Base
}

func NewCard(engine *xorm.Engine) *Card {
	card := &Card{}
	card.Init()
	card.Engine = engine
	card.Type = common.ProjectCard
	return card
}

func (c *Card) Router(router *gin.RouterGroup) {
	cards := router.Group("")
	cards.POST("check", c.Check)
	cards.POST("update/cache", c.UpdateCache)
	cards.GET(":agentId/:uniqueId", c.One)
}

/*
 * 返回图片识别结果id
 */
func (c *Card) Check(ctx *gin.Context) {
	appId := ctx.PostForm("appId")
	img, err := ctx.FormFile("img")
	if err != nil {
		c.ErrorBusiness(ctx, common.ErrorParams, "img file upload failed", err)
		return
	}
	filePath := ipc.GetImagePath(c.Config.Card.Ipc.Img, c.Type) + img.Filename
	err = ctx.SaveUploadedFile(img, filePath)
	if err != nil {
		c.ErrorBusiness(ctx, common.ErrorFiles, "img file save failed", err)
		return
	}

	data := ipc.IData{}
	modelPath := c.Config.Card.Ipc.Model
	data.SetParams(modelPath, filePath, appId, c.Type)
	imgUniqueId := data.Check()
	if imgUniqueId <= 0 {
		c.ErrorBusiness(ctx, common.ErrorCardIdentify, "img can not identify", nil)
		return
	}
	ret := map[string]int{
		"UniqueId": imgUniqueId,
	}
	c.Success(ctx, ret)
}

/*
 * 获取card配置
 */
func (c *Card) One(ctx *gin.Context) {
	uniqueId := ctx.Param("uniqueId")
	agentId := ctx.Param("agentId")

	cardSvc := services.NewCard(c.Engine)
	card, err := cardSvc.GetByUniqueId(uniqueId)
	if err != nil {
		c.ErrorBusiness(ctx, common.ErrorMysql, "card get by uniqueId failed", err)
		return
	}

	// 判断卡片所属专辑是否发布
	album := models.Album{}
	album.GeneralWithDeleted.Id = card.AlbumId
	albumSvc := services.NewAlbum(c.Engine)
	err = albumSvc.Get(&album)
	if err != nil {
		c.ErrorBusiness(ctx, common.ErrorMysql, "album get by Id failed", err)
		return
	}
	if album.Release == 0 {
		c.ErrorBusiness(ctx, common.ErrorCardDetail, "album not release yet", nil)
		return
	}

	// 判断对应agent是否发布
	releaseSvc := services.NewRelease(c.Engine)
	releases, err := releaseSvc.GetByAlbumId(card.AlbumId)
	if err != nil {
		c.ErrorBusiness(ctx, common.ErrorMysql, "release get by albumId failed", err)
		return
	}
	releaseFlag := false
	for _, one := range releases {
		if agentId == one.AgentId {
			releaseFlag = true
			break
		}
	}
	if !releaseFlag {
		c.ErrorBusiness(ctx, common.ErrorCardDetail, "release agent list don't have agentId:"+agentId, nil)
		return
	}

	// 获取卡片配置信息
	configSvc := services.NewConfig(c.Engine)
	configs, err := configSvc.GetByCardId(card.Id)
	if err != nil {
		c.ErrorBusiness(ctx, common.ErrorMysql, "configs get by card failed", err)
		return
	}
	card.Configs = configs
	c.Success(ctx, card)
}

/*
 * 更新card相关缓存
 */
func (c *Card) UpdateCache(ctx *gin.Context) {
	// 更新卡片所属专辑发布缓存
	// 更新卡片配置信息
	c.Success(ctx, struct{}{})
}
