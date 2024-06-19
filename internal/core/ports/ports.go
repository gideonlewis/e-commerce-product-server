package ports

import "github.com/gideonlewis/e-commerce-product-server/internal/core/domain"

type ProductService interface {
}

type ProductRepository interface {
}

type CategoryService interface {
	Create(data *domain.Category) error
	GetList() ([]*domain.Category, error)
}

type CategoryRepository interface {
	CreateCategory(data *domain.Category) error
	GetListCategory() ([]*domain.Category, error)
}
