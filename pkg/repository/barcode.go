// C:\GoProject\src\eShop\pkg\repository\barcode.go

package repository

import (
	"eShop/db"
	"eShop/errs"
	"eShop/logger"
	"eShop/models"
)

// GetProductForBarcodeByID получает информацию о продукте по его ID для генерации штрих-кода.
// Эта функция не учитывает удалённые товары и не загружает связанные данные.
func GetProductForBarcodeByID(productID uint) (models.Product, error) {
	var product models.Product

	err := db.GetDBConn().First(&product, productID).Error
	if err != nil {
		logger.Error.Printf("Error fetching product by ID %d: %v", productID, err)
		return product, errs.ErrProductNotFound
	}

	return product, nil
}
