// C:\GoProject\src\eShop\pkg\service\returns_products.go

package service

import (
	"eShop/models"
	"eShop/pkg/repository"
)

// AddReturnProduct добавляет новый возврат товара
func AddReturnProduct(returnProduct models.ReturnsProduct) error {
	return repository.AddReturnProduct(returnProduct)
}

// GetAllReturns возвращает список всех возвратов товаров
func GetAllReturns() ([]models.ReturnsProduct, error) {
	return repository.GetAllReturns()
}

// GetReturnByID возвращает информацию о возврате товара по ID
func GetReturnByID(id uint) (models.ReturnsProduct, error) {
	return repository.GetReturnByID(id)
}
