// C:\GoProject\src\eShop\pkg\repository\log_user_request_repository.go

package repository

import (
	"eShop/db"
	"eShop/models"
)

// // GetUserByID - функция, которая получает пользователя по его ID
// func GetUserByID(userID uint) (*models.User, error) {
// 	var user models.User

// 	// Запрос к базе данных для получения пользователя по ID
// 	err := db.GetDBConn().Where("id = ?", userID).First(&user).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &user, nil
// }

// LogRequestHistory - функция для записи данных в таблицу RequestHistory
func LogRequestHistory(requestHistory *models.RequestHistory) error {
	// Запись данных в таблицу RequestHistory
	err := db.GetDBConn().Create(requestHistory).Error
	if err != nil {
		return err
	}

	return nil
}
