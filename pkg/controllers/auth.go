// C:\GoProject\src\eShop\pkg\controllers\auth.go

package controllers

import (
	"eShop/models"
	"eShop/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SignUp
// @Summary Регистрация пользователя
// @Tags auth, users
// @Description Создаёт новый аккаунт пользователя. Только администратор имеет право выполнять эту операцию.
// @ID create-account
// @Accept json
// @Produce json
// @Param input body models.SwagUser true "Информация о пользователе"
// @Success 201 {string} string "User created successfully!!!"
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Failure 403 {object} ErrorResponse "Permission denied. Only Admin can create users..."
// @Failure 500 {object} ErrorResponse "Server error"
// @Failure default {object} ErrorResponse
// @Router /auth/sign-up [post]
// @Security ApiKeyAuth
func SignUp(c *gin.Context) {
	// Получаем роль текущего пользователя из контекста
	userRole, exists := c.Get(userRoleCtx)

	// fmt.Println("userRole: ", userRole)
	// fmt.Println("exists: ", exists)

	if !exists || userRole != "Admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied. Only Admin can create users..."})
		return
	}

	var user models.User
	if err := c.BindJSON(&user); err != nil {
		handleError(c, err) // Используем handleError для обработки ошибки
		return
	}
	if err := service.CreateUser(user); err != nil {
		handleError(c, err) // Используем handleError для обработки ошибки
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully!!!"})
}

// SignIn
// @Summary Вход в систему
// @Tags auth
// @Description Аутентификация пользователя и получение токена доступа
// @ID sign-in
// @Accept json
// @Produce json
// @Param input body models.SwagUser true "Данные для входа (логин и пароль)"
// @Success 200 {object} map[string]string "access_token"
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Failure 401 {object} ErrorResponse "Unauthorized"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /auth/sign-in [post]
func SignIn(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		handleError(c, err)
		return
	}

	accessToken, err := service.SignIn(user.Username, user.Password)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": accessToken})
}
