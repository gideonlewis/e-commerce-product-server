package services

import (
	"github.com/gideonlewis/e-commerce-product-server/internal/core/ports"
)

type ProductService struct {
	repo ports.ProductRepository
}

func NewProductService(repo ports.ProductRepository) *ProductService {
	return &ProductService{
		repo: repo,
	}
}
