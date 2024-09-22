// C:\GoProject\src\eShop\pkg\controllers\test_data.go

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
