// C:\GoProject\src\eShop\pkg\repository\taxes.go

package repository

import (
	"eShop/db"
	"eShop/logger"
	"eShop/models"
)

// GetAllTaxes получает все текущие налоги
func GetAllTaxes() (taxes []models.Taxes, err error) {
	err = db.GetDBConn().Find(&taxes).Error
	if err != nil {
		logger.Error.Printf("Error retrieving all taxes: %v", err)
		return nil, translateError(err)
	}
	logger.Info.Printf("Taxes retrieved successfully") // Лог успешного получения налогов
	return taxes, nil
}

// GetTaxByID получает налог по его ID
func GetTaxByID(id uint) (tax models.Taxes, err error) {
	err = db.GetDBConn().First(&tax, id).Error
	if err != nil {
		logger.Error.Printf("Error retrieving tax by ID: %v", err)
		return tax, translateError(err)
	}
	logger.Info.Printf("Tax retrieved successfully with ID: %d", id) // Лог успешного получения налога по ID
	return tax, nil
}

// UpdateTax обновляет налог в базе данных
func UpdateTax(tax models.Taxes) error {
	if err := db.GetDBConn().Save(&tax).Error; err != nil {
		logger.Error.Printf("Error updating tax: %v", err)
		return translateError(err)
	}
	logger.Info.Printf("Tax updated successfully with ID: %d", tax.ID) // Лог успешного обновления налога
	return nil
}

// package repository

// import (
// 	"eShop/db"
// 	"eShop/logger"
// 	"eShop/models"
// )

// // GetAllTaxes получает все текущие налоги
// func GetAllTaxes() (taxes []models.Taxes, err error) {
// 	err = db.GetDBConn().Find(&taxes).Error
// 	if err != nil {
// 		logger.Error.Printf("Error retrieving all taxes: %v", err)
// 		return nil, translateError(err)
// 	}
// 	return taxes, nil
// }

// // GetTaxByID получает налог по его ID
// func GetTaxByID(id uint) (tax models.Taxes, err error) {
// 	err = db.GetDBConn().First(&tax, id).Error
// 	if err != nil {
// 		logger.Error.Printf("Error retrieving tax by ID: %v", err)
// 		return tax, translateError(err)
// 	}
// 	return tax, nil
// }

// // UpdateTax обновляет налог в базе данных
//
//	func UpdateTax(tax models.Taxes) error {
//		if err := db.GetDBConn().Save(&tax).Error; err != nil {
//			logger.Error.Printf("Error updating tax: %v", err)
//			return translateError(err)
//		}
//		return nil
//	}
