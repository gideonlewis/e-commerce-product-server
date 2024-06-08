package main

import (
	"github.com/gideonlewis/e-commerce-product-server/internal/adapters/cache"
	handler "github.com/gideonlewis/e-commerce-product-server/internal/adapters/delivery/restful"
	"github.com/gideonlewis/e-commerce-product-server/internal/adapters/repository"
	"github.com/gideonlewis/e-commerce-product-server/internal/config"
	"github.com/gideonlewis/e-commerce-product-server/internal/core/domain"
	"github.com/gideonlewis/e-commerce-product-server/internal/core/services"
	"github.com/gideonlewis/e-commerce-product-server/internal/pkg/postgres"
	_ "github.com/lib/pq"
)

func main() {
	err := config.LoadConfig("config", ".")
	if err != nil {
		panic(err)
	}

	db := postgres.NewConnection(&postgres.Config{
		Host:     config.Postgres.Host,
		Port:     config.Postgres.Port,
		User:     config.Postgres.User,
		Password: config.Postgres.Pass,
		Database: config.Postgres.Name,
	})

	db.AutoMigrate(&domain.Product{}, &domain.Category{}, &domain.Inventory{}, &domain.ProductImage{})
	cache := cache.NewService(
		config.Redis.Address,
		config.Redis.Password,
	)

	// store
	store := repository.NewDB(db, cache)

	// services
	productService := services.NewProductService(store)

	// handlers
	producthandler := handler.NewProductHandler(*productService)

	// http server
	server := handler.NewAPIHandler(producthandler)
	server.Start(server.WithSetMode("release"))
}
