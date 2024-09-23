// C:\GoProject\src\eShop\pkg\repository\ktu.go

package repository

import (
	"eShop/db"
	"eShop/logger"
)

// SalesData структура для хранения данных о продажах пользователя
type SalesData struct {
	FullName   string  // Имя пользователя
	TotalSales float64 // Общая сумма продаж
}

// GetSalesDataByMonth возвращает данные по продажам за указанный месяц и год
func GetSalesDataByMonth(year, month int) ([]SalesData, error) {
	var salesData []SalesData

	// Выполняем запрос для получения суммы продаж каждого сотрудника
	query := `
		SELECT u.full_name, SUM(o.total_amount) as total_sales
		FROM orders o
		JOIN users u ON o.user_id = u.id
		WHERE EXTRACT(YEAR FROM o.created_at) = ? AND EXTRACT(MONTH FROM o.created_at) = ?
		GROUP BY u.full_name
	`

	err := db.GetDBConn().Raw(query, year, month).Scan(&salesData).Error
	if err != nil {
		logger.Error.Printf("[repository.GetSalesDataByMonth] Error fetching sales data: %v\n", err)
		return nil, err
	}

	logger.Info.Printf("[repository.GetSalesDataByMonth] Successfully fetched sales data for %d-%d\n", year, month) // Логирование успешного получения данных
	return salesData, nil
}

// package repository

// import (
// 	"eShop/db"
// 	"eShop/logger"
// )

// // SalesData структура для хранения данных о продажах пользователя
// type SalesData struct {
// 	FullName   string  // Имя пользователя
// 	TotalSales float64 // Общая сумма продаж
// }

// // GetSalesDataByMonth возвращает данные по продажам за указанный месяц и год
// func GetSalesDataByMonth(year, month int) ([]SalesData, error) {
// 	var salesData []SalesData

// 	// Выполняем запрос для получения суммы продаж каждого сотрудника
// 	query := `
// 		SELECT u.full_name, SUM(o.total_amount) as total_sales
// 		FROM orders o
// 		JOIN users u ON o.user_id = u.id
// 		WHERE EXTRACT(YEAR FROM o.created_at) = ? AND EXTRACT(MONTH FROM o.created_at) = ?
// 		GROUP BY u.full_name
// 	`
// 	// 	query := `
// 	//     SELECT u.full_name, u.role, ROUND(COALESCE(SUM(o.total_amount), 0), 2) as total_sales
// 	//     FROM orders o
// 	//     JOIN users u ON o.user_id = u.id
// 	//     WHERE EXTRACT(YEAR FROM o.created_at) = ? AND EXTRACT(MONTH FROM o.created_at) = ?
// 	//     GROUP BY u.full_name, u.role
// 	//     ORDER BY total_sales DESC
// 	// `

// 	err := db.GetDBConn().Raw(query, year, month).Scan(&salesData).Error
// 	if err != nil {
// 		logger.Error.Printf("[repository.GetSalesDataByMonth] Error fetching sales data: %v\n", err)
// 		return nil, err
// 	}

// 	return salesData, nil
// }
