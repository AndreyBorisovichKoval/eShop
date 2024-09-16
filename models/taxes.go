// C:\GoProject\src\eShop\models\taxes.go

package models

// final_price: это значит, что налог применяется к конечной цене товара или услуги.
// profit: это значит, что налог применяется к прибыли, то есть к разнице между продажной и себестоимостью товара.

type Taxes struct {
	ID      uint    `gorm:"primaryKey" json:"id"`             // Уникальный идентификатор налога.
	Title   string  `gorm:"size:255;not null" json:"title"`   // Название налога (например: НДС, акциз).
	Rate    float64 `gorm:"not null" json:"rate"`             // Процентная ставка налога.
	ApplyTo string  `gorm:"size:50;not null" json:"apply_to"` // Применяется к: 'final_price' или 'profit'.
}

func (Taxes) TableName() string {
	return "taxes"
}
