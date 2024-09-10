// C:\GoProject\src\eShop\pkg\controllers\users.go

package controllers

import (
	"eShop/errs"
	"eShop/logger"
	"eShop/models"
	"eShop/pkg/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAllUsers получает список всех пользователей...
func GetAllUsers(c *gin.Context) {
	// Логируем IP клиента при запросе списка всех пользователей...
	logger.Info.Printf("IP: [%s] requested list of all users\n", c.ClientIP())

	// Вызываем сервис для получения списка всех пользователей...
	users, err := service.GetAllUsers()
	if err != nil {
		// Логируем ошибку при получении списка пользователей...
		logger.Error.Printf("[controllers.GetAllUsers] error getting all users: %v\n", err)
		handleError(c, err)
		return
	}

	// Логируем IP клиента при успешной выдаче списка пользователей...
	logger.Info.Printf("IP: [%s] got list of all users\n", c.ClientIP())

	// Отправляем клиенту ответ со списком пользователей...
	c.JSON(http.StatusOK, users)
}

// GetUserByID получает данные конкретного пользователя по его ID...
func GetUserByID(c *gin.Context) {
	// Извлекаем ID пользователя из параметра запроса...
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// Логируем ошибку при некорректном параметре ID...
		logger.Error.Printf("[controllers.GetUserByID] invalid user_id path parameter: %s, IP: [%s]\n", c.Param("id"), c.ClientIP())
		handleError(c, errs.ErrValidationFailed)
		return
	}

	// Логируем IP клиента и запрашиваемый ID пользователя...
	logger.Info.Printf("IP: [%s] requested user with ID: %d\n", c.ClientIP(), id)

	// Вызываем сервис для получения данных пользователя по ID...
	user, err := service.GetUserByID(uint(id))
	if err != nil {
		handleError(c, err)
		return
	}

	// Логируем успешное получение данных пользователя...
	logger.Info.Printf("IP: [%s] got user with ID: %d\n", c.ClientIP(), id)

	// Отправляем клиенту ответ с данными пользователя...
	c.JSON(http.StatusOK, user)
}

// UpdateUserByID обновляет данные пользователя по его ID...
func UpdateUserByID(c *gin.Context) {
	// Получаем роль пользователя из контекста. Если она не задана, возвращаем ошибку валидации...
	userRole := c.GetString(userRoleCtx)
	if userRole == "" {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	// Проверяем, является ли пользователь администратором. Если нет, возвращаем ошибку "Доступ запрещен"...
	if userRole != "Admin" {
		handleError(c, errs.ErrPermissionDenied)
		return
	}

	// Извлекаем ID пользователя из параметра запроса и конвертируем его в число...
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// Логируем ошибку при некорректном параметре ID и возвращаем ошибку валидации...
		logger.Error.Printf("[controllers.UpdateUserByID] invalid user_id path parameter: %s, IP: [%s]\n", c.Param("id"), c.ClientIP())
		handleError(c, errs.ErrValidationFailed)
		return
	}

	// Логируем IP клиента и ID пользователя, для которого запрашивается обновление...
	logger.Info.Printf("IP: [%s] requested to update user with ID: %d\n", c.ClientIP(), id)

	// Извлекаем данные для обновления пользователя из тела запроса. Если данные некорректны, возвращаем ошибку валидации...
	var updatedUser models.User
	if err := c.BindJSON(&updatedUser); err != nil {
		// Логируем ошибку, если данные некорректны...
		logger.Error.Printf("[controllers.UpdateUserByID] invalid user data, IP: [%s]\n", c.ClientIP())
		handleError(c, errs.ErrValidationFailed)
		return
	}

	// Вызываем сервис для обновления данных пользователя. В случае ошибки передаем её клиенту...
	user, err := service.UpdateUserByID(uint(id), updatedUser)
	if err != nil {
		handleError(c, err)
		return
	}

	// Логируем успешное обновление пользователя...
	logger.Info.Printf("IP: [%s] successfully updated user with ID: %d\n", c.ClientIP(), id)

	// Отправляем клиенту ответ с обновлёнными данными пользователя...
	c.JSON(http.StatusOK, user)
}

// CreateUser создаёт нового пользователя...
func CreateUser(c *gin.Context) {
	var user models.User

	// Привязываем JSON тело запроса к модели пользователя...
	if err := c.BindJSON(&user); err != nil {
		// Возвращаем клиенту ошибку 400, если данные в запросе некорректные...
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Вызываем сервис для создания пользователя...
	err := service.CreateUser(user)
	if err != nil {
		// Возвращаем ошибку 500, если возникли проблемы на уровне сервиса...
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Возвращаем успешный ответ с кодом 201 при успешном создании пользователя...
	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully!!!",
	})
}
