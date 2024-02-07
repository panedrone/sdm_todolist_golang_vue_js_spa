package resp

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ErrResp struct {
	Error string `json:"error"`
}

func Abort400BadUri(ctx *gin.Context, err error) {
	Abort400hBadRequest(ctx, fmt.Sprintf("Invalid URI: %s", err.Error()))
}

func Abort400hBadRequest(ctx *gin.Context, message string) {
	AbortWithErrResp(ctx, http.StatusBadRequest, message)
}

func Abort400BadJson(ctx *gin.Context, err error) {
	Abort400hBadRequest(ctx, fmt.Sprintf("Invalid JSON: %s", err.Error()))
}

func Abort404(ctx *gin.Context, err error) {
	AbortWithErrResp(ctx, http.StatusNotFound, err.Error())
}

func Abort500(ctx *gin.Context, err error) {
	AbortWithErrResp(ctx, http.StatusInternalServerError, err.Error())
}

func AbortWithErrResp(ctx *gin.Context, httpStatusCode int, message string) {
	err := ErrResp{
		Error: message,
	}
	ctx.AbortWithStatusJSON(httpStatusCode, err)
}

func JSON(ctx *gin.Context, httpStatusCode int, jsonObject interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(httpStatusCode, jsonObject)
}
