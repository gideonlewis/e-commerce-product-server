package domain

type Category struct {
	CategoryID   int       `json:"category_id" gorm:"primary_key;auto_increment"`
	CategoryName string    `json:"category_name" gorm:"type:varchar(255);not null"`
	ParentID     int       `json:"parent_id" gorm:"index"`
	Parent       *Category `json:"parent,omitempty" gorm:"foreignKey:ParentID;references:CategoryID"`
}
