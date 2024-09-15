// C:\GoProject\src\eShop\pkg\controllers\auth.go

package controllers

import (
	"eShop/models"
	"eShop/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateUser
// @Summary Register a new user
// @Tags users
// @Description Register a new user (only Admin can do this)
// @ID create-user
// @Accept json
// @Produce json
// @Param input body models.SwagUser true "User Information"
// @Success 201 {string} string "User created successfully!!!"
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Failure 403 {object} ErrorResponse "Permission denied"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /users [post]
// @Security ApiKeyAuth
func CreateUser(c *gin.Context) {
	// Получаем роль текущего пользователя из контекста
	userRole, exists := c.Get(userRoleCtx)

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
// @Summary Log in user
// @Tags auth
// @Description User authentication (returns JWT token)
// @ID sign-in
// @Accept json
// @Produce json
// @Param input body models.SignInInput true "Login data"
// @Success 200 {object} accessTokenResponse "access_token"
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

	// Проверяем, нужно ли сбросить пароль...
	if user.PasswordResetRequired {
		c.JSON(http.StatusOK, gin.H{"message": "Password reset required", "reset_password": true})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": accessToken})
}
