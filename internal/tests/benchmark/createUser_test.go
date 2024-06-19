package benchmark

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/gideonlewis/e-commerce-product-server/internal/adapters/cache"
// 	"github.com/gideonlewis/e-commerce-product-server/internal/adapters/repository"
// 	"github.com/gideonlewis/e-commerce-product-server/internal/core/domain"

// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// func BenchmarkCreateUser(b *testing.B) {
// 	dsn := "postgres://test:test@localhost:5433/template1?sslmode=disable"
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
// 	if err != nil {
// 		panic(err)
// 	}
// 	db.AutoMigrate(&domain.Message{}, &domain.User{})

// 	redisCache := cache.NewService("localhost:6379", "")

// 	store := repository.NewDB(db, redisCache)

// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		email := fmt.Sprintf("test_user_%d@example.com", i)
// 		password := "password"
// 		// Delete user if it exists
// 		var user domain.User
// 		if err := db.Where("email = ?", email).First(&user).Error; err == nil {
// 			if err := db.Delete(&user).Error; err != nil {
// 				b.Fatalf("failed to delete user: %v", err)
// 			}
// 		}
// 		_, err := store.CreateUser(email, password)
// 		if err != nil {
// 			b.Fatalf("failed to create test user: %v", err)
// 		}
// 	}
// }
