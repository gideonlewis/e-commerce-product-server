package integration

// import (
// 	"errors"
// 	"io/fs"
// 	"os"
// 	"testing"

// 	"github.com/gideonlewis/e-commerce-product-server/internal/adapters/cache"
// 	"github.com/gideonlewis/e-commerce-product-server/internal/adapters/repository"
// 	"github.com/gideonlewis/e-commerce-product-server/internal/core/domain"

// 	"github.com/sirupsen/logrus"
// 	"golang.org/x/crypto/bcrypt"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// var store *repository.DB
// var logger *logrus.Logger

// func TestMain(m *testing.M) {
// 	dsn := "postgres://test:test@localhost:5433/template1?sslmode=disable"
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
// 	if err != nil {
// 		panic(err)
// 	}
// 	db.AutoMigrate(&domain.Message{}, &domain.User{})
// 	// defer db.Close()

// 	redisCache := cache.NewService("localhost:6379", "")

// 	store = repository.NewDB(db, redisCache)

// 	// defer store.Close()

// 	code := m.Run()

// 	os.Exit(code)
// }

// func TestDBIntegration(t *testing.T) {

// 	// create a test user
// 	email := "test1@example.com"
// 	password := "password"
// 	user, err := store.CreateUser(email, password)
// 	if err != nil {
// 		t.Fatalf("failed to create test user: %v", err)
// 	}

// 	// test reading a user
// 	readUser, err := store.ReadUser(user.ID)
// 	if err != nil {
// 		t.Fatalf("failed to read user: %v", err)
// 	}
// 	if readUser.Email != email {
// 		t.Errorf("expected email %q, got %q", email, readUser.Email)
// 	}

// 	// test reading all users
// 	users, err := store.ReadUsers()
// 	if err != nil {
// 		t.Fatalf("failed to read users: %v", err)
// 	}
// 	if len(users) != 1 {
// 		t.Errorf("expected 1 user, got %d", len(users))
// 	}
// 	if users[0].Email != email {
// 		t.Errorf("expected email %q, got %q", email, users[0].Email)
// 	}

// 	// test updating a user
// 	newEmail := "newemail@example.com"
// 	newPassword := "newpassword"
// 	err = store.UpdateUser(user.ID, newEmail, newPassword, newPassword)
// 	logger.WithField("updated user", user).Debugf("🪺updated user: %v", user)

// 	if err != nil {
// 		t.Fatalf("failed to update user: %v", err)
// 	}
// 	readUser, err = store.ReadUser(user.ID)
// 	logger.WithField("readUser", readUser).Debugf("🍄readUser: %v", readUser)

// 	if err != nil {
// 		t.Fatalf("failed to read updated user: %v", err)
// 	}
// 	if readUser.Email != newEmail {
// 		t.Errorf("expected email %q, got %q", newEmail, readUser.Email)
// 	}

// 	hashedNewPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
// 	if err != nil {
// 		t.Errorf("password not hashed: %v", err)
// 	}

// 	if readUser.Password != string(hashedNewPassword) {
// 		t.Errorf("expected password %q, got %q", newPassword, readUser.Password)
// 	}

// 	// test deleting a user
// 	err = store.DeleteUser(user.ID)
// 	if err != nil {
// 		t.Fatalf("failed to delete user: %v", err)
// 	}
// 	_, err = store.ReadUser(user.ID)
// 	if !errors.Is(err, fs.ErrNotExist) {
// 		t.Errorf("expected user not found error, got %v", err)
// 	}

// 	// test deleting the same user again should return error
// 	err = store.DeleteUser(user.ID)
// 	if !errors.Is(err, fs.ErrNotExist) {
// 		t.Errorf("expected user not found error, got %v", err)
// 	}
// }
