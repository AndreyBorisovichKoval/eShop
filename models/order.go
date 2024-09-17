// C:\GoProject\src\eShop\models\order.go

package models

import "time"

type Order struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	UserID      uint       `gorm:"not null" json:"user_id"`      // Внешний ключ на продавца.
	TotalAmount float64    `gorm:"not null" json:"total_amount"` // Общая сумма всех заказов.
	IsPaid      bool       `gorm:"default:false" json:"is_paid"` // Статус оплаты.
	CreatedAt   time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
	IsDeleted   bool       `gorm:"default:false" json:"is_deleted"`

	// Связи
	User       User        `gorm:"foreignKey:UserID" json:"user"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderID" json:"order_items"` // Связь с OrderItem
}

func (Order) TableName() string {
	return "orders"
}
