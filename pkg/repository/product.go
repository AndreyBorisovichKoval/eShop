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

// GetProductByID получает продукт по его ID
func GetProductByID(id uint) (models.Product, error) {
	var product models.Product
	err := db.GetDBConn().Preload("Category").Preload("Supplier").Where("id = ? AND is_deleted = ?", id, false).First(&product).Error
	if err != nil {
		logger.Error.Printf("[repository.GetProductByID] error retrieving product by id: %v\n", err)
		return product, translateError(err)
	}
	return product, nil
}

// GetProductByBarcode получает продукт по его штрих-коду
func GetProductByBarcode(barcode string) (models.Product, error) {
	var product models.Product
	err := db.GetDBConn().Preload("Category").Preload("Supplier").Where("barcode = ? AND is_deleted = ?", barcode, false).First(&product).Error
	if err != nil {
		logger.Error.Printf("[repository.GetProductByBarcode] error retrieving product by barcode: %v\n", err)
		return product, translateError(err)
	}
	return product, nil
}

// UpdateProduct обновляет существующий продукт в базе данных
func UpdateProduct(product models.Product) error {
	if err := db.GetDBConn().Save(&product).Error; err != nil {
		logger.Error.Printf("[repository.UpdateProduct] error updating product: %v\n", err)
		return translateError(err)
	}
	return nil
}
