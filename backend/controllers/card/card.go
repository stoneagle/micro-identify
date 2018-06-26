package card

import (
	"identify/backend/common"
	"identify/backend/controllers"
	middles "identify/backend/middles"
	models "identify/backend/models/card"
	"identify/backend/rpc"
	services "identify/backend/services/card"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Card struct {
	controllers.Base
	RpcClient rpc.Image
}

func NewCard() *Card {
	card := &Card{}
	card.Prepare(common.ProjectCard)
	card.RpcClient = rpc.Image{
		Host: card.Config.Card.Rpc.Host,
		Port: card.Config.Card.Rpc.Port,
	}
	return card
}

func (c *Card) Router(router *gin.RouterGroup) {
	cards := router.Group("")
	cards.POST("check", middles.RAuthCheck("JSON"), c.Check)
	cards.POST("detail", middles.RAuthCheck("POST"), c.PostOne)
	cards.POST("", middles.RAuthCheck("POST"), c.PostOneSlime)
	cards.GET(":source/:clientId/:token/:appId/:uniqueId", middles.RAuthCheck("GET"), c.One)
	cards.DELETE("cache/release/:albumId", c.DeleteReleaseCache)
	cards.DELETE("cache/config/:uniqueId", c.DeleteConfigCache)
}

/*
 * 返回图片识别结果id
 */
func (c *Card) Check(ctx *gin.Context) {
	appId := ctx.MustGet("appId").(string)
	img, err := ctx.FormFile("img")
	if err != nil {
		common.ResponseErrorBusiness(ctx, common.ErrorParams, "img file upload failed", err)
		return
	}
	filePath := common.GetImagePath(c.Config.Card.Rpc.Img, c.Type) + img.Filename
	err = ctx.SaveUploadedFile(img, filePath)
	if err != nil {
		common.ResponseErrorBusiness(ctx, common.ErrorFiles, "img file save failed", err)
		return
	}

	imgUniqueId, err := c.RpcClient.Identify(appId, filePath, c.Type)
	if imgUniqueId <= 0 || err != nil {
		common.ResponseErrorBusiness(ctx, common.ErrorCardIdentify, "img can not identify", err)
		return
	}
	ret := map[string]interface{}{
		"uniqueId": imgUniqueId,
	}

	detail := ctx.MustGet("detail").(bool)
	if detail {
		card, successFlag := c.getCardDetail(strconv.Itoa(imgUniqueId), appId, ctx)
		if !successFlag {
			return
		}
		cardServiceFormat := formatCardService(card)
		ret["Card"] = cardServiceFormat
	}

	common.ResponseSuccess(ctx, ret)
}

/*
 * 获取card配置
 */
func (c *Card) One(ctx *gin.Context) {
	uniqueId := ctx.Param("uniqueId")
	appId := ctx.Param("appId")
	card, successFlag := c.getCardDetail(uniqueId, appId, ctx)
	if successFlag {
		common.ResponseSuccess(ctx, card)
	}
}

func (c *Card) PostOne(ctx *gin.Context) {
	uniqueId := ctx.MustGet("uniqueId").(string)
	appId := ctx.MustGet("appId").(string)
	card, successFlag := c.getCardDetail(uniqueId, appId, ctx)
	if successFlag {
		common.ResponseSuccess(ctx, card)
	}
}

func (c *Card) PostOneSlime(ctx *gin.Context) {
	uniqueId := ctx.MustGet("uniqueId").(string)
	appId := ctx.MustGet("appId").(string)
	card, successFlag := c.getCardDetail(uniqueId, appId, ctx)
	if successFlag {
		cardServiceFormat := formatCardService(card)
		common.ResponseSuccess(ctx, cardServiceFormat)
	}
}

func formatCardService(card models.Card) models.CardServiceModel {
	cardServiceFormat := models.CardServiceModel{
		Id:          card.Id,
		UniqueId:    card.UniqueId,
		Name:        card.Name,
		AlbumName:   card.Album.Name,
		AlbumSource: card.Album.Source,
	}
	cardServiceFormat.Messages = make([][]models.MessageServiceModel, 0)
	for _, config := range card.Configs {
		messageServiceFormtSlice := make([]models.MessageServiceModel, 0)
		for _, message := range config.Messages {
			messageServiceFormat := models.MessageServiceModel{
				Type:   message.Type,
				Detail: message.Detail,
			}
			messageServiceFormtSlice = append(messageServiceFormtSlice, messageServiceFormat)
		}
		if len(messageServiceFormtSlice) > 0 {
			cardServiceFormat.Messages = append(cardServiceFormat.Messages, messageServiceFormtSlice)
		}
	}
	return cardServiceFormat
}

func (c *Card) getCardDetail(uniqueId, appId string, ctx *gin.Context) (card models.Card, flag bool) {
	flag = false
	cardSvc := services.NewCard(c.Engine, c.Cache)
	card, err := cardSvc.GetByUniqueId(uniqueId)
	if err != nil {
		common.ResponseErrorBusiness(ctx, common.ErrorMysql, "card get by uniqueId failed", err)
		return
	}

	// 判断卡片所属专辑是否发布
	if card.Album.Release == 0 {
		common.ResponseErrorBusiness(ctx, common.ErrorCardDetail, "album not release yet", nil)
		return
	}

	// 判断对应agent是否发布
	releaseSvc := services.NewRelease(c.Engine, c.Cache)
	agentMap, err := releaseSvc.GetCache(card.AlbumId)
	if err != nil {
		common.ResponseErrorBusiness(ctx, common.ErrorRedis, "agent lists get from cache failed", err)
		return
	}
	if len(agentMap) == 0 {
		agentMap = make(map[string]int)
		releases, err := releaseSvc.GetByAlbumId(card.AlbumId)
		if err != nil {
			common.ResponseErrorBusiness(ctx, common.ErrorMysql, "release get by albumId failed", err)
			return
		}
		for _, one := range releases {
			agentMap[one.AgentId] = 1
		}
		err = releaseSvc.SetCache(card.AlbumId, agentMap)
		if err != nil {
			common.ResponseErrorBusiness(ctx, common.ErrorRedis, "agent lists set cache failed", err)
			return
		}
	}
	if _, ok := agentMap[appId]; !ok {
		common.ResponseErrorBusiness(ctx, common.ErrorCardDetail, "release agent list don't have appId:"+appId, nil)
		return
	}

	// 获取卡片配置信息
	configSvc := services.NewConfig(c.Engine, c.Cache)
	configs, err := configSvc.GetCache(uniqueId)
	if err != nil {
		common.ResponseErrorBusiness(ctx, common.ErrorRedis, "configs get from cache failed", err)
		return
	}
	if len(configs) == 0 {
		configs, err = configSvc.GetByCardId(card.Id)
		if err != nil {
			common.ResponseErrorBusiness(ctx, common.ErrorMysql, "configs get by card failed", err)
			return
		}
		err = configSvc.SetCache(uniqueId, configs)
		if err != nil {
			common.ResponseErrorBusiness(ctx, common.ErrorRedis, "configs set cache failed", err)
			return
		}
	}
	card.Configs = configs
	return card, true
}

/*
 * 删除card相关缓存
 */
func (c *Card) DeleteConfigCache(ctx *gin.Context) {
	uniqueId := ctx.Param("uniqueId")
	configSvc := services.NewConfig(c.Engine, c.Cache)
	err := configSvc.DelCache(uniqueId)
	if err != nil {
		common.ResponseErrorBusiness(ctx, common.ErrorRedis, "delete config cache failed", err)
		return
	}
	common.ResponseSuccess(ctx, struct{}{})
}

func (c *Card) DeleteReleaseCache(ctx *gin.Context) {
	albumId := ctx.Param("albumId")
	releaseSvc := services.NewRelease(c.Engine, c.Cache)
	err := releaseSvc.DelCache(albumId)
	if err != nil {
		common.ResponseErrorBusiness(ctx, common.ErrorRedis, "delete release agent list cache failed", err)
		return
	}
	common.ResponseSuccess(ctx, struct{}{})
}
