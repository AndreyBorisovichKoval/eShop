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
	BlockedAt *time.Time `json:"blocked_at"`                                  // Время блокировки продавца...
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"` // Время создания записи.
	// UpdatedAt *time.Time `gorm:"-" json:"updated_at"`                         // Время последнего обновления записи.
	UpdatedAt *time.Time `json:"updated_at"` // Время последнего обновления записи.
	// CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime"` // Время создания записи.
	// UpdatedAt *time.Time `json:"updated_at" gorm:"autoUpdateTime"` // Время последнего обновления записи.
	DeletedAt *time.Time `json:"deleted_at"`                      // Время удаления записи.
	IsDeleted bool       `gorm:"default:false" json:"is_deleted"` // Флаг удаления записи.

	// Связи
	Orders []Order `json:"orders"` // Связь с таблицей заказов, связанными с продавцом.
}

func (User) TableName() string {
	return "users"
}

// SwagUser представляет структуру пользователя с основными данными...
type SwagUser struct {
	FullName string `json:"full_name"` // Полное имя пользователя...
	Username string `json:"username"`  // Логин для входа в систему...
	Email    string `json:"email"`     // Адрес электронной почты пользователя...
	Password string `json:"password"`  // Пароль для аутентификации...
}

// ErrorResponse представляет структуру для обработки сообщений об ошибках...
type ErrorResponse struct {
	Error string `json:"error"` // Описание возникшей ошибки...
}

// TokenResponse представляет ответ с токеном доступа и идентификатором пользователя...
type TokenResponse struct {
	AccessToken string `json:"access_token"` // JWT токен для аутентификации пользователя...
}
