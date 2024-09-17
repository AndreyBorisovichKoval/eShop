// C:\GoProject\src\eShop\pkg\repository\product.go

package repository

import (
	"eShop/db"
	"eShop/errs"
	"eShop/logger"
	"eShop/models"
	"strings"
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
// 		logger.Error.Printf("[repository.CreateProduct] error creating product: %v\n", err)
// 		return translateError(err)
// 	}
// 	return nil
// }

// CreateProduct добавляет новый продукт в базу данных
func CreateProduct(product models.Product) error {
	if err := db.GetDBConn().Create(&product).Error; err != nil {
		// Проверка на нарушение уникальности (например, дубликат barcode)
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			logger.Warning.Printf("[repository.CreateProduct] duplicate barcode error: %v\n", err)
			return errs.ErrUniquenessViolation
		}
		logger.Error.Printf("[repository.CreateProduct] error creating product: %v\n", err)
		return translateError(err)
	}
	return nil
}

// // GetCategoryByID получает категорию по её ID (только активные)
// func GetCategoryByID(id uint) (category models.Category, err error) {
// 	err = db.GetDBConn().Where("id = ? AND is_deleted = ?", id, false).First(&category).Error
// 	if err != nil {
// 		logger.Error.Printf("[repository.GetCategoryByID] error getting category by id: %v\n", err)
// 		return category, translateError(err)
// 	}
// 	return category, nil
// }

// // GetSupplierByID получает поставщика по его ID (только активные)
// func GetSupplierByID(id uint) (supplier models.Supplier, err error) {
// 	err = db.GetDBConn().Where("id = ? AND is_deleted = ?", id, false).First(&supplier).Error
// 	if err != nil {
// 		logger.Error.Printf("[repository.GetSupplierByID] error getting supplier by id: %v\n", err)
// 		return supplier, translateError(err)
// 	}
// 	return supplier, nil
// }

// // GetAllTaxes получает все текущие налоги
// func GetAllTaxes() (taxes []models.Taxes, err error) {
// 	err = db.GetDBConn().Find(&taxes).Error
// 	if err != nil {
// 		logger.Error.Printf("[repository.GetAllTaxes] error retrieving all taxes: %v\n", err)
// 		return nil, translateError(err)
// 	}
// 	return taxes, nil
// }
