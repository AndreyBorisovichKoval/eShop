// C:\GoProject\src\eShop\db\connection.go

package db

import (
	"eShop/logger"
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbConn *gorm.DB

func ConnectToDB() error {
	connStr := securityConfig()

	logger.Info.Println("Connecting to database...")

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		logger.Error.Printf("Failed to connect to database: %v", err)
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	dbConn = db
	logger.Info.Println("Successfully connected to database!!!")
	return nil
}

func CloseDB() error {
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

func GetDB() *gorm.DB {
	return dbConn
}
