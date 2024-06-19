package restful

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *CategoryHandler) GetList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req GetListCategoryRequest
		if err := ctx.Bind(&req); err != nil {
			handleError(ctx, http.StatusBadRequest, err)
			return
		}

		if err := req.validate(); err != nil {
			handleError(ctx, http.StatusBadRequest, err)
			return
		}

		data, err := h.svc.GetList()
		if err != nil {
			handleError(ctx, http.StatusInternalServerError, err)
			return
		}

		handleSuccess(ctx, http.StatusOK, data)
	}
}

type GetListCategoryRequest struct {
}

func (r *GetListCategoryRequest) validate() error {
	return nil
}
