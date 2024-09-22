// C:\GoProject\src\eShop\utils\barcode.go

package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
)

// GenerateBarcode генерирует уникальный штрих-код длиной 12 цифр, начинающийся с цифры "2"
func GenerateBarcode() (string, error) {
	const length = 12
	const charset = "0123456789"

	barcode := make([]byte, length)
	barcode[0] = '2' // Первый символ фиксирован — "2", для внутреннего использования

	for i := 1; i < length; i++ { // Генерация остальных символов штрих-кода
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", fmt.Errorf("failed to generate barcode: %v", err)
		}
		barcode[i] = charset[num.Int64()]
	}

	return string(barcode), nil
}

// /
// ParseBarcode анализирует штрих-код для получения данных о продукте и весе
func ParseBarcode(barcode string) (productID int, weight float64, err error) {
	if len(barcode) != 18 {
		return 0, 0, fmt.Errorf("invalid barcode length")
	}

	// Первые 6 цифр — ID продукта, следующие 5 — вес в граммах
	productID, err = strconv.Atoi(barcode[2:7])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid product ID in barcode")
	}

	// Преобразуем вес из граммов обратно в килограммы (или единицы)
	weightInGrams, err := strconv.ParseFloat(barcode[7:12], 64)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid weight in barcode")
	}

	weight = weightInGrams / 1000 // Преобразуем обратно в кг (или единицы)

	return productID, weight, nil
}
