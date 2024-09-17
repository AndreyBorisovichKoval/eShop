// C:\GoProject\src\eShop\pkg\controllers\products.go

package controllers

import (
	"eShop/errs"
	"eShop/logger"
	"eShop/models"
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

// // AddProduct
// // @Summary Add a new product
// // @Tags products
// // @Description Add a new product with calculated taxes
// // @ID add-product
// // @Accept json
// // @Produce json
// // @Param input body models.Product true "Product data"
// // @Success 201 {string} string "Product added successfully!"
// // @Failure 400 {object} ErrorResponse "Invalid input"
// // @Failure 500 {object} ErrorResponse "Server error"
// // @Router /products [post]
// func AddProduct(c *gin.Context) {
// 	var product models.Product

// 	if err := c.BindJSON(&product); err != nil {
// 		logger.Error.Printf("[controllers.AddProduct] error binding product data: %v\n", err)
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
// 		return
// 	}

// 	if err := service.AddProduct(product); err != nil {
// 		logger.Error.Printf("[controllers.AddProduct] error adding product: %v\n", err)
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
// 		return
// 	}

// 	logger.Info.Printf("Product %s added successfully", product.Title)
// 	c.JSON(http.StatusCreated, gin.H{"message": "Product added successfully!"})
// }

// AddProduct
// @Summary Add a new product
// @Tags products
// @Description Add a new product with calculated taxes
// @ID add-product
// @Accept json
// @Produce json
// @Param input body models.Product true "Product data"
// @Success 201 {string} string "Product added successfully!"
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Failure 409 {object} ErrorResponse "Product already exists"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /products [post]
func AddProduct(c *gin.Context) {
	var product models.Product

	if err := c.BindJSON(&product); err != nil {
		logger.Error.Printf("[controllers.AddProduct] error binding product data: %v\n", err)
		handleError(c, errs.ErrValidationFailed)
		return
	}

	if err := service.AddProduct(product); err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("Product %s added successfully", product.Title)
	c.JSON(http.StatusCreated, gin.H{"message": "Product added successfully!"})
}
