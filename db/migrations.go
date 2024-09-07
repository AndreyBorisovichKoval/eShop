// C:\GoProject\src\eShop\db\migrations.go

package db

import (
	"eShop/logger"
	"eShop/models"
)

func MigrateDB() error {
	logger.Info.Println("Starting database migration...")
	err := dbConn.AutoMigrate(
		models.Category{},
		models.Order{},
		models.OrderItem{},
		models.Payment{},
		models.Product{},
		models.ProductHistory{},
		models.ReturnsProduct{},
		models.Supplier{},
		models.Taxes{},
		models.User{},
	)
	if err != nil {
		logger.Error.Printf("Migration failed: %v", err)
		return err
	}
	logger.Info.Println("Database migration completed successfully!!!")
	return nil
}
