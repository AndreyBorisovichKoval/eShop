// // C:\GoProject\src\eShop\pkg\repository\reports.go

// package repository

// import (
// 	"eShop/db"
// 	"eShop/logger"
// 	"eShop/models"
// 	"time"
// )

// // GetSalesReport получает отчет о продажах за указанный период.
// // Возвращает общую сумму продаж, количество проданных товаров и топ-5 самых продаваемых продуктов.
// func GetSalesReport(startDate, endDate time.Time) (models.SalesReport, error) {
// 	var report models.SalesReport

// 	// Подсчет общего количества и общей суммы продаж за указанный период
// 	err := db.GetDBConn().Model(&models.OrderItem{}).
// 		Select("SUM(order_items.total) as total_sales, SUM(order_items.quantity) as total_quantity").
// 		Where("order_items.created_at BETWEEN ? AND ?", startDate, endDate).
// 		Scan(&report).Error
// 	if err != nil {
// 		logger.Error.Printf("Error generating sales report: %v", err)
// 		return report, err
// 	}

// 	// Получение топ-5 самых продаваемых товаров за указанный период
// 	err = db.GetDBConn().Model(&models.OrderItem{}).
// 		Select("order_items.product_id, products.title, SUM(order_items.quantity) as quantity, SUM(order_items.total) as total").
// 		Joins("JOIN products ON order_items.product_id = products.id").
// 		Where("order_items.created_at BETWEEN ? AND ?", startDate, endDate).
// 		Group("order_items.product_id, products.title").
// 		Order("quantity DESC").
// 		Limit(5).
// 		Scan(&report.TopSelling).Error
// 	if err != nil {
// 		logger.Error.Printf("Error generating top selling products: %v", err)
// 		return report, err
// 	}

// 	return report, nil
// }

// // GetLowStockProducts получает список товаров с запасом, меньшим или равным указанному порогу.
// func GetLowStockProducts(threshold float64) ([]models.LowStockReport, error) {
// 	var lowStockProducts []models.LowStockReport
// 	err := db.GetDBConn().Model(&models.Product{}).
// 		Select("id as product_id, title, stock").
// 		Where("stock <= ?", threshold).
// 		Scan(&lowStockProducts).Error
// 	if err != nil {
// 		logger.Error.Printf("[repository.GetLowStockProducts] error retrieving low stock products: %v\n", err)
// 		return nil, err
// 	}
// 	return lowStockProducts, nil
// }

// func GetSupplierReport() ([]models.SupplierReport, error) {
// 	var supplierReport []models.SupplierReport

// 	// Подсчет количества товаров и общей суммы поставок для каждого поставщика
// 	err := db.GetDBConn().Model(&models.Product{}).
// 		Select("supplier_id, suppliers.title as supplier_name, COUNT(products.id) as product_count, SUM(products.total_price) as total_supplies").
// 		Joins("JOIN suppliers ON suppliers.id = products.supplier_id").
// 		Group("supplier_id, suppliers.title").
// 		Scan(&supplierReport).Error
// 	if err != nil {
// 		logger.Error.Printf("[repository.GetSupplierReport] error generating supplier report: %v\n", err)
// 		return nil, err
// 	}

// 	return supplierReport, nil
// }

// func GetSellerReport() ([]models.SellerReport, error) {
// 	var report []models.SellerReport

// 	err := db.GetDBConn().Model(&models.Order{}).
// 		Select("orders.user_id AS seller_id, users.full_name AS seller_name, COUNT(orders.id) AS order_count, SUM(orders.total_amount) AS total_revenue").
// 		Joins("JOIN users ON users.id = orders.user_id").
// 		Group("orders.user_id, users.full_name").
// 		Scan(&report).Error
// 	if err != nil {
// 		logger.Error.Printf("[repository.GetSellerReport] error generating seller report: %v\n", err)
// 		return nil, err
// 	}

// 	return report, nil
// }

// func GetCategorySalesReport(startDate, endDate time.Time) ([]models.CategorySalesReport, error) {
// 	var categoryReport []models.CategorySalesReport

// 	// Подсчет выручки по каждой категории за указанный период
// 	err := db.GetDBConn().Model(&models.OrderItem{}).
// 		Select("products.category_id, categories.title as category_name, SUM(order_items.total) as total_sales").
// 		Joins("JOIN products ON products.id = order_items.product_id").
// 		Joins("JOIN categories ON categories.id = products.category_id").
// 		Where("order_items.created_at BETWEEN ? AND ?", startDate, endDate).
// 		Group("products.category_id, categories.title").
// 		Scan(&categoryReport).Error
// 	if err != nil {
// 		logger.Error.Printf("[repository.GetCategorySalesReport] error generating category sales report: %v\n", err)
// 		return nil, err
// 	}

