// C:\GoProject\src\eShop\pkg\controllers\auth.go

package controllers

import (
	"eShop/models"
	"eShop/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var seller models.Seller
	if err := c.BindJSON(&seller); err != nil {
		handleError(c, err)
		return
	}

	err := service.CreateSeller(seller)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully!!!"})
}

func SignIn(c *gin.Context) {
	var seller models.Seller
	if err := c.BindJSON(&seller); err != nil {
		handleError(c, err)
		return
	}

	accessToken, err := service.SignIn(seller.SellerName, seller.Password)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": accessToken})
}
