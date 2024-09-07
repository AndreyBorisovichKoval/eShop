// C:\GoProject\src\eShop\db\connection.go

package db

import (
	"eShop/configs"
	"eShop/logger"
	"fmt"
	"os"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbConn *gorm.DB

func ConnectToDB() error {
	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		configs.AppSettings.PostgresParams.Host,
		configs.AppSettings.PostgresParams.Port,
		configs.AppSettings.PostgresParams.User,
		configs.AppSettings.PostgresParams.Database,
		os.Getenv("DB_PASSWORD"))

	logger.Info.Println("Connecting to database...")

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		logger.Error.Printf("Failed to connect to database: %v", err)
		return err
	}
	logger.Info.Println("Successfully connected to database!!!")
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
	logger.Info.Println("Database connection closed successfully...")
	return nil
}

func GetDBConn() *gorm.DB {
	return dbConn
}
