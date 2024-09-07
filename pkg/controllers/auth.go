// C:\GoProject\src\eShop\pkg\controllers\auth.go

package controllers

import (
	"eShop/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		handleError(c, err)
		return
	}

	err := service.CreateUser(user)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user created successfully"})
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
