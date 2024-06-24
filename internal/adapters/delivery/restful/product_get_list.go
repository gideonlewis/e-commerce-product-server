package restful

import (
	"errors"
	"net/http"

	"github.com/gideonlewis/e-commerce-product-server/internal/core/domain"
	"github.com/gideonlewis/e-commerce-product-server/pkg/datatypes"
	"github.com/gin-gonic/gin"
)

func (h *ProductHandler) GetList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req GetListProductRequest
		if err := ctx.BindQuery(&req); err != nil {
			handleError(ctx, http.StatusBadRequest, err)
			return
		}

		if err := req.validate(); err != nil {
			handleError(ctx, http.StatusBadRequest, err)
			return
		}

		data, total, err := h.svc.GetList(req.Paginator, req.Search, req.OrderBy, req.SortBy.String())
		if err != nil {
			handleError(ctx, http.StatusInternalServerError, err)
			return
		}

		handleSuccess(ctx, http.StatusOK, productsResponse(data, total, req.Page, req.Limit))
	}
}

func productsResponse(data []*domain.Product, total int64, page, limit int) *GetListProductResponse {
	products := make([]GetProductResponse, 0, len(data))

	for _, item := range data {
		products = append(products, GetProductResponse{
			ID:          item.ID,
			Name:        item.Name,
			Description: item.Description,
		})
	}

	resp := &GetListProductResponse{
		Products: products,
		Metadata: map[string]interface{}{
			"total": total,
			"page":  page,
			"limit": limit,
		},
	}

	return resp
}

type GetListProductRequest struct {
	datatypes.Paginator
	datatypes.Pair
	Search string `form:"search"`
}

func (r *GetListProductRequest) validate() error {
	if !r.SortBy.IsNull() && !r.SortBy.IsValid() {
		return errors.New("invalid sort_by")
	}
	return nil
}

type GetProductResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type GetListProductResponse struct {
	Products []GetProductResponse   `json:"products"`
	Metadata map[string]interface{} `json:"metadata"`
}
