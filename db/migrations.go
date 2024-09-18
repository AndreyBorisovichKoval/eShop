// C:\GoProject\src\eShop\db\migrations.go

package db

import (
	"eShop/logger"
	"eShop/models"
	"eShop/utils"
	"os"

	"gorm.io/gorm"
)

func MigrateDB() error {
	logger.Info.Println("Starting database migration...")
	err := dbConn.AutoMigrate(
		models.Category{},
		models.Order{},
		models.OrderItem{},
		// models.Payment{},
		models.Product{},
		// models.ProductHistory{},
		models.ReturnsProduct{},
		models.Supplier{},
		models.Taxes{},
		models.User{},
		models.UserSettings{},
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

	// Добавляем начального администратора
	if err := addInitialAdmin(dbConn); err != nil {
		logger.Error.Printf("Failed to add initial admin: %v", err)
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

// addInitialAdmin проверяет и добавляет начального администратора, если его нет
func addInitialAdmin(db *gorm.DB) error {
	var count int64
	db.Model(&models.User{}).Where("role = ?", "Admin").Count(&count)

	// Если администраторов нет, добавляем начального администратора
	if count == 0 {
		// Получаем пароль администратора из переменной окружения
		adminPassword := os.Getenv("ADMIN_PASSWORD")
		// if adminPassword == "" {
		// 	adminPassword = "Admin_123" // Используем дефолтный пароль, если не задан в .env
		// }

		// Хешируем пароль
		hashedPassword := utils.GenerateHash(adminPassword)

		admin := models.User{
			FullName: "Fred Doe",
			Username: "Fred",
			Email:    "fred.doe@example.com",
			Password: hashedPassword, // Сохраняем хешированный пароль
			Role:     "Admin",
		}

		// Создаем администратора
		if err := db.Create(&admin).Error; err != nil {
			return err
		}

		// После создания администратора создаём запись в таблице user_settings
		adminSettings := models.UserSettings{
			UserID:                admin.ID, // ID уже доступен после создания пользователя
			AddConfirmation:       true,
			UpdateConfirmation:    true,
			DeleteConfirmation:    true,
			DisplayLanguage:       "Russian",
			DesktopTheme:          "Green animation",
			DarkModeTheme:         false,
			Font:                  "Helvetica",
			FontSize:              11,
			AccessibilityOptions:  "High contrast",
			NotificationSound:     true,
			EmailNotifications:    false,
			NotificationFrequency: "daily",
		}

		// Сохраняем настройки администратора
		if err := db.Create(&adminSettings).Error; err != nil {
			return err
		}

		logger.Info.Println("Initial admin 'Fred Doe' has been added to the Users table.")
	} else {
		logger.Info.Println("Admin already exists, skipping initial admin insert.")
	}

	return nil
}

// ALTER TABLE orders DROP COLUMN updated_at;
// ALTER TABLE orders DROP COLUMN deleted_at;
// ALTER TABLE orders DROP COLUMN is_deleted;

// ALTER TABLE order_items DROP COLUMN updated_at;
// ALTER TABLE order_items DROP COLUMN deleted_at;
// ALTER TABLE order_items DROP COLUMN is_deleted;
