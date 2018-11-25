package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseWithCount struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Count   int         `json:"count"`
}

func SendResponse(context *gin.Context, response Response) {
	context.JSON(http.StatusOK, response)
}

func SendResponseWithError(context *gin.Context, err error, response Response) {
	if err != nil {
		context.AbortWithError(
			404, err)
		return
	}
	context.JSON(http.StatusOK, response)
}
func SendResponseCount(context *gin.Context, response ResponseWithCount) {
	context.JSON(http.StatusOK, response)
}
func SendResponseCountWithError(context *gin.Context, err error, response ResponseWithCount) {
	if err != nil {
		context.AbortWithError(404, err)
		return
	}

	context.JSON(http.StatusOK, response)
}
