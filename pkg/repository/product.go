// C:\GoProject\src\eShop\pkg\repository\product.go

package repository

import (
	"eShop/db"
	"eShop/logger"
	"eShop/models"
)

// GetAllProducts получает все активные продукты с их категориями и поставщиками
func GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	err := db.GetDBConn().Preload("Category").Preload("Supplier").Where("is_deleted = ?", false).Find(&products).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllProducts] error retrieving products: %v", err)
		return nil, err
	}
	return products, nil
}

// // CreateProduct добавляет новый продукт в базу данных
// func CreateProduct(product models.Product) error {
// 	if err := db.GetDBConn().Create(&product).Error; err != nil {
// 		// Проверка на нарушение уникальности (например, дубликат barcode)
// 		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
// 			logger.Warning.Printf("[repository.CreateProduct] duplicate barcode error: %v\n", err)
// 			return errs.ErrUniquenessViolation
// 		}
// 		logger.Error.Printf("[repository.CreateProduct] error creating product: %v\n", err)
// 		return translateError(err)
// 	}
// 	return nil
// }

// CreateProduct добавляет новый продукт в базу данных
func CreateProduct(product models.Product) error {
	if err := db.GetDBConn().Create(&product).Error; err != nil {
		logger.Error.Printf("[repository.CreateProduct] error creating product: %v\n", err)
		return translateError(err)
	}
	return nil
}

// CheckBarcodeExists проверяет, существует ли штрих-код в базе данных
func CheckBarcodeExists(barcode string) (bool, error) {
	var count int64
	err := db.GetDBConn().Model(&models.Product{}).Where("barcode = ?", barcode).Count(&count).Error
	if err != nil {
		logger.Error.Printf("[repository.CheckBarcodeExists] error checking barcode: %v\n", err)
		return false, err
	}
	return count > 0, nil
}
