package services

import (
	"fmt"

	"github.com/gideonlewis/e-commerce-product-server/internal/core/domain"
	"github.com/gideonlewis/e-commerce-product-server/internal/core/ports"
	"github.com/gideonlewis/e-commerce-product-server/pkg/datatypes"
	"gorm.io/gorm"
)

type ProductService struct {
	repo ports.ProductRepository
}

func NewProductService(repo ports.ProductRepository) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

func (s *ProductService) Create(data *domain.Product) error {
	return s.repo.CreateProduct(data)
}

func (s *ProductService) GetByID(id int64, unscope bool) (*domain.Product, error) {
	return s.repo.GetByIDProduct(id, unscope)
}

func (s *ProductService) GetList(paginator datatypes.Paginator, search string, orderBy string, sortBy string) ([]*domain.Product, int64, error) {
	paginator.Format()
	var (
		order      = make([]string, 0)
		conditions = map[string]interface{}{}
	)
	// Default order_by - product_name
	if orderBy == "" {
		orderBy = "product_name"
	}
	// Default sort_by - ASC
	if sortBy == "" {
		sortBy = string(datatypes.SortTypeASC)
	}

	order = append(order, fmt.Sprintf("%s %s", orderBy, sortBy))
	if search != "" {
		conditions["product_name"] = gorm.Expr("product_name AND product_name LIKE ?", "%"+search+"%")
	}

	return s.repo.GetListProduct(paginator, conditions, order)
}
