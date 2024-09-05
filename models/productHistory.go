// C:\GoProject\src\eShop\models\productHistory.go

package models

import "time"

type ProductHistory struct {
	ID                 uint       `gorm:"primaryKey" json:"id"`                        // Уникальный идентификатор записи.
	ProductID          uint       `gorm:"not null" json:"product_id"`                  // Внешний ключ на продукт.
	Title              string     `gorm:"size:255;not null" json:"title"`              // Название товара.
	SupplierID         uint       `gorm:"not null" json:"supplier_id"`                 // Внешний ключ на поставщика.
	Quantity           float64    `gorm:"not null" json:"quantity"`                    // Количество товара.
	SupplierPrice      float64    `gorm:"not null" json:"supplier_price"`              // Цена товара у поставщика.
	RetailPrice        float64    `gorm:"not null" json:"retail_price"`                // Розничная цена товара.
	IsVATApplicable    bool       `gorm:"default:true" json:"is_vat_applicable"`       // Применяется ли НДС.
	IsExciseApplicable bool       `gorm:"default:false" json:"is_excise_applicable"`   // Применяется ли акциз.
	ExpirationDate     *time.Time `json:"expiration_date"`                             // Дата истечения срока годности.
	Discount           float64    `gorm:"default:0" json:"discount"`                   // Размер скидки на товар.
	DiscountDetails    string     `gorm:"size:255" json:"discount_details"`            // Детали скидки.
	Unit               string     `gorm:"size:50;not null" json:"unit"`                // Единица измерения товара.
	CreatedAt          time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"` // Время создания записи.
	UpdatedAt          *time.Time `json:"updated_at"`                                  // Время последнего обновления записи.
	DeletedAt          *time.Time `json:"deleted_at"`                                  // Время удаления записи.
	IsDeleted          bool       `gorm:"default:false" json:"is_deleted"`             // Флаг удаления записи.

	// Связи
	Supplier Supplier `gorm:"foreignKey:SupplierID" json:"supplier"` // Связь с таблицей поставщиков.
	Product  Product  `gorm:"foreignKey:ProductID" json:"product"`   // Связь с таблицей продуктов.
}

func (ProductHistory) TableName() string {
	return "product_history"
}
