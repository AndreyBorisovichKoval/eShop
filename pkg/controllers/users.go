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

// UpdateUserByID обновляет информацию о пользователе по его ID...
func UpdateUserByID(c *gin.Context) {
	// Получаем ID пользователя из параметра запроса...
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// Логируем и обрабатываем ошибку, если ID некорректный...
		logger.Error.Printf("[controllers.UpdateUserByID] invalid user_id path parameter: %s, IP: [%s]\n", c.Param("id"), c.ClientIP())
		handleError(c, errs.ErrValidationFailed)
		return
	}

	var updatedUser models.User
	// Парсим данные пользователя из тела запроса...
	if err := c.BindJSON(&updatedUser); err != nil {
		// Логируем и обрабатываем ошибку при некорректных данных...
		logger.Error.Printf("[controllers.UpdateUserByID] invalid request body, IP: [%s]\n", c.ClientIP())
		handleError(c, errs.ErrValidationFailed)
		return
	}

	// Логируем успешный запрос на обновление пользователя...
	logger.Info.Printf("IP: [%s] requested update for user with ID: %d\n", c.ClientIP(), id)

	// Вызываем сервис для обновления пользователя...
	updatedUser, err = service.UpdateUserByID(uint(id), updatedUser)
	if err != nil {
		// Обрабатываем ошибки через handleError...
		handleError(c, err)
		return
	}

	// Логируем успешное обновление пользователя...
	logger.Info.Printf("IP: [%s] successfully updated user with ID: %d\n", c.ClientIP(), id)
	// Отправляем ответ с обновленными данными пользователя...
	c.JSON(http.StatusOK, updatedUser)
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
