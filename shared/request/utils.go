package request

import (
	"github.com/gin-gonic/gin"
	"sdm_demo_todolist/shared/resp"
)

func BindUri(ctx *gin.Context, uri interface{}) error {
	if err := ctx.ShouldBindUri(uri); err != nil {
		resp.Abort400BadUri(ctx, err)
		return err
	}
	return nil
}

func BindJSON(ctx *gin.Context, req interface{}) error {
	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp.Abort400BadJson(ctx, err)
		return err
	}
	return nil
}
