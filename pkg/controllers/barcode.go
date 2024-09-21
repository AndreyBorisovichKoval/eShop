// C:\GoProject\src\eShop\pkg\controllers\barcode.go

package controllers

import (
	"eShop/errs"
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
