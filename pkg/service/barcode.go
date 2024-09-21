// C:\GoProject\src\eShop\pkg\service\barcode.go

package service

import (
	"eShop/logger"
	"eShop/pkg/repository"
	"fmt"
)

// GenerateBarcode генерирует штрих-код для товара на основе его ID, веса и цены за единицу.
func GenerateBarcode(productID int, weight float64) (string, error) {
	// Получаем продукт по ID для генерации штрих-кода
	product, err := repository.GetProductForBarcodeByID(uint(productID))
	if err != nil {
		logger.Error.Printf("Error retrieving product data for barcode generation: %v", err)
		return "", err
	}

	// Рассчитываем итоговую стоимость товара
	totalPrice := weight * product.RetailPrice

	// Логируем рассчитанную итоговую стоимость
	logger.Info.Printf("Total price for product ID: %d with weight: %.2f is %.2f", productID, weight, totalPrice)

	// Генерация штрих-кода
	// Штрих-код включает код товара, вес и итоговую цену
	barcode := fmt.Sprintf("20%05d%05d%05d%1d", productID, int(weight*1000), int(totalPrice*100), calculateChecksum(productID, weight))

	logger.Info.Printf("Generated barcode: %s for product ID: %d with weight: %.2f and total price: %.2f", barcode, productID, weight, totalPrice)

	return barcode, nil
}

// calculateChecksum генерирует контрольную цифру для штрих-кода
func calculateChecksum(productID int, weight float64) int {
	sum := productID + int(weight*1000)
	return sum % 10 // Контрольная цифра
}
