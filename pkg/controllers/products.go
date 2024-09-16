// C:\GoProject\src\eShop\pkg\controllers\products.go

package controllers

import (
	"eShop/logger"
	"eShop/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllProducts
// @Summary Retrieve all products
// @Tags products
// @Description Get a list of all active products (Admin/Manager only)
// @ID get-all-products
// @Produce json
// @Success 200 {array} models.Product "List of active products"
// @Failure 403 {object} ErrorResponse "Permission denied"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /products [get]
// @Security ApiKeyAuth
func GetAllProducts(c *gin.Context) {
	logger.Info.Printf("IP: [%s] requested list of all products\n", c.ClientIP())

	products, err := service.GetAllProducts()
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("IP: [%s] successfully retrieved list of products\n", c.ClientIP())
	c.JSON(http.StatusOK, products)
}
