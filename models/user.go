// C:\GoProject\src\eShop\models\user.go

package models

import "time"

type User struct {
	ID        uint       `gorm:"primaryKey" json:"id"`                        // Уникальный идентификатор продавца.
	FullName  string     `gorm:"size:255;not null" json:"full_name"`          // Полное имя продавца.
	Username  string     `gorm:"size:255;not null" json:"username"`           // Логин продавца.
	Email     string     `gorm:"size:255;not null;unique" json:"email"`       // Email продавца.
	Password  string     `gorm:"size:255;not null" json:"password"`           // Пароль продавца.
	Role      string     `gorm:"size:50;not null;default:seller" json:"role"` // Роль продавца (например: 'admin', 'manager', 'seller').
	IsBlocked bool       `gorm:"default:false" json:"is_blocked"`             // Заблокирован ли продавец.
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"` // Время создания записи.
	UpdatedAt *time.Time `json:"updated_at"`                                  // Время последнего обновления записи.
	DeletedAt *time.Time `json:"deleted_at"`                                  // Время удаления записи.
	IsDeleted bool       `gorm:"default:false" json:"is_deleted"`             // Флаг удаления записи.

	// Связи
	Orders []Order `json:"orders"` // Связь с таблицей заказов, связанными с продавцом.
}

func (User) TableName() string {
	return "users"
}
