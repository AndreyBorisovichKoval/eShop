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
		UserID:          user.ID,
		Username:        user.Username,
		FullName:        user.FullName,
		Email:           user.Email,
		Phone:           user.Phone,
		Role:            user.Role,
		Path:            fullURL,
		Method:          c.Request.Method,
		ClientIPAddress: c.ClientIP(),
		CreatedAt:       time.Now(),
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
