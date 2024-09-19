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

func GetReturnByID(id uint) (models.ReturnResponse, error) {
	// Получаем возврат через репозиторий
	returnProduct, err := repository.GetReturnByID(id)
	if err != nil {
		return models.ReturnResponse{}, err
	}

	// Формируем ответ
	return models.ReturnResponse{
		ID:           returnProduct.ID, // Используем ID возврата
		ProductName:  returnProduct.Product.Title,
		SupplierName: returnProduct.Product.Supplier.Title,
		Quantity:     returnProduct.Quantity,
		ReturnReason: returnProduct.ReturnReason,
		ReturnedAt:   returnProduct.ReturnedAt.Format("2006-01-02T15:04:05Z"),
	}, nil
}

func GetAllReturns() ([]models.ReturnResponse, error) {
	// Получаем список всех возвратов через репозиторий
	returnProducts, err := repository.GetAllReturns()
	if err != nil {
		return nil, err
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
			ReturnedAt:   returnProduct.ReturnedAt.Format("2006-01-02T15:04:05Z"),
		})
	}

	return returnResponses, nil
}
