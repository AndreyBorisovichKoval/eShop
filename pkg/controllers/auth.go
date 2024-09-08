// C:\GoProject\src\eShop\pkg\controllers\auth.go

package controllers

import (
	"eShop/models"
	"eShop/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
