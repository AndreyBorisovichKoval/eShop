// C:\GoProject\src\eShop\db\connection.go

package db

import (
	"eShop/configs"
	"eShop/logger"
	"fmt"
	"os"
	"time"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbConn *gorm.DB

// ConnectToDB устанавливает соединение с базой данных PostgreSQL и учитывает временную зону...
func ConnectToDB() error {
	// Формируем строку подключения с указанием временной зоны Asia/Tashkent, так как для Душанбе в Windows временная зона отсутствует...
	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s TimeZone=Asia/Dushanbe",
		configs.AppSettings.PostgresParams.Host,
		configs.AppSettings.PostgresParams.Port,
		configs.AppSettings.PostgresParams.User,
		configs.AppSettings.PostgresParams.Database,
		os.Getenv("DB_PASSWORD"))

	// Логируем попытку подключения к базе данных...
	logger.Info.Println("Connecting to database...")

	// Открываем соединение с базой данных с использованием GORM...
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{NowFunc: func() time.Time {
		return time.Now().Local()
	}})
	if err != nil {
		// Логируем ошибку, если не удалось подключиться к базе данных...
		logger.Error.Printf("Failed to connect to database: %v", err)
		return err
	}

	// Логируем успешное подключение к базе данных...
	logger.Info.Println("Successfully connected to database!")

	// Сохраняем подключение для дальнейшего использования...
	dbConn = db
	return nil
}

func CloseDBConn() error {
	sqlDB, err := dbConn.DB()
	if err != nil {
		logger.Error.Printf("Failed to retrieve SQL DB from gorm.DB: %v", err)
		return err
	}
	err = sqlDB.Close()
	if err != nil {
		logger.Error.Printf("Failed to close database connection: %v", err)
		return err
	}
	logger.Info.Println("\nDatabase connection closed successfully!")
	return nil
}

func GetDBConn() *gorm.DB {
	return dbConn
}
