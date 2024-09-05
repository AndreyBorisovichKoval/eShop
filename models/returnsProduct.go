// C:\GoProject\src\eShop\models\returnsProduct.go

package models

import "time"

type ReturnsProduct struct {
	ID           uint       `gorm:"primaryKey" json:"id"`
	ProductID    uint       `gorm:"not null" json:"product_id"`             // Внешний ключ на продукт.
	SupplierID   uint       `gorm:"not null" json:"supplier_id"`            // Внешний ключ на поставщика.
	Quantity     float64    `gorm:"not null" json:"quantity"`               // Количество возвращаемого товара.
	ReturnReason string     `gorm:"size:255;not null" json:"return_reason"` // Причина возврата.
	ReturnedAt   time.Time  `gorm:"not null" json:"returned_at"`            // Дата возврата.
	CreatedAt    time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
	IsDeleted    bool       `gorm:"default:false" json:"is_deleted"`

	// Связи
	Product  Product  `gorm:"foreignKey:ProductID" json:"product"`
	Supplier Supplier `gorm:"foreignKey:SupplierID" json:"supplier"`
}

func (ReturnsProduct) TableName() string {
	return "returns_product"
}
