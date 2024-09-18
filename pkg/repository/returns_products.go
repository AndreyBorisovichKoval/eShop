// C:\GoProject\src\eShop\pkg\repository\returns_products.go

package repository

import (
	"eShop/db"
	"eShop/logger"
	"eShop/models"
)

// AddReturnProduct добавляет новую запись о возврате товара в базу данных
func AddReturnProduct(returnProduct models.ReturnsProduct) error {
	if err := db.GetDBConn().Create(&returnProduct).Error; err != nil {
		logger.Error.Printf("[repository.AddReturnProduct] error adding return: %v\n", err)
		return translateError(err)
	}
	return nil
}

// GetAllReturns возвращает список всех возвратов товаров с загрузкой данных о продукте и поставщике
func GetAllReturns() ([]models.ReturnsProduct, error) {
	var returns []models.ReturnsProduct
	err := db.GetDBConn().Preload("Product").Preload("Product.Category").Preload("Supplier").Find(&returns).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllReturns] error retrieving returns: %v\n", err)
		return nil, translateError(err)
	}
	return returns, nil
}

// GetReturnByID возвращает возврат товара по ID с загрузкой данных о продукте и поставщике
func GetReturnByID(id uint) (models.ReturnsProduct, error) {
	var returnProduct models.ReturnsProduct
	err := db.GetDBConn().Preload("Product").Preload("Product.Category").Preload("Supplier").First(&returnProduct, id).Error
	if err != nil {
		logger.Error.Printf("[repository.GetReturnByID] error retrieving return by ID: %v\n", err)
		return returnProduct, translateError(err)
	}
	return returnProduct, nil
}
