// C:\GoProject\src\eShop\models\user_settings.go

package models

import (
	"time"
)

// UserSettings представляет настройки пользователя.
type UserSettings struct {
	ID                    uint       `gorm:"primaryKey" json:"id"`                                          // Уникальный идентификатор настроек пользователя.
	UserID                uint       `gorm:"not null" json:"user_id"`                                       // Связь с пользователем.
	AddConfirmation       bool       `gorm:"default:true" json:"add_confirmation"`                          // Подтверждение добавления.
	UpdateConfirmation    bool       `gorm:"default:true" json:"update_confirmation"`                       // Подтверждение обновления.
	DeleteConfirmation    bool       `gorm:"default:true" json:"delete_confirmation"`                       // Подтверждение удаления.
	DisplayLanguage       string     `gorm:"size:255;default:'Russian'" json:"display_language"`            // Язык отображения.
	DesktopTheme          string     `gorm:"size:255;default:'Green animation'" json:"desktop_theme"`       // Тема рабочего стола.
	DarkModeTheme         bool       `gorm:"default:false" json:"dark_mode_theme"`                          // Темная тема.
	Font                  string     `gorm:"size:255;default:'Helvetica'" json:"font"`                      // Шрифт.
	FontSize              int        `gorm:"default:11" json:"font_size"`                                   // Размер шрифта.
	AccessibilityOptions  string     `gorm:"size:255;default:'High contrast'" json:"accessibility_options"` // Опции доступности.
	NotificationSound     bool       `gorm:"default:true" json:"notification_sound"`                        // Звук уведомлений.
	EmailNotifications    bool       `gorm:"default:false" json:"email_notifications"`                      // Email уведомления.
	NotificationFrequency string     `gorm:"size:255" json:"notification_frequency"`                        // Частота уведомлений.
	CreatedAt             time.Time  `json:"created_at"`                                                    // Дата создания настроек.
	UpdatedAt             *time.Time `json:"updated_at"`                                                    // Дата последнего обновления настроек.
}

// TableName устанавливает кастомное имя таблицы для модели UserSettings.
func (UserSettings) TableName() string {
	return "user_settings"
}
