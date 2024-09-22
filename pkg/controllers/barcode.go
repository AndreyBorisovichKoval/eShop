// C:\GoProject\src\eShop\pkg\controllers\barcode.go

package controllers

import (
	"eShop/errs"
	"eShop/logger"
	"eShop/pkg/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GenerateBarcode
// @Summary Generate barcode for a weighted product
// @Tags barcode
// @Description Generates a barcode based on product ID, weight, and price per unit
// @ID generate-barcode
// @Produce json
// @Param product_id query int true "Product ID"
// @Param weight query float64 true "Weight of the product"
// @Success 200 {object} map[string]string "Barcode generated"
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /barcode/generate [get]
func GenerateBarcode(c *gin.Context) {
	productIDStr := c.Query("product_id")
	weightStr := c.Query("weight")

	if productIDStr == "" || weightStr == "" {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	weight, err := strconv.ParseFloat(weightStr, 64)
	if err != nil || weight <= 0 {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	// Генерация штрих-кода через сервис
	barcode, err := service.GenerateBarcode(productID, weight)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"barcode": barcode})
}

// /

// InsertProductByBarcode
// @Summary Insert a product into the order by scanning the barcode
// @Tags barcode
// @Description Decodes a temporary barcode and inserts a product into the order based on the provided barcode
// @ID insert-product-by-barcode
// @Produce json
// @Param barcode query string true "Scanned barcode"
// @Param order_id query int true "Order ID"
// @Success 200 {object} map[string]string "Product inserted into order"
// @Failure 400 {object} ErrorResponse "Invalid barcode or order ID"
// @Failure 404 {object} ErrorResponse "Product not found"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /orders/add-from-barcode [post]
func InsertProductByBarcode(c *gin.Context) {
	barcode := c.Query("barcode")
	orderIDStr := c.Query("order_id")

	logger.Info.Printf("Received barcode: %s, order_id: %s", barcode, orderIDStr)

	if barcode == "" || orderIDStr == "" {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	orderID, err := strconv.Atoi(orderIDStr)
	if err != nil || orderID <= 0 {
		logger.Error.Printf("Invalid order_id: %v", err)
		handleError(c, errs.ErrValidationFailed)
		return
	}

	err = service.InsertProductToOrder(barcode, uint(orderID))
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product inserted into order"})
}
