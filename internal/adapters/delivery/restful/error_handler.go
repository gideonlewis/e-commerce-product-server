package restful

import "github.com/gin-gonic/gin"

func handleError(ctx *gin.Context, statusCode int, err error) {
	ctx.JSON(statusCode, gin.H{
		"error": err.Error(),
	})
}
