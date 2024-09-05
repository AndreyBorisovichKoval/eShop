// C:\GoProject\src\eShop\models\taxes.go

package models

import "time"

type Taxes struct {
	ID        uint       `gorm:"primaryKey" json:"id"`                        // Уникальный идентификатор налога.
	Title     string     `gorm:"size:255;not null" json:"title"`              // Название налога (например: НДС, акциз).
	Rate      float64    `gorm:"not null" json:"rate"`                        // Процентная ставка налога.
	ApplyTo   string     `gorm:"size:50;not null" json:"apply_to"`            // Применяется к: 'final_price' или 'profit'.
	CreatedAt time.Time  `json:"created_at"`                                  // Время создания записи.
	UpdatedAt *time.Time `json:"updated_at"`                                  // Время последнего обновления записи. Может быть пустым.
	DeletedAt *time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"deleted_at"` // Время удаления записи. По умолчанию устанавливается текущее время.
	IsDeleted bool       `gorm:"default:false" json:"is_deleted"`             // Флаг удаления записи (soft delete). По умолчанию false.
}

func (Taxes) TableName() string {
	return "taxes"
}
