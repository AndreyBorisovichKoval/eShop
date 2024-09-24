// C:\GoProject\src\eShop\pkg\controllers\taxes.go

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

// GetAllTaxes
// @Summary Retrieve all current tax rates
// @Tags taxes
// @Description Get a list of all current taxes
// @ID get-all-taxes
// @Produce json
// @Success 200 {array} models.Taxes "List of taxes"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /taxes [get]
// @Security ApiKeyAuth
func GetAllTaxes(c *gin.Context) {
	logger.Info.Printf("IP: [%s] requested current tax rates", c.ClientIP())

	taxes, err := service.GetAllTaxes()
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("IP: [%s] successfully retrieved tax rates", c.ClientIP())
	c.JSON(http.StatusOK, taxes)
}

// UpdateTaxByID
// @Summary Update tax rate by ID
// @Tags taxes
// @Description Update tax rate (Admin only)
// @ID update-tax-by-id
// @Accept json
// @Produce json
// @Param id path int true "Tax ID"
// @Param input body models.Taxes true "Updated tax information"
// @Success 200 {object} models.Taxes "Updated tax"
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Failure 404 {object} ErrorResponse "Tax not found"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /taxes/{id} [patch]
// @Security ApiKeyAuth
func UpdateTaxByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	var updatedTax models.Taxes
	if err := c.BindJSON(&updatedTax); err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	logger.Info.Printf("IP: [%s] requested to update tax with ID: %d", c.ClientIP(), id)

	tax, err := service.UpdateTaxByID(uint(id), updatedTax)
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("IP: [%s] successfully updated tax with ID: %d", c.ClientIP(), id)
	c.JSON(http.StatusOK, tax)
}
