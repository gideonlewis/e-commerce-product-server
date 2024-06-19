package repository

import (
	"fmt"

	"github.com/gideonlewis/e-commerce-product-server/internal/core/domain"
)

func (c *DB) CreateCategory(data *domain.Category) error {
	req := c.db.Create(data)
	if req.Error != nil {
		return req.Error
	}

	if req.RowsAffected == 0 {
		return fmt.Errorf("category not created")
	}

	return nil
}

func (c *DB) GetListCategory() ([]*domain.Category, error) {
	var data []*domain.Category
	req := c.db.Find(&data)
	if req.Error != nil {
		return nil, req.Error
	}

	return data, nil
}