// 	return categoryReport, nil
// }

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
		Select("SUM(order_items.total) as total_sales, SUM(order_items.quantity) as total_quantity"). // Вычисляем общую сумму продаж и количество проданных товаров
		Where("order_items.created_at BETWEEN ? AND ?", startDate, endDate).                          // Указываем период времени для выборки
		Scan(&report).Error                                                                           // Результаты сохраняем в структуру report
	if err != nil {
		// Логируем ошибку в случае неудачи
		logger.Error.Printf("Error generating sales report: %v", err)
		return report, err
	}

	// Получение топ-5 самых продаваемых товаров за указанный период
	err = db.GetDBConn().Model(&models.OrderItem{}).
		Select("order_items.product_id, products.title, SUM(order_items.quantity) as quantity, SUM(order_items.total) as total"). // Получаем данные о самых продаваемых продуктах
		Joins("JOIN products ON order_items.product_id = products.id").                                                           // Присоединяем таблицу продуктов для получения названий
		Where("order_items.created_at BETWEEN ? AND ?", startDate, endDate).                                                      // Ограничиваем выборку временным периодом
		Group("order_items.product_id, products.title").                                                                          // Группируем по идентификатору и названию продукта
		Order("quantity DESC").                                                                                                   // Сортируем по количеству в порядке убывания
		Limit(5).                                                                                                                 // Ограничиваем результат пятью продуктами
		Scan(&report.TopSelling).Error                                                                                            // Результаты сохраняем в поле TopSelling структуры report
	if err != nil {
		logger.Error.Printf("Error generating top selling products: %v", err)
		return report, err
	}

	// Возвращаем полный отчет по продажам
	return report, nil
}

// GetLowStockProducts получает список товаров с запасом, меньшим или равным указанному порогу (threshold).
func GetLowStockProducts(threshold float64) ([]models.LowStockReport, error) {
	var lowStockProducts []models.LowStockReport

	// Выбираем товары, у которых количество на складе меньше или равно заданному порогу
	err := db.GetDBConn().Model(&models.Product{}).
		Select("id as product_id, title, stock"). // Выбираем ID продукта, название и текущее количество на складе
		Where("stock <= ?", threshold).           // Устанавливаем условие по количеству товаров
		Scan(&lowStockProducts).Error             // Результаты сохраняем в массив lowStockProducts
	if err != nil {
		// Логируем ошибку
		logger.Error.Printf("[repository.GetLowStockProducts] error retrieving low stock products: %v\n", err)
		return nil, err
	}

	// Возвращаем список товаров с низким запасом
	return lowStockProducts, nil
}

// GetSupplierReport возвращает отчет по поставщикам с указанием количества товаров и общей стоимости поставок.
func GetSupplierReport() ([]models.SupplierReport, error) {
	var supplierReport []models.SupplierReport

	// Подсчитываем количество товаров и общую стоимость поставок для каждого поставщика
	err := db.GetDBConn().Model(&models.Product{}).
		Select("supplier_id, suppliers.title as supplier_name, COUNT(products.id) as product_count, SUM(products.total_price) as total_supplies"). // Собираем данные по поставщикам
		Joins("JOIN suppliers ON suppliers.id = products.supplier_id").                                                                            // Присоединяем таблицу поставщиков для получения названий
		Group("supplier_id, suppliers.title").                                                                                                     // Группируем по поставщикам
		Scan(&supplierReport).Error                                                                                                                // Результаты сохраняем в массив supplierReport
	if err != nil {
		// Логируем ошибку
		logger.Error.Printf("[repository.GetSupplierReport] error generating supplier report: %v\n", err)
		return nil, err
	}

	// Возвращаем отчет по поставщикам
	return supplierReport, nil
}

// GetSellerReport возвращает отчет по продавцам с указанием количества заказов и общей выручки каждого продавца.
func GetSellerReport() ([]models.SellerReport, error) {
	var report []models.SellerReport

	// Собираем информацию о количестве заказов и общей выручке для каждого продавца
	err := db.GetDBConn().Model(&models.Order{}).
		Select("orders.user_id AS seller_id, users.full_name AS seller_name, COUNT(orders.id) AS order_count, SUM(orders.total_amount) AS total_revenue"). // Собираем данные о продавцах
		Joins("JOIN users ON users.id = orders.user_id").                                                                                                  // Присоединяем таблицу пользователей для получения имен продавцов
		Group("orders.user_id, users.full_name").                                                                                                          // Группируем по продавцам
		Scan(&report).Error                                                                                                                                // Результаты сохраняем в массив report
	if err != nil {
		// Логируем ошибку
		logger.Error.Printf("[repository.GetSellerReport] error generating seller report: %v\n", err)
		return nil, err
	}

	// Возвращаем отчет по продавцам
	return report, nil
}

// GetCategorySalesReport возвращает отчет по категориям товаров за указанный период.
// В отчете указывается общая выручка по каждой категории товаров.
func GetCategorySalesReport(startDate, endDate time.Time) ([]models.CategorySalesReport, error) {
	var categoryReport []models.CategorySalesReport

	// Подсчитываем выручку по каждой категории товаров за указанный период
	err := db.GetDBConn().Model(&models.OrderItem{}).
		Select("products.category_id, categories.title as category_name, SUM(order_items.total) as total_sales"). // Собираем данные по категориям товаров
		Joins("JOIN products ON products.id = order_items.product_id").                                           // Присоединяем таблицу продуктов
		Joins("JOIN categories ON categories.id = products.category_id").                                         // Присоединяем таблицу категорий
		Where("order_items.created_at BETWEEN ? AND ?", startDate, endDate).                                      // Ограничиваем выборку временным диапазоном
		Group("products.category_id, categories.title").                                                          // Группируем по категориям
		Scan(&categoryReport).Error                                                                               // Результаты сохраняем в массив categoryReport
	if err != nil {
		// Логируем ошибку
		logger.Error.Printf("[repository.GetCategorySalesReport] error generating category sales report: %v\n", err)
		return nil, err
	}

	// Возвращаем отчет по категориям товаров
	return categoryReport, nil
}
