// C:\GoProject\src\eShop\pkg\controllers\returns_products.go

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

// AddReturnProduct
// @Summary Add a product return
// @Tags returns
// @Description Add a new record for a product return
// @Accept json
// @Produce json
// @Param input body models.ReturnsProduct true "Product return information"
// @Success 201 {string} string "Return added successfully!"
// @Failure 400 {object} ErrorResponse "Input error"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /returns [post]
func AddReturnProduct(c *gin.Context) {
	var returnProduct models.ReturnsProduct
	if err := c.BindJSON(&returnProduct); err != nil {
		logger.Error.Printf("[controllers.AddReturnProduct] error binding return data: %v\n", err)
		handleError(c, errs.ErrValidationFailed)
		return
	}

	if err := service.AddReturnProduct(returnProduct); err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("Return of product %d added successfully", returnProduct.ProductID)
	c.JSON(http.StatusCreated, gin.H{"message": "Возврат добавлен успешно!"})
}

// GetAllReturns
// @Summary Get all returns
// @Tags returns
// @Description Retrieve a list of all product returns
// @Produce json
// @Success 200 {array} models.ReturnResponse "List of returns"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /returns [get]
func GetAllReturns(c *gin.Context) {
	// Получаем список всех возвратов через сервис
	returns, err := service.GetAllReturns()
	if err != nil {
		handleError(c, err)
		return
	}

	// Возвращаем список всех возвратов в нужном формате
	c.JSON(http.StatusOK, returns)
}

// GetReturnByID
// @Summary Get a return by ID
// @Tags returns
// @Description Retrieve product return information by ID
// @Param id path int true "Return ID"
// @Produce json
// @Success 200 {object} models.ReturnResponse "Return information"
// @Failure 404 {object} ErrorResponse "Return not found"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /returns/{id} [get]
func GetReturnByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	// Получаем данные возврата через сервис
	returnProduct, err := service.GetReturnByID(uint(id))
	if err != nil {
		handleError(c, err)
		return
	}

	// Возвращаем результат в нужном формате
	c.JSON(http.StatusOK, returnProduct)
}
