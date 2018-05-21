package card

import (
	"identify/backend/common"
	"identify/backend/controllers"
	"identify/backend/ipc"

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
	cards.GET("", c.One)
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
	// 判断卡片所属专辑是否发布
	// 获取卡片配置信息
	c.Success(ctx, struct{}{})
}

/*
 * 更新card相关缓存
 */
func (c *Card) UpdateCache(ctx *gin.Context) {
	// 更新卡片所属专辑发布缓存
	// 更新卡片配置信息
	c.Success(ctx, struct{}{})
}
