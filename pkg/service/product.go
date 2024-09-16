// C:\GoProject\src\eShop\pkg\service\product.go

package service

import (
	"eShop/logger"
	"eShop/models"
	"eShop/pkg/repository"
)

// GetAllProducts получает все активные продукты
func GetAllProducts() ([]models.Product, error) {
	products, err := repository.GetAllProducts()
	if err != nil {
		logger.Error.Printf("[service.GetAllProducts] error retrieving products: %v", err)
		return nil, err
	}

	return products, nil
}
