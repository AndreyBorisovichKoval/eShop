// C:\GoProject\src\eShop\pkg\service\log_user_request.go

package service

import (
	"eShop/logger"
	"eShop/models"
	"eShop/pkg/repository"
	"time"

	"github.com/gin-gonic/gin"
)

// LogUserRequest - сервисная функция для логирования запроса пользователя
func LogUserRequest(c *gin.Context, userID uint) error {
	// Получаем пользователя из репозитория
	user, err := repository.GetUserByID(userID)
	if err != nil {
		logger.Error.Printf("Failed to fetch user with ID %d: %v", userID, err)
		return err
	}

	// Получаем полный URL вместе с query-параметрами
	fullURL := c.Request.URL.Path + "?" + c.Request.URL.RawQuery

	// Заполняем структуру RequestHistory для записи в базу данных
	requestHistory := &models.RequestHistory{
		UserID:    user.ID,
		Username:  user.Username,
		FullName:  user.FullName,
		Email:     user.Email,
		Phone:     user.Phone,
		Role:      user.Role,
		Path:      fullURL, // Сохраняем полный URL с query-параметрами
		Method:    c.Request.Method,
		IP:        c.ClientIP(),
		CreatedAt: time.Now(),
	}

	// Логируем запрос в базу через репозиторий
	err = repository.LogRequestHistory(requestHistory)
	if err != nil {
		logger.Error.Printf("Failed to log request history for user ID %d: %v", userID, err)
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

// 	"github.com/gin-gonic/gin"
// )

// // LogUserRequest - функция, которая по известному ID пользователя делает запрос к таблице user и записывает данные в RequestHistory
// func LogUserRequest(c *gin.Context, userID uint) error {
// 	// Объект для хранения данных пользователя
// 	var user models.User

// 	// Запрос в таблицу user для извлечения данных по ID
// 	err := db.GetDBConn().Where("id = ?", userID).First(&user).Error
// 	if err != nil {
// 		logger.Error.Printf("Failed to fetch user with ID %d: %v", userID, err)
// 		return err
// 	}

// 	// Получаем полный URL вместе с query-параметрами
// 	fullURL := c.Request.URL.Path + "?" + c.Request.URL.RawQuery

// 	// Заполняем структуру RequestHistory для записи в базу данных
// 	requestHistory := models.RequestHistory{
// 		UserID:    user.ID,
// 		Username:  user.Username,
// 		FullName:  user.FullName,
// 		Email:     user.Email,
// 		Phone:     user.Phone,
// 		Role:      user.Role,
// 		Path:      fullURL, // Теперь сохраняем полный URL с query-параметрами
// 		Method:    c.Request.Method,
// 		IP:        c.ClientIP(),
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
