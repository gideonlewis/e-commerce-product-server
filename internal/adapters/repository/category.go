package repository

import (
	"fmt"

	"github.com/gideonlewis/e-commerce-product-server/internal/core/domain"
	"github.com/gideonlewis/e-commerce-product-server/pkg/datatypes"
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

func (c *DB) GetListCategory(
	paginator datatypes.Paginator,
	conditions interface{},
	order []string,
) ([]*domain.Category, int64, error) {
	var (
		db     = c.db.Model(&domain.Category{})
		data   = make([]*domain.Category, 0)
		total  int64
		offset int
	)

	if conditions != nil {
		db = db.Where(conditions)
	}

	for i := range order {
		db = db.Order(order[i])
	}

	if paginator.Page != 1 {
		offset = paginator.Limit * (paginator.Page - 1)
	}

	if paginator.Limit != -1 {
		err := db.Count(&total).Error
		if err != nil {
			return nil, 0, err
		}
	}

	err := db.Limit(paginator.Limit).Offset(offset).Find(&data).Error
	if err != nil {
		return nil, 0, err
	}

	if paginator.Limit == -1 {
		total = int64(len(data))
	}

	return data, total, nil
}
