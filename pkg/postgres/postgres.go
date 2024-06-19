package postgres

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var instanceDB *gorm.DB

type Config struct {
	Host            string
	Port            int
	User            string
	Password        string
	Database        string
	MaxIdleConns    *int
	MaxOpenConns    *int
	ConnMaxLifetime *time.Duration
}

func NewConnection(cfg *Config) *gorm.DB {
	// singleton connection
	if instanceDB == nil {
		dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database)

		gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
		if err != nil {
			fmt.Println("Error creating: ", err.Error())
			panic(err)
		}

		sqlDB, _ := gormDB.DB()

		// Set connection pool settings
		if cfg.MaxIdleConns != nil {
			sqlDB.SetMaxIdleConns(*cfg.MaxIdleConns)
		}

		if cfg.MaxOpenConns != nil {
			sqlDB.SetMaxOpenConns(*cfg.MaxOpenConns)
		}

		if cfg.ConnMaxLifetime != nil {
			sqlDB.SetConnMaxLifetime(*cfg.ConnMaxLifetime)
		}

		instanceDB = gormDB
	}

	return instanceDB
}
