package restful

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleSuccess(ctx *gin.Context, statusCode int, data interface{}) {
	ctx.JSON(statusCode, gin.H{
		"code":    statusCode,
		"message": http.StatusText(statusCode),
		"data":    data,
	})
}
