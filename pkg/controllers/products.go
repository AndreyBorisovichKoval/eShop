// C:\GoProject\src\eShop\pkg\controllers\products.go

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

// GetAllProducts
// @Summary Retrieve all products
// @Tags products
// @Description Get a list of all active products
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

// AddProduct
// @Summary Add a new product
// @Tags products
// @Description Add a new product with calculated taxes  (Admin/Manager only)
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
	// Проверка роли пользователя
	userRole := c.GetString(userRoleCtx)
	if userRole != "Admin" && userRole != "Manager" {
		handleError(c, errs.ErrPermissionDenied)
		return
	}

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

// GetProductByID
// @Summary Retrieve a product by ID
// @Tags products
// @Description Get product information by ID
// @ID get-product-by-id
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} models.Product "Product information"
// @Failure 403 {object} ErrorResponse "Permission denied"
// @Failure 404 {object} ErrorResponse "Product not found"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /products/{id} [get]
func GetProductByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	logger.Info.Printf("IP: [%s] requested product with ID: %d\n", c.ClientIP(), id)

	product, err := service.GetProductByID(uint(id))
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("Product with ID %d retrieved successfully", id)
	c.JSON(http.StatusOK, product)
}

// GetProductByBarcode
// @Summary Retrieve a product by barcode
// @Tags products
// @Description Get product information by barcode
// @ID get-product-by-barcode
// @Produce json
// @Param barcode path string true "Product barcode"
// @Success 200 {object} models.Product "Product information"
// @Failure 403 {object} ErrorResponse "Permission denied"
// @Failure 404 {object} ErrorResponse "Product not found"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /products/barcode/{barcode} [get]
func GetProductByBarcode(c *gin.Context) {
	barcode := c.Param("barcode")

	logger.Info.Printf("IP: [%s] requested product with barcode: %s\n", c.ClientIP(), barcode)

	product, err := service.GetProductByBarcode(barcode)
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("Product with barcode %s retrieved successfully", barcode)
	c.JSON(http.StatusOK, product)
}

// UpdateProductByID
// @Summary Update a product by ID
// @Tags products
// @Description Update a product's information by ID (Admin/Manager only)
// @ID update-product-by-id
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param input body models.Product true "Updated product information"
// @Success 200 {object} models.Product "Updated product"
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Failure 404 {object} ErrorResponse "Product not found"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /products/{id} [patch]
func UpdateProductByID(c *gin.Context) {
	// Проверка роли пользователя
	userRole := c.GetString(userRoleCtx)
	if userRole != "Admin" && userRole != "Manager" {
		handleError(c, errs.ErrPermissionDenied)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	var updatedProduct models.Product
	if err := c.BindJSON(&updatedProduct); err != nil {
		logger.Error.Printf("[controllers.UpdateProductByID] error binding product data: %v\n", err)
		handleError(c, errs.ErrValidationFailed)
		return
	}

	product, err := service.UpdateProductByID(uint(id), updatedProduct)
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("Product with ID %d updated successfully", id)
	c.JSON(http.StatusOK, product)
}

// SoftDeleteProductByID
// @Summary Soft delete a product by ID
// @Tags products
// @Description Soft delete a product by ID (Admin/Manager only)
// @ID soft-delete-product-by-id
// @Param id path int true "Product ID"
// @Success 200 {string} string "Product soft deleted successfully!"
// @Failure 403 {object} ErrorResponse "Permission denied"
// @Failure 404 {object} ErrorResponse "Product not found"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /products/{id} [delete]
func SoftDeleteProductByID(c *gin.Context) {
	// Проверка роли пользователя
	userRole := c.GetString(userRoleCtx)
	if userRole != "Admin" && userRole != "Manager" {
		handleError(c, errs.ErrPermissionDenied)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	err = service.SoftDeleteProductByID(uint(id))
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("Product with ID %d soft deleted successfully", id)
	c.JSON(http.StatusOK, gin.H{"message": "Product soft deleted successfully!"})
}

// RestoreProductByID
// @Summary Restore a soft deleted product by ID
// @Tags products
// @Description Restore a soft deleted product by ID (Admin/Manager only)
// @ID restore-product-by-id
// @Param id path int true "Product ID"
// @Success 200 {string} string "Product restored successfully!"
// @Failure 403 {object} ErrorResponse "Permission denied"
// @Failure 404 {object} ErrorResponse "Product not found"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /products/{id}/restore [put]
func RestoreProductByID(c *gin.Context) {
	// Проверка роли пользователя
	userRole := c.GetString(userRoleCtx)
	if userRole != "Admin" && userRole != "Manager" {
		handleError(c, errs.ErrPermissionDenied)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	err = service.RestoreProductByID(uint(id))
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("Product with ID %d restored successfully", id)
	c.JSON(http.StatusOK, gin.H{"message": "Product restored successfully!"})
}
