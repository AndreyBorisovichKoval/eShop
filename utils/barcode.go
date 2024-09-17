// C:\GoProject\src\eShop\utils\barcode.go

package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// // GenerateBarcode генерирует уникальный штрих-код длиной 12 цифр
// func GenerateBarcode() (string, error) {
// 	const length = 12
// 	const charset = "0123456789"

// 	barcode := make([]byte, length)
// 	for i := range barcode {
// 		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
// 		if err != nil {
// 			return "", fmt.Errorf("failed to generate barcode: %v", err)
// 		}
// 		barcode[i] = charset[num.Int64()]
// 	}

// 	return string(barcode), nil
// }

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
