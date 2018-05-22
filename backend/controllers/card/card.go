package card

import (
	"identify/backend/common"
	"identify/backend/controllers"
	"identify/backend/ipc"
	services "identify/backend/services/card"

	"github.com/gin-gonic/gin"
)

type Card struct {
	controllers.Base
}

func NewCard() *Card {
	card := &Card{}
	card.Prepare(common.ProjectCard)
	return card
}

func (c *Card) Router(router *gin.RouterGroup) {
	cards := router.Group("")
	cards.POST("check", c.Check)
	cards.DELETE("cache/release/:albumId", c.DeleteReleaseCache)
	cards.DELETE("cache/config/:uniqueId", c.DeleteConfigCache)
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

	cardSvc := services.NewCard(c.Engine, c.Cache)
	card, err := cardSvc.GetByUniqueId(uniqueId)
	if err != nil {
		c.ErrorBusiness(ctx, common.ErrorMysql, "card get by uniqueId failed", err)
		return
	}

	// 判断卡片所属专辑是否发布
	if card.Album.Release == 0 {
		c.ErrorBusiness(ctx, common.ErrorCardDetail, "album not release yet", nil)
		return
	}

	// 判断对应agent是否发布
	releaseSvc := services.NewRelease(c.Engine, c.Cache)
	agentMap, err := releaseSvc.GetCache(card.AlbumId)
	if err != nil {
		c.ErrorBusiness(ctx, common.ErrorRedis, "agent lists get from cache failed", err)
		return
	}
	if len(agentMap) == 0 {
		agentMap = make(map[string]int)
		releases, err := releaseSvc.GetByAlbumId(card.AlbumId)
		if err != nil {
			c.ErrorBusiness(ctx, common.ErrorMysql, "release get by albumId failed", err)
			return
		}
		for _, one := range releases {
			agentMap[one.AgentId] = 1
		}
		err = releaseSvc.SetCache(card.AlbumId, agentMap)
		if err != nil {
			c.ErrorBusiness(ctx, common.ErrorRedis, "agent lists set cache failed", err)
			return
		}
	}
	if _, ok := agentMap[agentId]; !ok {
		c.ErrorBusiness(ctx, common.ErrorCardDetail, "release agent list don't have agentId:"+agentId, nil)
		return
	}

	// 获取卡片配置信息
	configSvc := services.NewConfig(c.Engine, c.Cache)
	configs, err := configSvc.GetCache(uniqueId)
	if err != nil {
		c.ErrorBusiness(ctx, common.ErrorRedis, "configs get from cache failed", err)
		return
	}
	if len(configs) == 0 {
		configs, err = configSvc.GetByCardId(card.Id)
		if err != nil {
			c.ErrorBusiness(ctx, common.ErrorMysql, "configs get by card failed", err)
			return
		}
		err = configSvc.SetCache(uniqueId, configs)
		if err != nil {
			c.ErrorBusiness(ctx, common.ErrorRedis, "configs set cache failed", err)
			return
		}
	}
	card.Configs = configs

	c.Success(ctx, card)
}

/*
 * 删除card相关缓存
 */
func (c *Card) DeleteConfigCache(ctx *gin.Context) {
	uniqueId := ctx.Param("uniqueId")
	configSvc := services.NewConfig(c.Engine, c.Cache)
	err := configSvc.DelCache(uniqueId)
	if err != nil {
		c.ErrorBusiness(ctx, common.ErrorRedis, "delete config cache failed", err)
		return
	}
	c.Success(ctx, struct{}{})
}

func (c *Card) DeleteReleaseCache(ctx *gin.Context) {
	albumId := ctx.Param("albumId")
	releaseSvc := services.NewRelease(c.Engine, c.Cache)
	err := releaseSvc.DelCache(albumId)
	if err != nil {
		c.ErrorBusiness(ctx, common.ErrorRedis, "delete release agent list cache failed", err)
		return
	}
	c.Success(ctx, struct{}{})
}
