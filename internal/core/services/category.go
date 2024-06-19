package services

import (
	"github.com/gideonlewis/e-commerce-product-server/internal/core/domain"
	"github.com/gideonlewis/e-commerce-product-server/internal/core/ports"
)

type CategoryService struct {
	repo ports.CategoryRepository
}

func NewCategoryService(repo ports.CategoryRepository) *CategoryService {
	return &CategoryService{
		repo: repo,
	}
}

func (s *CategoryService) Create(data *domain.Category) error {
	return s.repo.CreateCategory(data)
}

func (s *CategoryService) GetList() ([]*domain.Category, error) {
	return s.repo.GetListCategory()
}
