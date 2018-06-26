package common

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	RequestParamsKey = "requestParams"
)

type Response struct {
	Result ErrorCode   `json:"result"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
}

func Redirect(ctx *gin.Context, uri string) {
	ctx.Redirect(http.StatusFound, uri)
}

func ResponseSuccess(ctx *gin.Context, data interface{}) {
	response := Response{
		Result: ErrorOk,
		Data:   data,
		Msg:    "success",
	}
	ctx.JSON(http.StatusOK, response)
	FormatResponseLog(ctx, response)
}

func ResponseErrorBusiness(ctx *gin.Context, code ErrorCode, desc string, err error) {
	if err != nil {
		desc += ":" + err.Error()
	}
	response := Response{
		Result: code,
		Data:   struct{}{},
		Msg:    desc,
	}
	ctx.JSON(http.StatusOK, response)
	ctx.Abort()
	FormatResponseLog(ctx, response)
}

func ResponseErrorServer(ctx *gin.Context, desc string) {
	response := Response{
		Result: ErrorServer,
		Data:   struct{}{},
		Msg:    desc,
	}
	ctx.JSON(http.StatusOK, response)
	ctx.Abort()
	FormatResponseLog(ctx, response)
}

func FormatResponseLog(ctx *gin.Context, response Response) {
	logRequest := ctx.MustGet(RequestParamsKey).(string)
	logResponse, _ := json.Marshal(response)
	if logRequest != "" {
		GetLogger().Infow("request:【" + logRequest + "】")
	}
	if string(logResponse) != "" {
		GetLogger().Infow("response:【" + string(logResponse) + "】")
	}
}
