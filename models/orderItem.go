// C:\GoProject\src\eShop\models\orderItem.go

package models

import "time"

type OrderItem struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	OrderID   uint       `gorm:"not null" json:"order_id"`   // Внешний ключ на заказ.
	ProductID uint       `gorm:"not null" json:"product_id"` // Внешний ключ на продукт.
	Quantity  int        `gorm:"not null" json:"quantity"`   // Количество товара в заказе.
	Price     float64    `gorm:"not null" json:"price"`      // Цена товара в момент продажи.
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	IsDeleted bool       `gorm:"default:false" json:"is_deleted"`

	// Связи
	Order   Order   `gorm:"foreignKey:OrderID" json:"order"`
	Product Product `gorm:"foreignKey:ProductID" json:"product"`
}

func (OrderItem) TableName() string {
	return "order_items"
}
