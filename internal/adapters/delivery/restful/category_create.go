package restful

import (
	"net/http"

	"github.com/gideonlewis/e-commerce-product-server/internal/core/domain"
	"github.com/gin-gonic/gin"
)

func (h *CategoryHandler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req CreateCategoryRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			handleError(ctx, http.StatusBadRequest, err)
			return
		}

		if err := req.validate(); err != nil {
			handleError(ctx, http.StatusBadRequest, err)
			return
		}

		data := req.toDomain()
		if err := h.svc.Create(data); err != nil {
			handleError(ctx, http.StatusInternalServerError, err)
			return
		}

		var resp = &CreateCategoryResponse{
			ID: data.ID,
		}

		handleSuccess(ctx, http.StatusCreated, resp)
	}
}

type CreateCategoryRequest struct {
	Name     string `json:"name"`
	Icon     string `json:"icon"`
	ParentID int    `json:"parent_id"`
}

func (r *CreateCategoryRequest) validate() error {
	return nil
}

func (r *CreateCategoryRequest) toDomain() *domain.Category {
	return &domain.Category{
		Name:     r.Name,
		Icon:     r.Icon,
		ParentID: &r.ParentID,
	}
}

type CreateCategoryResponse struct {
	ID int `json:"id"`
}
