package middles

import (
	"encoding/json"
	"errors"
	"identify/backend/common"

	"github.com/gin-gonic/gin"
)

type rres struct {
	Result int
	Data   interface{}
	Msg    string
}

type params struct {
	AppId    string `json:"appId" binding:"required"`
	ClientId string `json:"clientId" binding:"required"`
	Token    string `json:"token" binding:"required"`
	Source   string `json:"source" binding:"required"`
	UniqueId string `json:"uniqueId"`
}

const (
	DeviceSource   string = "dev"
	StoryboxSource string = "rtoy"
)

func RAuthCheck(method string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var authParams params
		switch method {
		case "GET":
			authParams.AppId = ctx.Param("appId")
			authParams.Token = ctx.Param("token")
			authParams.ClientId = ctx.Param("clientId")
			authParams.Source = ctx.Param("source")
		case "POST":
			err := ctx.ShouldBindJSON(&authParams)
			if err != nil {
				common.ResponseErrorBusiness(ctx, common.ErrorParams, "params get failed", err)
			}
			ctx.Set("appId", authParams.AppId)
			ctx.Set("uniqueId", authParams.UniqueId)
		case "JSON":
			jsonParams := ctx.PostForm("json")
			err := json.Unmarshal([]byte(jsonParams), &authParams)
			if err != nil {
				common.ResponseErrorBusiness(ctx, common.ErrorParams, "params get failed", err)
			}
			ctx.Set("appId", authParams.AppId)
		}

		supertoken := common.GetConfig().App.Supertoken
		if authParams.Token != supertoken {
			switch authParams.Source {
			case DeviceSource:
				err := checkDevAppClient(authParams.AppId, authParams.Token, authParams.ClientId)
				if err != nil {
					common.ResponseErrorBusiness(ctx, common.ErrorAuth, "device app and client check failed", err)
				}
			case StoryboxSource:
			default:
				common.ResponseErrorBusiness(ctx, common.ErrorAuth, "request source matched failed", nil)
			}
		}
		return
	}
}

func checkDevAppClient(appId, token, clientId string) (err error) {
	conf := common.GetConfig().Service.Stoken
	url := conf.Url + "/accesskeys/check"

	params := map[string]string{
		"appId":       appId,
		"accessToken": token,
		"clientId":    clientId,
	}
	jsonStr, err := json.Marshal(params)
	if err != nil {
		return
	}

	response, err := common.DoHttpPost(url, jsonStr)
	if err != nil {
		return
	}
	result := rres{}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.Result != 0 {
		return errors.New(" stoken check failed: " + result.Msg)
	}
	return nil
}