// C:\GoProject\src\eShop\pkg\controllers\auth.go

package controllers

import (
	"eShop/models"
	"eShop/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
