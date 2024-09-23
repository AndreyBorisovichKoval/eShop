package controllers

import (
	"eShop/errs"
	"eShop/pkg/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GenerateRandomOrders
// @Summary Generate random orders for testing
// @Tags orders_test
// @Description Generates a specified number of random orders with random items for testing
// @ID generate-random-orders
// @Produce json
// @Param count path int true "Number of orders to generate"
// @Success 200 {string} string "Random orders generated"
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /orders/generate-random/{count} [post]
// GenerateRandomOrders генерирует случайные заказы
func GenerateRandomOrders(c *gin.Context) {
	countStr := c.Param("count")
	count, err := strconv.Atoi(countStr)
	if err != nil || count <= 0 {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	err = service.GenerateRandomOrders(count)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Random orders generated"})
}
