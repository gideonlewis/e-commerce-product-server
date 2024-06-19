package domain

type Category struct {
	DefaultModel
	ID       int       `gorm:"column:category_id"`
	Name     string    `gorm:"column:category_name"`
	Icon     string    `gorm:"column:category_icon"`
	ParentID *int      `gorm:"column:parent_id"`
	Parent   *Category `gorm:"foreignKey:ParentID;references:ID"`
}
