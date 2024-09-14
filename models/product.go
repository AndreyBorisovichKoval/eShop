// C:\GoProject\src\eShop\models\product.go

package models

import "time"

type Product struct {
	ID                 uint       `gorm:"primaryKey" json:"id"`                        // Уникальный идентификатор товара.
	Barcode            string     `gorm:"size:50;unique;not null" json:"barcode"`      // Штрих-код товара.
	CategoryID         uint       `gorm:"not null" json:"category_id"`                 // Внешний ключ на категорию товара.
	Title              string     `gorm:"size:255;not null" json:"title"`              // Название товара.
	SupplierID         uint       `gorm:"not null" json:"supplier_id"`                 // Внешний ключ на поставщика.
	Quantity           float64    `gorm:"not null" json:"quantity"`                    // Общее количество товара.
	Stock              float64    `gorm:"not null" json:"stock"`                       // Остаток товара на складе.
	SupplierPrice      float64    `gorm:"not null" json:"supplier_price"`              // Цена товара у поставщика.
	RetailPrice        float64    `gorm:"not null" json:"retail_price"`                // Розничная цена товара.
	TotalPrice         float64    `gorm:"not null" json:"total_price"`                 // Общая цена товара.
	IsPaidToSupplier   bool       `gorm:"default:false" json:"is_paid_to_supplier"`    // Оплачен ли товар поставщику.
	IsVATApplicable    bool       `gorm:"default:true" json:"is_vat_applicable"`       // Применяется ли НДС.
	IsExciseApplicable bool       `gorm:"default:false" json:"is_excise_applicable"`   // Применяется ли акциз.
	ExpirationDate     *time.Time `json:"expiration_date"`                             // Дата истечения срока годности.
	Discount           float64    `gorm:"default:0" json:"discount"`                   // Размер скидки на товар.
	DiscountDetails    string     `gorm:"size:255" json:"discount_details"`            // Детали скидки.
	Unit               string     `gorm:"size:50;not null" json:"unit"`                // Единица измерения товара.
	StorageLocation    string     `gorm:"size:255" json:"storage_location"`            // Место хранения на складе.
	CreatedAt          time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"` // Время создания записи.
	UpdatedAt          *time.Time `json:"updated_at"`                                  // Время последнего обновления записи.
	DeletedAt          *time.Time `json:"deleted_at"`                                  // Время удаления записи.
	IsDeleted          bool       `gorm:"default:false" json:"is_deleted"`             // Флаг удаления товара.

	// Связи
	Supplier Supplier    `gorm:"foreignKey:SupplierID" json:"supplier"` // Связь с таблицей поставщиков.
	Category Category    `gorm:"foreignKey:CategoryID" json:"category"` // Связь с таблицей категорий.
	Orders   []OrderItem `json:"orders"`                                // Связь с таблицей заказов.
}

func (Product) TableName() string {
	return "products"
}
