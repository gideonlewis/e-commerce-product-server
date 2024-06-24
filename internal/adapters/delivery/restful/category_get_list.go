package restful

import (
	"errors"
	"net/http"

	"github.com/gideonlewis/e-commerce-product-server/internal/core/domain"
	"github.com/gideonlewis/e-commerce-product-server/pkg/datatypes"
	"github.com/gin-gonic/gin"
)

func (h *CategoryHandler) GetList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req GetListCategoryRequest
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

		handleSuccess(ctx, http.StatusOK, catagoriesResponse(data, total, req.Page, req.Limit))
	}
}

func catagoriesResponse(data []*domain.Category, total int64, page, limit int) *GetListCategoryResponse {
	categories := make([]GetCategoryResponse, 0, len(data))

	for _, item := range data {
		categories = append(categories, GetCategoryResponse{
			ID:   item.ID,
			Name: item.Name,
			Icon: item.Icon,
		})
	}

	resp := &GetListCategoryResponse{
		Categories: categories,
		Metadata: map[string]interface{}{
			"total": total,
			"page":  page,
			"limit": limit,
		},
	}

	return resp
}

type GetListCategoryRequest struct {
	datatypes.Paginator
	datatypes.Pair
	Search string `form:"search"`
}

func (r *GetListCategoryRequest) validate() error {
	if !r.SortBy.IsNull() && !r.SortBy.IsValid() {
		return errors.New("invalid sort_by")
	}
	return nil
}

type GetCategoryResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type GetListCategoryResponse struct {
	Categories []GetCategoryResponse  `json:"categories"`
	Metadata   map[string]interface{} `json:"metadata"`
}
