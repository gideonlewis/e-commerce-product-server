package main

import (
	"time"

	"github.com/gideonlewis/e-commerce-product-server/internal/config"
	"github.com/gideonlewis/e-commerce-product-server/internal/core/domain"
	applogger "github.com/gideonlewis/e-commerce-product-server/pkg/logger"
	"github.com/gideonlewis/e-commerce-product-server/pkg/postgres"
	"gorm.io/gorm"
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

	if err := gormMigrate(db); err != nil {
		logger.Fatalf("cannot execute migration: %v\n", err)
	}

	// Init data
	db.Create(categories())
	db.Create(products())
}

func gormMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		domain.Category{},
		domain.Product{},
	)
}

func categories() []domain.Category {
	return []domain.Category{
		{DefaultModel: domain.DefaultModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}, ID: 1, Name: "Thời Trang Nam", Icon: "Icon 1", ParentID: nil},
		{DefaultModel: domain.DefaultModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}, ID: 2, Name: "Thời Trang Nữ", Icon: "Icon 2", ParentID: nil},
		{DefaultModel: domain.DefaultModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}, ID: 3, Name: "Điện Thoại & Phụ Kiện", Icon: "Icon 3", ParentID: nil},
		{DefaultModel: domain.DefaultModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}, ID: 4, Name: "Máy Tính & Laptop", Icon: "Icon 4", ParentID: nil},
		{DefaultModel: domain.DefaultModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}, ID: 5, Name: "Mẹ & Bé", Icon: "Icon 5", ParentID: nil},
		{DefaultModel: domain.DefaultModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}, ID: 6, Name: "Máy Ảnh & Máy Quay Phim", Icon: "Icon 6", ParentID: nil},
		{DefaultModel: domain.DefaultModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}, ID: 7, Name: "Đồng Hồ", Icon: "Icon 7", ParentID: nil},
		{DefaultModel: domain.DefaultModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}, ID: 8, Name: "Giày Dép", Icon: "Icon 8", ParentID: nil},
		{DefaultModel: domain.DefaultModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}, ID: 9, Name: "Thiết Bị Gia Dụng", Icon: "Icon 9", ParentID: nil},
		{DefaultModel: domain.DefaultModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}, ID: 10, Name: "Thể Thao & Du Lịch", Icon: "Icon 10", ParentID: nil},
		{DefaultModel: domain.DefaultModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}, ID: 11, Name: "Nhà Cửa & Đời Sống", Icon: "Icon 11", ParentID: nil},
		{DefaultModel: domain.DefaultModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}, ID: 12, Name: "Sắc Đẹp", Icon: "Icon 12", ParentID: nil},
		{DefaultModel: domain.DefaultModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}, ID: 13, Name: "Sức Khoẻ", Icon: "Icon 13", ParentID: nil},
		{DefaultModel: domain.DefaultModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}, ID: 14, Name: "Trang Sức & Nữ Trang", Icon: "Icon 14", ParentID: nil},
	}
}

