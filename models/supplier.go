// C:\GoProject\src\eShop\models\supplier.go

package models

import "time"

type Supplier struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	Name        string     `gorm:"size:255;not null" json:"name"`         // Название или имя поставщика.
	ContactInfo string     `gorm:"size:255;not null" json:"contact_info"` // Контактная информация.
	CreatedAt   time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
	IsDeleted   bool       `gorm:"default:false" json:"is_deleted"`

	// Связи
	Products []Product `json:"products"`
}

func (Supplier) TableName() string {
	return "suppliers"
}
