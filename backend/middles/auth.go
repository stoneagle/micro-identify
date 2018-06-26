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
	ClientId string `json:"clientId"`
	Token    string `json:"token" binding:"required"`
	UniqueId string `json:"uniqueId"`
	Detail   bool   `json:"detail"`
}

func RAuthCheck(method string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var authParams params
		switch method {
		case "GET":
			authParams.AppId = ctx.Param("appId")
			authParams.Token = ctx.Param("token")
			authParams.ClientId = ctx.Param("clientId")
		case "POST":
			err := ctx.ShouldBindJSON(&authParams)
			if err != nil {
				common.ResponseErrorBusiness(ctx, common.ErrorParams, "params get failed", err)
			}
			ctx.Set("appId", authParams.AppId)
			ctx.Set("uniqueId", authParams.UniqueId)
			requestParams, _ := json.Marshal(authParams)
			ctx.Set(common.RequestParamsKey, string(requestParams))
		case "JSON":
			jsonParams := ctx.PostForm("json")
			err := json.Unmarshal([]byte(jsonParams), &authParams)
			if err != nil {
				common.ResponseErrorBusiness(ctx, common.ErrorParams, "params get failed", err)
			}
			ctx.Set("appId", authParams.AppId)
			ctx.Set("detail", authParams.Detail)
			ctx.Set(common.RequestParamsKey, jsonParams)
		}

		supertoken := common.GetConfig().App.Supertoken
		if authParams.Token != supertoken {
			err := checkDevAppClient(authParams.AppId, authParams.Token, authParams.ClientId)
			if err != nil {
				common.ResponseErrorBusiness(ctx, common.ErrorAuth, "device app and client check failed", err)
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
