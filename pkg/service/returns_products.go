// C:\GoProject\src\eShop\pkg\service\returns_products.go

package service

import (
	"eShop/models"
	"eShop/pkg/repository"
)

// AddReturnProduct добавляет новый возврат товара в систему.
// Входной параметр: returnProduct - структура, содержащая информацию о возврате.
// Возвращает ошибку в случае неудачи.
func AddReturnProduct(returnProduct models.ReturnsProduct) error {
	// Получаем товар по его ID из репозитория
	product, err := repository.GetProductByID(returnProduct.ProductID)
	if err != nil {
		return err // Если товар не найден, возвращаем ошибку
	}

	// Устанавливаем SupplierID на основании найденного товара
	returnProduct.SupplierID = product.SupplierID

	// Добавляем возврат товара в базу данных через репозиторий
	return repository.AddReturnProduct(returnProduct)
}

// GetReturnByID получает информацию о возврате по его ID.
// Входной параметр: id - уникальный идентификатор возврата.
// Возвращает структуру ReturnResponse с информацией о возврате и ошибку в случае неудачи.
func GetReturnByID(id uint) (models.ReturnResponse, error) {
	// Получаем возврат через репозиторий по ID
	returnProduct, err := repository.GetReturnByID(id)
	if err != nil {
		return models.ReturnResponse{}, err // Возвращаем пустую структуру и ошибку в случае неудачи
	}

	// Формируем ответ с информацией о возврате
	return models.ReturnResponse{
		ID:           returnProduct.ID, // Используем ID возврата
		ProductName:  returnProduct.Product.Title,
		SupplierName: returnProduct.Product.Supplier.Title,
		Quantity:     returnProduct.Quantity,
		ReturnReason: returnProduct.ReturnReason,
		ReturnedAt:   returnProduct.CreatedAt.Format("2006-01-02T15:04:05Z"), // Используем CreatedAt как дату возврата
	}, nil
}

// GetAllReturns получает список всех возвратов в системе.
// Возвращает срез структур ReturnResponse и ошибку в случае неудачи.
func GetAllReturns() ([]models.ReturnResponse, error) {
	// Получаем список всех возвратов через репозиторий
	returnProducts, err := repository.GetAllReturns()
	if err != nil {
		return nil, err // Возвращаем nil и ошибку в случае неудачи
	}

	// Формируем список для ответа
	var returnResponses []models.ReturnResponse
	for _, returnProduct := range returnProducts {
		returnResponses = append(returnResponses, models.ReturnResponse{
			ID:           returnProduct.ID, // Используем ID возврата
			ProductName:  returnProduct.Product.Title,
			SupplierName: returnProduct.Product.Supplier.Title,
			Quantity:     returnProduct.Quantity,
			ReturnReason: returnProduct.ReturnReason,
			ReturnedAt:   returnProduct.CreatedAt.Format("2006-01-02T15:04:05Z"), // Используем CreatedAt как дату возврата
		})
	}

	return returnResponses, nil // Возвращаем сформированный список возвратов
}
