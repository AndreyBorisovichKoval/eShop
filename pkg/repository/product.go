// C:\GoProject\src\eShop\pkg\repository\product.go

package repository

import (
	"eShop/db"
	"eShop/logger"
	"eShop/models"
)

// // GetAllProducts получает все активные продукты
// func GetAllProducts() ([]models.Product, error) {
// 	var products []models.Product
// 	err := db.GetDBConn().Where("is_deleted = ?", false).Find(&products).Error
// 	if err != nil {
// 		logger.Error.Printf("[repository.GetAllProducts] error retrieving products: %v", err)
// 		return nil, err
// 	}
// 	return products, nil
// }

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
