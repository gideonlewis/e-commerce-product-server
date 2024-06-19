package restful

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductCreateRequest struct{}

func (h *ProductHandler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req ProductCreateRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			handleError(ctx, http.StatusBadRequest, err)
			return
		}
	}
}
