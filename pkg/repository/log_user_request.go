// C:\GoProject\src\eShop\pkg\repository\log_user_request_repository.go

package repository

import (
	"eShop/db"
	"eShop/models"
)

// LogRequestHistory - функция для записи данных в таблицу RequestHistory
func LogRequestHistory(requestHistory *models.RequestHistory) error {
	// Запись данных в таблицу RequestHistory
	err := db.GetDBConn().Create(requestHistory).Error
	if err != nil {
		return err
	}

	return nil
}
