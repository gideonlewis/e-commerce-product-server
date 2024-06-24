package restful

import (
	"net/http"
	"strconv"

	"github.com/gideonlewis/e-commerce-product-server/internal/core/domain"
	"github.com/gin-gonic/gin"
)

func (h *ProductHandler) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, _ := strconv.Atoi(ctx.Param("id"))

		var req GetByIDProductRequest
		req.ID = int64(id)

		if err := ctx.BindQuery(&req); err != nil {
			handleError(ctx, http.StatusBadRequest, err)
			return
		}

		if err := req.validate(); err != nil {
			handleError(ctx, http.StatusBadRequest, err)
			return
		}

		data, err := h.svc.GetByID(req.ID, req.Unscope)
		if err != nil {
			handleError(ctx, http.StatusInternalServerError, err)
			return
		}

		handleSuccess(ctx, http.StatusOK, productResponse(data))
	}
}

func productResponse(data *domain.Product) *GetByIDProductResponse {
	resp := &GetByIDProductResponse{
		ID:          data.ID,
		Name:        data.Name,
		Description: data.Description,
	}

	return resp
}

type GetByIDProductRequest struct {
	ID      int64 `form:"id"`
	Unscope bool  `form:"unscope"`
}

func (r *GetByIDProductRequest) validate() error {
	return nil
}

type GetByIDProductResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
