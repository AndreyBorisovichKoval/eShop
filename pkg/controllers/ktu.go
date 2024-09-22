package controllers

import (
	"eShop/errs"
	"eShop/logger"
	"eShop/pkg/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetKTU
// @Summary Calculate KTU for all employees
// @Tags ktu
// @Description Calculates KTU for each employee based on monthly sales (admin, managers, sellers)
// @ID get-ktu
// @Produce json
// @Param year query int true "Year for KTU calculation"
// @Param month query int true "Month for KTU calculation"
// @Success 200 {object} map[string]float64 "KTU calculated"
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /ktu [get]
func GetKTU(c *gin.Context) {
	// Получаем параметры запроса
	yearStr := c.Query("year")
	monthStr := c.Query("month")

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	month, err := strconv.Atoi(monthStr)
	if err != nil || month < 1 || month > 12 {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	logger.Info.Printf("Calculating KTU for year: %d, month: %d\n", year, month)

	// Вызываем сервис для расчета КТУ
	ktuResult, err := service.CalculateKTU(year, month)
	if err != nil {
		handleError(c, err)
		return
	}

	// Возвращаем результат
	c.JSON(http.StatusOK, ktuResult)
}
