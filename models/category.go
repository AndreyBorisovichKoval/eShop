// C:\GoProject\src\eShop\models\category.go

package models

import "time"

type Category struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	Title       string     `gorm:"size:255;not null;unique" json:"title"`       // Название категории
	Description string     `gorm:"size:500" json:"description"`                 // Описание категории
	CreatedAt   time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"` // Дата создания
	UpdatedAt   *time.Time `json:"updated_at"`                                  // Дата последнего обновления
	DeletedAt   *time.Time `json:"deleted_at"`                                  // Дата удаления
	IsDeleted   bool       `gorm:"default:false" json:"is_deleted"`             // Флаг мягкого удаления

	// Связи
	Products []Product `json:"products"` // Продукты, связанные с категорией
}

func (Category) TableName() string {
	return "categories"
}
