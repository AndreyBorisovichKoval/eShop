// C:\GoProject\src\eShop\pkg\service\log_user_request.go

package service

import (
	"eShop/db"
	"eShop/logger"
	"eShop/models"
	"time"
)

// LogUserRequest - функция, которая по известному ID пользователя делает запрос к таблице user и записывает данные в RequestHistory
func LogUserRequest(userID uint, path string, method string, ip string) error {
	// Объект для хранения данных пользователя
	var user models.User

	// Запрос в таблицу user для извлечения данных по ID
	err := db.GetDBConn().Where("id = ?", userID).First(&user).Error
	if err != nil {
		logger.Error.Printf("Failed to fetch user with ID %d: %v", userID, err)
		return err
	}

	// Заполняем структуру RequestHistory для записи в базу данных
	requestHistory := models.RequestHistory{
		UserID:    user.ID,
		Username:  user.Username,
		FullName:  user.FullName, // Добавляем полное имя
		Email:     user.Email,    // Добавляем email
		Phone:     user.Phone,    // Добавляем телефон
		Role:      user.Role,
		Path:      path,
		Method:    method,
		IP:        ip,
		CreatedAt: time.Now(),
	}

	// Записываем данные в таблицу RequestHistory
	err = db.GetDBConn().Create(&requestHistory).Error
	if err != nil {
		logger.Error.Printf("Failed to save request history for user ID %d: %v", userID, err)
		return err
	}

	logger.Info.Printf("Request history for user ID %d successfully saved", userID)
	return nil
}

// // C:\GoProject\src\eShop\pkg\service\log_user_request.go

// package service

// import (
// 	"eShop/db"
// 	"eShop/logger"
// 	"eShop/models"
// 	"time"
// )

// // LogUserRequest - функция, которая по известному ID пользователя делает запрос к таблице user и записывает данные в RequestHistory
// func LogUserRequest(userID uint, path string, method string, ip string) error {
// 	// Объект для хранения данных пользователя
// 	var user models.User

// 	// Запрос в таблицу user для извлечения данных по ID
// 	err := db.GetDBConn().Where("id = ?", userID).First(&user).Error
// 	if err != nil {
// 		logger.Error.Printf("Failed to fetch user with ID %d: %v", userID, err)
// 		return err
// 	}

// 	// Заполняем структуру RequestHistory для записи в базу данных
// 	requestHistory := models.RequestHistory{
// 		UserID:    user.ID,
// 		Username:  user.Username,
// 		Role:      user.Role,
// 		Path:      path,
// 		Method:    method,
// 		IP:        ip,
// 		CreatedAt: time.Now(),
// 	}

// 	// Записываем данные в таблицу RequestHistory
// 	err = db.GetDBConn().Create(&requestHistory).Error
// 	if err != nil {
// 		logger.Error.Printf("Failed to save request history for user ID %d: %v", userID, err)
// 		return err
// 	}

// 	logger.Info.Printf("Request history for user ID %d successfully saved", userID)
// 	return nil
// }
