package card

import (
	"identify/backend/common"
	"identify/backend/controllers"
	"identify/backend/ipc"
	svc "identify/backend/services/card"

	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
)

type Card struct {
	controllers.Base
	Service *svc.Card
}

func NewCard(engine *xorm.Engine) *Card {
	card := &Card{}
	card.Init()
	card.Service = svc.NewCard(card.Engine)
	card.Engine = engine
	card.Type = common.ProjectCard
	return card
}

func (c *Card) Router(router *gin.RouterGroup) {
	users := router.Group("")
	users.POST("check", c.Check)
	users.GET("", c.One)
}

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
	ret := data.Check()
	c.Success(ctx, ret)
}

func (c *Card) One(ctx *gin.Context) {
	data := ipc.IData{}
	modelPath := c.Config.Card.Ipc.Model
	data.SetParams(modelPath, "", "", c.Type)
	ret := data.Check()
	c.Success(ctx, ret)
}
