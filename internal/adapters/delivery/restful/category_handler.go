package restful

import (
	"github.com/gideonlewis/e-commerce-product-server/internal/core/services"
)

type CategoryHandler struct {
	svc *services.CategoryService
}

func NewCategoryHandler(categoryService *services.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		svc: categoryService,
	}
}
