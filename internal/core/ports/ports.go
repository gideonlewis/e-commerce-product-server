package ports

import (
	"github.com/gideonlewis/e-commerce-product-server/internal/core/domain"
	"github.com/gideonlewis/e-commerce-product-server/pkg/datatypes"
)

type ProductService interface {
	Create(data *domain.Product) error
	GetByID(id int64, unscope bool) (*domain.Product, error)
	GetList() ([]*domain.Product, error)
}

type ProductRepository interface {
	CreateProduct(data *domain.Product) error
	GetByIDProduct(id int64, unscope bool) (*domain.Product, error)
	GetListProduct(paginator datatypes.Paginator, condition interface{}, order []string) ([]*domain.Product, int64, error)
}

type CategoryService interface {
	Create(data *domain.Category) error
	GetList() ([]*domain.Category, error)
}

type CategoryRepository interface {
	CreateCategory(data *domain.Category) error
	GetListCategory(paginator datatypes.Paginator, condition interface{}, order []string) ([]*domain.Category, int64, error)
}
