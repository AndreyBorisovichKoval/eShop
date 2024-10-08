// C:\GoProject\src\eShop\pkg\repository\reports.go

package repository

import (
	"eShop/db"
	"eShop/logger"
	"eShop/models"
	"time"
)

// GetSalesReport получает отчет о продажах за указанный период.
func GetSalesReport(startDate, endDate time.Time) (models.SalesReport, error) {
	var report models.SalesReport

	err := db.GetDBConn().Model(&models.OrderItem{}).
		Select("SUM(order_items.total) as total_sales, SUM(order_items.quantity) as total_quantity").
		Where("order_items.created_at BETWEEN ? AND ?", startDate, endDate).
		Scan(&report).Error
	if err != nil {
		logger.Error.Printf("Error generating sales report: %v", err)
		return report, err
	}

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

	logger.Info.Printf("Sales report generated successfully for period: %v - %v", startDate, endDate) // Лог успешного создания отчета
	return report, nil
}

// GetLowStockProducts получает список товаров с низким запасом.
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

	logger.Info.Printf("Low stock products retrieved successfully with threshold: %f", threshold) // Лог успешного получения товаров с низким запасом
	return lowStockProducts, nil
}

// GetSupplierReport возвращает отчет по поставщикам.
func GetSupplierReport() ([]models.SupplierReport, error) {
	var supplierReport []models.SupplierReport

	err := db.GetDBConn().Model(&models.Product{}).
		Select("supplier_id, suppliers.title as supplier_name, COUNT(products.id) as product_count, SUM(products.total_price) as total_supplies").
		Joins("JOIN suppliers ON suppliers.id = products.supplier_id").
		Group("supplier_id, suppliers.title").
		Scan(&supplierReport).Error
	if err != nil {
		logger.Error.Printf("[repository.GetSupplierReport] error generating supplier report: %v\n", err)
		return nil, err
	}

	logger.Info.Printf("Supplier report generated successfully") // Лог успешного создания отчета по поставщикам
	return supplierReport, nil
}

// GetSellerReport возвращает отчет по продавцам.
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

	logger.Info.Printf("Seller report generated successfully") // Лог успешного создания отчета по продавцам
	return report, nil
}

// GetCategorySalesReport возвращает отчет по категориям товаров.
func GetCategorySalesReport(startDate, endDate time.Time) ([]models.CategorySalesReport, error) {
	var categoryReport []models.CategorySalesReport

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

	logger.Info.Printf("Category sales report generated successfully for period: %v - %v", startDate, endDate) // Лог успешного создания отчета по категориям
	return categoryReport, nil
}

