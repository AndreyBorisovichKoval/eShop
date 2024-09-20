// C:\GoProject\src\eShop\pkg\repository\reports.go

package repository

import (
	"eShop/db"
	"eShop/logger"
	"eShop/models"
	"time"
)

// GetSalesReport получает отчет о продажах за указанный период.
// Возвращает общую сумму продаж, количество проданных товаров и топ-5 самых продаваемых продуктов.
func GetSalesReport(startDate, endDate time.Time) (models.SalesReport, error) {
	var report models.SalesReport

	// Подсчет общего количества и общей суммы продаж за указанный период
	err := db.GetDBConn().Model(&models.OrderItem{}).
		Select("SUM(order_items.total) as total_sales, SUM(order_items.quantity) as total_quantity").
		Where("order_items.created_at BETWEEN ? AND ?", startDate, endDate).
		Scan(&report).Error
	if err != nil {
		logger.Error.Printf("Error generating sales report: %v", err)
		return report, err
	}

	// Получение топ-5 самых продаваемых товаров за указанный период
	err = db.GetDBConn().Model(&models.OrderItem{}).
		Select("order_items.product_id, products.title, SUM(order_items.quantity) as quantity, SUM(order_items.total) as total").
		Joins("JOIN products ON order_items.product_id = products.id").
		Where("order_items.created_at BETWEEN ? AND ?", startDate, endDate).
		Group("order_items.product_id, products.title").
		Order("quantity DESC").
		Limit(5).
		Scan(&report.TopSelling).Error
	if err != nil {
		logger.Error.Printf("Error generating top selling products: %v", err)
		return report, err
	}

	return report, nil
}

// GetLowStockProducts получает список товаров с запасом, меньшим или равным указанному порогу.
func GetLowStockProducts(threshold float64) ([]models.LowStockReport, error) {
	var lowStockProducts []models.LowStockReport
	err := db.GetDBConn().Model(&models.Product{}).
		Select("id as product_id, title, stock").
		Where("stock <= ?", threshold).
		Scan(&lowStockProducts).Error
	if err != nil {
		logger.Error.Printf("[repository.GetLowStockProducts] error retrieving low stock products: %v\n", err)
		return nil, err
	}
	return lowStockProducts, nil
}

func GetSupplierReport() ([]models.SupplierReport, error) {
	var supplierReport []models.SupplierReport

	// Подсчет количества товаров и общей суммы поставок для каждого поставщика
	err := db.GetDBConn().Model(&models.Product{}).
		Select("supplier_id, suppliers.title as supplier_name, COUNT(products.id) as product_count, SUM(products.total_price) as total_supplies").
		Joins("JOIN suppliers ON suppliers.id = products.supplier_id").
		Group("supplier_id, suppliers.title").
		Scan(&supplierReport).Error
	if err != nil {
		logger.Error.Printf("[repository.GetSupplierReport] error generating supplier report: %v\n", err)
		return nil, err
	}

	return supplierReport, nil
}

func GetSellerReport() ([]models.SellerReport, error) {
	var report []models.SellerReport

	err := db.GetDBConn().Model(&models.Order{}).
		Select("orders.user_id AS seller_id, users.full_name AS seller_name, COUNT(orders.id) AS order_count, SUM(orders.total_amount) AS total_revenue").
		Joins("JOIN users ON users.id = orders.user_id").
		Group("orders.user_id, users.full_name").
		Scan(&report).Error
	if err != nil {
		logger.Error.Printf("[repository.GetSellerReport] error generating seller report: %v\n", err)
		return nil, err
	}

	return report, nil
}


func GetCategorySalesReport(startDate, endDate time.Time) ([]models.CategorySalesReport, error) {
	var categoryReport []models.CategorySalesReport

	// Подсчет выручки по каждой категории за указанный период
	err := db.GetDBConn().Model(&models.OrderItem{}).
		Select("products.category_id, categories.title as category_name, SUM(order_items.total) as total_sales").
		Joins("JOIN products ON products.id = order_items.product_id").
		Joins("JOIN categories ON categories.id = products.category_id").
		Where("order_items.created_at BETWEEN ? AND ?", startDate, endDate).
		Group("products.category_id, categories.title").
		Scan(&categoryReport).Error
	if err != nil {
		logger.Error.Printf("[repository.GetCategorySalesReport] error generating category sales report: %v\n", err)
		return nil, err
	}

	return categoryReport, nil
}
