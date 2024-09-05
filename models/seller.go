// C:\GoProject\src\eShop\models\seller.go

package models

import "time"

type Seller struct {
	ID        uint       `gorm:"primaryKey" json:"id"`                          // Уникальный идентификатор продавца.
	FullName  string     `gorm:"size:255;not null" json:"full_name"`            // Полное имя продавца.
	UserName  string     `gorm:"size:255;not null" json:"user_name"`            // Логин продавца.
	Email     string     `gorm:"size:255;not null;unique" json:"email"`         // Email продавца.
	Password  string     `gorm:"size:255;not null" json:"password"`             // Пароль продавца.
	Role      string     `gorm:"size:50;not null;default:'seller'" json:"role"` // Роль продавца (например: 'admin', 'seller').
	IsBlocked bool       `gorm:"default:false" json:"is_blocked"`               // Заблокирован ли продавец.
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	IsDeleted bool       `gorm:"default:false" json:"is_deleted"`

	// Связи
	Orders []Order `json:"orders"` // Связь с таблицей заказов, связанными с продавцом.
}

func (Seller) TableName() string {
	return "users"
}
