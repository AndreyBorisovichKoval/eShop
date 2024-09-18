// C:\GoProject\src\eShop\models\report.go

package models

type SalesReport struct {
	TotalSales    float64             `json:"total_sales"`
	TotalQuantity float64             `json:"total_quantity"`
	TopSelling    []TopSellingProduct `json:"top_selling" gorm:"-"` // Игнорируем это поле в GORM
}

type TopSellingProduct struct {
	ProductID uint    `json:"product_id"` // ID товара
	Title     string  `json:"title"`      // Название товара
	Quantity  float64 `json:"quantity"`   // Количество проданного товара
	Total     float64 `json:"total"`      // Общая сумма продаж этого товара
}

type LowStockReport struct {
	ProductID uint    `json:"product_id"`
	Title     string  `json:"title"`
	Stock     float64 `json:"stock"`
}

type SellerReport struct {
	SellerID     uint    `json:"seller_id"`
	SellerName   string  `json:"seller_name"`
	OrderCount   int     `json:"order_count"`
	TotalRevenue float64 `json:"total_revenue"`
}

// /

type SupplierReport struct {
	SupplierID    uint    `json:"supplier_id"`
	SupplierName  string  `json:"supplier_name"`
	ProductCount  int     `json:"product_count"`
	TotalSupplies float64 `json:"total_supplies"`
}

type CategorySalesReport struct {
	CategoryID   uint    `json:"category_id"`
	CategoryName string  `json:"category_name"`
	TotalSales   float64 `json:"total_sales"`
}
