// C:\GoProject\src\eShop\db\migrations.go

package db

import (
	"eShop/logger"
	"eShop/models"

	"gorm.io/gorm"
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

	// После миграции добавляем начальные записи в таблицу Taxes
	if err := addInitialTaxes(dbConn); err != nil {
		logger.Error.Printf("Failed to add initial taxes: %v", err)
		return err
	}

	logger.Info.Println("Database migration completed successfully!")
	return nil
}

// addInitialTaxes проверяет и добавляет записи в таблицу Taxes, если их нет
func addInitialTaxes(db *gorm.DB) error {
	var count int64
	db.Model(&models.Taxes{}).Count(&count)

	// Если записей нет, добавляем их
	if count == 0 {
		taxes := []models.Taxes{
			{
				Title:   "VAT",         // VAT НДС (налог на добавленную стоимость)
				Rate:    18.0,          // Процентная ставка
				ApplyTo: "final_price", // Применяется к конечной цене
			},
			{
				Title:   "Excise", // Excise Акциз
				Rate:    5.0,      // Процентная ставка
				ApplyTo: "profit", // Применяется к прибыли
			},
		}

		// Добавляем записи в таблицу
		for _, tax := range taxes {
			if err := db.Create(&tax).Error; err != nil {
				return err
			}
		}
		logger.Info.Println("Initial tax records have been added to the Taxes table.")
	} else {
		logger.Info.Println("Taxes table already has records, skipping initial insert.")
	}

	return nil
}
