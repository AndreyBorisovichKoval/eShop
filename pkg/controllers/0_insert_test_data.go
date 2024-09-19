// C:\GoProject\src\eShop\pkg\controllers\test_data.go

// package controllers

// import (
// 	"eShop/db"
// 	"eShop/logger"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// // InsertTestData вставляет тестовые данные в базу данных
// func InsertTestData(c *gin.Context) {
// 	// SQL-запрос для вставки данных
// 	query := `
// 		-- Вставка тестовых данных в таблицу suppliers
// 		INSERT INTO suppliers (title, email, phone, created_at, is_deleted) VALUES
// 		('ООО Манижа-СОГД', 'supplier1@example.com', '123-456-7890', CURRENT_TIMESTAMP, false),
// 		('ООО Сафар', 'supplier2@example.com', '234-567-8901', CURRENT_TIMESTAMP, false),
// 		-- (Добавить остальные строки)

// 		-- Вставка тестовых данных в таблицу categories
// 		INSERT INTO categories (title, description, created_at, is_deleted) VALUES
// 		('Фрукты', 'Категория для всех видов фруктов', CURRENT_TIMESTAMP, false),
// 		-- (Добавить остальные строки)

// 		-- Вставка тестовых данных в таблицу products
// 		INSERT INTO products (barcode, category_id, title, supplier_id, quantity, stock, supplier_price, retail_price, total_price, is_paid_to_supplier, is_vat_applicable, is_excise_applicable, unit, storage_location, created_at, is_deleted) VALUES
// 		('123456789001', 1, 'Яблоки', 1, 100.0, 80.0, 50.0, 75.0, 5000.0, false, true, false, 'кг', 'Склад 1', CURRENT_TIMESTAMP, false),
// 		-- (Добавить остальные строки)
// 	`

// 	// Выполнение SQL-запроса
// 	err := db.GetDBConn().Exec(query).Error
// 	if err != nil {
// 		logger.Error.Printf("[InsertTestData] Error inserting test data: %v", err)
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert test data"})
// 		return
// 	}

// 	logger.Info.Println("[InsertTestData] Test data inserted successfully")
// 	c.JSON(http.StatusOK, gin.H{"message": "Test data inserted successfully"})
// }

package controllers

import (
	"eShop/db"
	"eShop/logger"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// InsertTestData читает данные из SQL файла и вставляет тестовые данные в базу данных
func InsertTestData(c *gin.Context) {
	// Открываем SQL файл
	file, err := os.Open("C:/GoProject/src/eShop/db/insert_test_data.sql")
	if err != nil {
		logger.Error.Printf("[InsertTestData] Error opening SQL file: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open SQL file"})
		return
	}
	defer file.Close()

	// Читаем содержимое файла
	sqlBytes, err := ioutil.ReadAll(file)
	if err != nil {
		logger.Error.Printf("[InsertTestData] Error reading SQL file: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read SQL file"})
		return
	}

	// Преобразуем байты в строку
	sqlQuery := string(sqlBytes)

	// Выполняем SQL-запрос
	err = db.GetDBConn().Exec(sqlQuery).Error
	if err != nil {
		logger.Error.Printf("[InsertTestData] Error executing SQL query: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute SQL query"})
		return
	}

	logger.Info.Println("[InsertTestData] Test data inserted successfully")
	c.JSON(http.StatusOK, gin.H{"message": "Test data inserted successfully"})
}
