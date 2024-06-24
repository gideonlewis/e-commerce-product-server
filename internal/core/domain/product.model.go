package domain

type Product struct {
	DefaultModel
	ID          int    `gorm:"column:product_id"`
	Name        string `gorm:"column:product_name"`
	Description string `gorm:"column:description"`
	CategoryID  int64  `gorm:"column:category_id"`
}

type Variant struct {
	ID           int     `gorm:""`
	ProductID    int     `gorm:""`
	VariantName  string  `gorm:""`
	VariantValue string  `gorm:""`
	Product      Product `gorm:""`
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
