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
// @Summary Добавить возврат товара
// @Tags returns
// @Description Добавить новую запись о возврате товара
// @Accept json
// @Produce json
// @Param input body models.ReturnsProduct true "Информация о возврате товара"
// @Success 201 {string} string "Возврат добавлен успешно!"
// @Failure 400 {object} ErrorResponse "Ошибка ввода"
// @Failure 500 {object} ErrorResponse "Ошибка сервера"
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
// @Summary Получить все возвраты
// @Tags returns
// @Description Получить список всех возвратов
// @Produce json
// @Success 200 {array} models.ReturnsProduct "Список возвратов"
// @Failure 500 {object} ErrorResponse "Ошибка сервера"
// @Router /returns [get]
func GetAllReturns(c *gin.Context) {
	returns, err := service.GetAllReturns()
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("Retrieved list of all returns")
	c.JSON(http.StatusOK, returns)
}

// GetReturnByID
// @Summary Получить возврат по ID
// @Tags returns
// @Description Получить информацию о возврате по ID
// @Param id path int true "Return ID"
// @Produce json
// @Success 200 {object} models.ReturnsProduct "Информация о возврате"
// @Failure 404 {object} ErrorResponse "Возврат не найден"
// @Failure 500 {object} ErrorResponse "Ошибка сервера"
// @Router /returns/{id} [get]
func GetReturnByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	returnProduct, err := service.GetReturnByID(uint(id))
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("Return with ID %d retrieved successfully", id)
	c.JSON(http.StatusOK, returnProduct)
}
