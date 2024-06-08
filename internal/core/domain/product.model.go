package domain

import "time"

type Product struct {
	ID          int       `gorm:"primary_key;auto_increment"`
	CategoryID  int       `gorm:"not null"`
	ProductName string    `gorm:"type:varchar(255);not null"`
	Description string    `gorm:"type:text"`
	Price       float64   `gorm:"type:decimal(10,2);not null"`
	CreatedAt   time.Time `gorm:"type:timestamp;default:current_timestamp"`
}

type Variant struct {
	ID           int     `gorm:"primary_key;auto_increment"`
	ProductID    int     `gorm:"not null"`
	VariantName  string  `gorm:"type:varchar(255)"`
	VariantValue string  `gorm:"type:varchar(255)"`
	Product      Product `gorm:"foreignKey:ProductID;references:ProductID"`
}

type ProductImage struct {
	ImageID   int `gorm:"primary_key;auto_increment"`
	ProductID int `gorm:"not null"`
	VariantID int
	ImageURL  string  `gorm:"type:varchar(255)"`
	Product   Product `gorm:"foreignKey:ProductID;references:ProductID"`
	Variant   Variant `gorm:"foreignKey:VariantID;references:VariantID"`
}

type Inventory struct {
	ID        int     `gorm:"primary_key;auto_increment"`
	VariantID int     `gorm:"not null"`
	Quantity  int     `gorm:"not null"`
	Variant   Variant `gorm:"foreignKey:VariantID;references:VariantID"`
}
