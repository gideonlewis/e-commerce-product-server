package services

import (
	"fmt"

	"github.com/gideonlewis/e-commerce-product-server/internal/core/domain"
	"github.com/gideonlewis/e-commerce-product-server/internal/core/ports"
	"github.com/gideonlewis/e-commerce-product-server/pkg/datatypes"
	"gorm.io/gorm"
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

func (s *CategoryService) GetList(paginator datatypes.Paginator, search string, orderBy string, sortBy string) ([]*domain.Category, int64, error) {
	paginator.Format()
	var (
		order      = make([]string, 0)
		conditions = map[string]interface{}{}
	)
	// Default order_by - category_name
	if orderBy == "" {
		orderBy = "category_name"
	}
	// Default sort_by - ASC
	if sortBy == "" {
		sortBy = string(datatypes.SortTypeASC)
	}

	order = append(order, fmt.Sprintf("%s %s", orderBy, sortBy))
	if search != "" {
		conditions["category_name"] = gorm.Expr("category_name AND category_name LIKE ?", "%"+search+"%")
	}

	return s.repo.GetListCategory(paginator, conditions, order)
}
