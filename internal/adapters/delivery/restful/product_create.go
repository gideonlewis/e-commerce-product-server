package restful

import (
	"net/http"

	"github.com/gideonlewis/e-commerce-product-server/internal/core/domain"
	"github.com/gin-gonic/gin"
)

func (h *ProductHandler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req ProductCreateRequest
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

		var resp = &CreateProductResponse{
			ID: data.ID,
		}

		handleSuccess(ctx, http.StatusCreated, resp)
	}
}

type ProductCreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	CategoryID  int64  `json:"category_id"`
}

func (p *ProductCreateRequest) validate() error {
	return nil
}

func (p *ProductCreateRequest) toDomain() *domain.Product {
	return &domain.Product{
		Name:        p.Name,
		Description: p.Description,
		CategoryID:  p.CategoryID,
	}
}

type CreateProductResponse struct {
	ID int `json:"id"`
}
