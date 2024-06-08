package handler

import (
	"github.com/gideonlewis/e-commerce-product-server/internal/core/services"
)

type ProductHandler struct {
	svc services.ProductService
}

func NewProductHandler(productService services.ProductService) *ProductHandler {
	return &ProductHandler{
		svc: productService,
	}
}
