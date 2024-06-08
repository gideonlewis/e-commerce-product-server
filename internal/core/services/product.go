package services

import (
	"github.com/gideonlewis/e-commerce-product-server/internal/core/ports"
)

type ProductService struct {
	repo ports.ProductService
}

func NewProductService(repo ports.ProductService) *ProductService {
	return &ProductService{
		repo: repo,
	}
}