func products() []domain.Product {
	return []domain.Product{
		// Thời Trang Nam
		{ID: 1, Name: "Áo Sơ Mi Nam", Description: "Áo sơ mi nam chất liệu cotton", CategoryID: 1},
		{ID: 2, Name: "Quần Jeans Nam", Description: "Quần jeans nam phong cách", CategoryID: 1},
		{ID: 3, Name: "Áo Khoác Nam", Description: "Áo khoác nam mùa đông", CategoryID: 1},
		{ID: 4, Name: "Giày Tây Nam", Description: "Giày tây nam da thật", CategoryID: 1},
		{ID: 5, Name: "Thắt Lưng Nam", Description: "Thắt lưng nam da bò", CategoryID: 1},

		// Thời Trang Nữ
		{ID: 6, Name: "Đầm Dạ Hội", Description: "Đầm dạ hội sang trọng", CategoryID: 2},
		{ID: 7, Name: "Áo Khoác Nữ", Description: "Áo khoác nữ thời trang", CategoryID: 2},
		{ID: 8, Name: "Giày Cao Gót", Description: "Giày cao gót phong cách", CategoryID: 2},
		{ID: 9, Name: "Quần Legging", Description: "Quần legging nữ co giãn", CategoryID: 2},
		{ID: 10, Name: "Túi Xách Nữ", Description: "Túi xách nữ thời trang", CategoryID: 2},

		// Điện Thoại & Phụ Kiện
		{ID: 11, Name: "iPhone 13", Description: "Điện thoại iPhone 13 mới nhất", CategoryID: 3},
		{ID: 12, Name: "Samsung Galaxy S21", Description: "Điện thoại Samsung Galaxy S21", CategoryID: 3},
		{ID: 13, Name: "Ốp Lưng Điện Thoại", Description: "Ốp lưng bảo vệ điện thoại", CategoryID: 3},
		{ID: 14, Name: "Tai Nghe Bluetooth", Description: "Tai nghe Bluetooth không dây", CategoryID: 3},
		{ID: 15, Name: "Sạc Dự Phòng", Description: "Pin sạc dự phòng dung lượng cao", CategoryID: 3},

		// Máy Tính & Laptop
		{ID: 16, Name: "MacBook Air", Description: "Laptop MacBook Air M1", CategoryID: 4},
		{ID: 17, Name: "Dell XPS 13", Description: "Laptop Dell XPS 13", CategoryID: 4},
		{ID: 18, Name: "Bàn Phím Cơ", Description: "Bàn phím cơ chơi game", CategoryID: 4},
		{ID: 19, Name: "Chuột Không Dây", Description: "Chuột không dây tiện dụng", CategoryID: 4},
		{ID: 20, Name: "Balo Laptop", Description: "Balo chống sốc đựng laptop", CategoryID: 4},

		// Mẹ & Bé
		{ID: 21, Name: "Sữa Bột", Description: "Sữa bột dinh dưỡng cho bé", CategoryID: 5},
		{ID: 22, Name: "Xe Đẩy Em Bé", Description: "Xe đẩy cho bé tiện lợi", CategoryID: 5},
		{ID: 23, Name: "Ghế Ăn Dặm", Description: "Ghế ăn dặm cho bé", CategoryID: 5},
		{ID: 24, Name: "Đồ Chơi Giáo Dục", Description: "Đồ chơi giáo dục phát triển trí tuệ", CategoryID: 5},
		{ID: 25, Name: "Bình Sữa", Description: "Bình sữa chống sặc", CategoryID: 5},

		// Máy Ảnh & Máy Quay Phim
		{ID: 26, Name: "Máy Ảnh Canon", Description: "Máy ảnh Canon EOS R5", CategoryID: 6},
		{ID: 27, Name: "Máy Quay GoPro", Description: "Máy quay GoPro Hero 9", CategoryID: 6},
		{ID: 28, Name: "Ống Kính Máy Ảnh", Description: "Ống kính máy ảnh chất lượng cao", CategoryID: 6},
		{ID: 29, Name: "Chân Máy Ảnh", Description: "Chân máy ảnh chuyên nghiệp", CategoryID: 6},
		{ID: 30, Name: "Thẻ Nhớ SD", Description: "Thẻ nhớ SD dung lượng lớn", CategoryID: 6},

		// Đồng Hồ
		{ID: 31, Name: "Đồng Hồ Rolex", Description: "Đồng hồ Rolex sang trọng", CategoryID: 7},
		{ID: 32, Name: "Đồng Hồ Casio", Description: "Đồng hồ Casio bền bỉ", CategoryID: 7},
		{ID: 33, Name: "Đồng Hồ Thông Minh", Description: "Đồng hồ thông minh Apple Watch", CategoryID: 7},
		{ID: 34, Name: "Đồng Hồ Fossil", Description: "Đồng hồ Fossil thời trang", CategoryID: 7},
		{ID: 35, Name: "Đồng Hồ Seiko", Description: "Đồng hồ Seiko chính hãng", CategoryID: 7},

		// Giày Dép
		{ID: 36, Name: "Giày Thể Thao", Description: "Giày thể thao Adidas", CategoryID: 8},
		{ID: 37, Name: "Dép Quai Ngang", Description: "Dép quai ngang thoải mái", CategoryID: 8},
		{ID: 38, Name: "Giày Lười Nam", Description: "Giày lười nam tiện dụng", CategoryID: 8},
		{ID: 39, Name: "Giày Cao Gót Nữ", Description: "Giày cao gót nữ thanh lịch", CategoryID: 8},
		{ID: 40, Name: "Dép Xỏ Ngón", Description: "Dép xỏ ngón phong cách", CategoryID: 8},

		// Thiết Bị Gia Dụng
		{ID: 41, Name: "Máy Giặt", Description: "Máy giặt cửa ngang LG", CategoryID: 9},
		{ID: 42, Name: "Tủ Lạnh", Description: "Tủ lạnh Inverter tiết kiệm điện", CategoryID: 9},
		{ID: 43, Name: "Lò Vi Sóng", Description: "Lò vi sóng đa năng", CategoryID: 9},
		{ID: 44, Name: "Máy Hút Bụi", Description: "Máy hút bụi công suất lớn", CategoryID: 9},
		{ID: 45, Name: "Quạt Điều Hòa", Description: "Quạt điều hòa làm mát không khí", CategoryID: 9},

		// Thể Thao & Du Lịch
		{ID: 46, Name: "Xe Đạp Địa Hình", Description: "Xe đạp địa hình bền bỉ", CategoryID: 10},
		{ID: 47, Name: "Balo Du Lịch", Description: "Balo du lịch đa năng", CategoryID: 10},
		{ID: 48, Name: "Giày Chạy Bộ", Description: "Giày chạy bộ Nike", CategoryID: 10},
		{ID: 49, Name: "Lều Cắm Trại", Description: "Lều cắm trại chống nước", CategoryID: 10},
		{ID: 50, Name: "Dụng Cụ Leo Núi", Description: "Dụng cụ leo núi an toàn", CategoryID: 10},

		// Nhà Cửa & Đời Sống
		{ID: 51, Name: "Bộ Chăn Ga", Description: "Bộ chăn ga cao cấp", CategoryID: 11},
		{ID: 52, Name: "Đèn Trang Trí", Description: "Đèn trang trí phòng khách", CategoryID: 11},
		{ID: 53, Name: "Bình Hoa", Description: "Bình hoa thủy tinh", CategoryID: 11},
		{ID: 54, Name: "Thảm Trải Sàn", Description: "Thảm trải sàn êm ái", CategoryID: 11},
		{ID: 55, Name: "Giường Ngủ", Description: "Giường ngủ gỗ tự nhiên", CategoryID: 11},

		// Sắc Đẹp
		{ID: 56, Name: "Kem Dưỡng Da", Description: "Kem dưỡng da chống lão hóa", CategoryID: 12},
	}
}

// func sqlMigrate(db *gorm.DB) error {
// 	sqlDB, err := db.DB()
// 	if err != nil {
// 		return err
// 	}

// 	migrations := &migrate.FileMigrationSource{
// 		Dir: "internal/migrations",
// 	}

// 	total, err := migrate.Exec(sqlDB, "postgres", migrations, migrate.Up)
// 	if err != nil {
// 		return err
// 	}

// 	logger.Infof("applied %d migrations\n", total)
// 	return nil
// }
