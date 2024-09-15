// C:\GoProject\src\eShop\models\supplier.go

package models

import "time"

type Supplier struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Title     string     `gorm:"size:255;not null" json:"title"` // Название или имя поставщика...
	Email     string     `gorm:"size:255" json:"email"`          // Электронная почта поставщика...
	Phone     string     `gorm:"size:15" json:"phone"`           // Телефон поставщика...
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	IsDeleted bool       `gorm:"default:false" json:"is_deleted"`

	// Связи
	Products []Product `json:"products"`
}

func (Supplier) TableName() string {
	return "suppliers"
}
