// C:\GoProject\src\eShop\models\payment.go

package models

import "time"

type Payment struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	OrderID   uint       `gorm:"not null" json:"order_id"` // Внешний ключ на заказ.
	Amount    float64    `gorm:"not null" json:"amount"`   // Сумма платежа.
	PaidAt    time.Time  `json:"paid_at"`                  // Дата оплаты.
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	IsDeleted bool       `gorm:"default:false" json:"is_deleted"`

	// Связи
	Order Order `gorm:"foreignKey:OrderID" json:"order"`
}

func (Payment) TableName() string {
	return "payments"
}
