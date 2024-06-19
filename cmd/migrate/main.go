package main

import (
	"github.com/gideonlewis/e-commerce-product-server/internal/config"
	applogger "github.com/gideonlewis/e-commerce-product-server/pkg/logger"
	"github.com/gideonlewis/e-commerce-product-server/pkg/postgres"
	migrate "github.com/rubenv/sql-migrate"
)

func main() {
	logger := applogger.CreateLoggerInstant()

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

	migrations := &migrate.FileMigrationSource{
		Dir: "internal/migrations",
	}

	sqlDB, err := db.DB()
	if err != nil {
		logger.Fatalf("cannot get db connection: %v\n", err)
	}

	total, err := migrate.Exec(sqlDB, "postgres", migrations, migrate.Up)
	if err != nil {
		logger.Fatalf("cannot execute migration: %v\n", err)
	}

	logger.Infof("applied %d migrations\n", total)
}
