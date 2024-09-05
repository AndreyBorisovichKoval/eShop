// C:\GoProject\src\eShop\models\category.go

package models

import "time"

type Category struct {
	ID        uint       `gorm:"primaryKey" json:"id"`           // Уникальный идентификатор категории товара.
	Title     string     `gorm:"size:255;not null" json:"title"` // Название категории товара.
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	IsDeleted bool       `gorm:"default:false" json:"is_deleted"`
}

func (Category) TableName() string {
	return "category"
}
