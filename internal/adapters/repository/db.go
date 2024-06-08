package repository

import (
	"github.com/gideonlewis/e-commerce-product-server/internal/adapters/cache"
	"gorm.io/gorm"
)

type DB struct {
	db    *gorm.DB
	cache *cache.Cache
}

// new database
func NewDB(db *gorm.DB, cache *cache.Cache) *DB {
	return &DB{
		db:    db,
		cache: cache,
	}
}
