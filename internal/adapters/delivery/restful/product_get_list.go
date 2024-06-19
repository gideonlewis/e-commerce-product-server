package restful

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductGetListRequest struct{}

// func (r *ProductGetListRequest) Bind
type ProductGetListResponse struct{}

func (h *ProductHandler) GetList(ctx *gin.Context) {
	var req ProductGetListRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		handleError(ctx, http.StatusBadRequest, err)
		return
	}
}
