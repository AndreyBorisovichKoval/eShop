// C:\GoProject\src\eShop\models\returns_product.go

package models

import "time"

type ReturnsProduct struct {
	ID           uint    `gorm:"primaryKey" json:"id"`
	ProductID    uint    `gorm:"not null" json:"product_id"`
	SupplierID   uint    `gorm:"not null" json:"supplier_id"`
	Quantity     float64 `gorm:"not null" json:"quantity"`
	ReturnReason string  `gorm:"size:255;not null" json:"return_reason"`
	// ReturnedAt   time.Time  `gorm:"not null" json:"returned_at"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	IsDeleted bool       `gorm:"default:false" json:"is_deleted"`

	Product  Product  `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	Supplier Supplier `gorm:"foreignKey:SupplierID" json:"supplier,omitempty"`
}

func (ReturnsProduct) TableName() string {
	return "returns_products"
}

// ReturnResponse - кастомная структура для возвратов.
type ReturnResponse struct {
	ID           uint    `json:"id"`            // ID записи возврата
	ProductName  string  `json:"product_name"`  // Название товара
	SupplierName string  `json:"supplier_name"` // Название поставщика
	Quantity     float64 `json:"quantity"`      // Количество
	ReturnReason string  `json:"return_reason"` // Причина возврата
	ReturnedAt   string  `json:"returned_at"`   // Дата возврата (дата создания записи)
}
