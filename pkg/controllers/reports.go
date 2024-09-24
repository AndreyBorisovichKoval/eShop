// C:\GoProject\src\eShop\pkg\controllers\reports.go

package controllers

import (
	"eShop/errs"
	"eShop/pkg/service"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// GetReport
// @Summary Retrieve a report for a given type (sales, low-stock, seller, supplier, category-sales)
// @Tags reports
// @Description Get a report in JSON, CSV, XLSX, or ZIP format
// @ID get-report
// @Produce json, application/octet-stream, application/zip
// @Param report_type path string true "Report type (sales, low-stock, seller, supplier, category-sales)"
// @Param start_date query string false "Start date in format YYYY-MM-DD (for reports requiring a period)"
// @Param end_date query string false "End date in format YYYY-MM-DD (for reports requiring a period)"
// @Param threshold query float64 false "Minimum stock threshold (for low stock report)"
// @Param format query string false "File format (json, csv, xlsx, csvzip, or xlsxzip)"
// @Success 200 {object} interface{} "Report"
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /reports/{report_type} [get]
func GetReport(c *gin.Context) {
	reportType := c.Param("report_type")
	format := c.Query("format")

	var startDate, endDate time.Time
	var threshold float64
	var err error

	// Для отчётов, требующих диапазона дат
	if reportType == "sales" || reportType == "category-sales" {
		startDate, err = parseDateQuery(c, "start_date")
		if err != nil {
			handleError(c, errs.ErrValidationFailed)
			return
		}

		endDate, err = parseDateQuery(c, "end_date")
		if err != nil {
			handleError(c, errs.ErrValidationFailed)
			return
		}
	}

	// Для отчёта по низкому запасу
	if reportType == "low-stock" {
		threshold, err = parseThresholdQuery(c)
		if err != nil {
			handleError(c, errs.ErrValidationFailed)
			return
		}
	}

	// Выбор формата вывода (JSON или файл)
	if format == "json" || format == "" {
		report, err := service.GenerateReport(reportType, startDate, endDate, threshold)
		if err != nil {
			handleError(c, err)
			return
		}
		c.JSON(http.StatusOK, report)
		return
	}

	// Генерация файла отчёта
	fileBuffer, fileName, err := service.GenerateReportFile(reportType, startDate, endDate, threshold, format)
	if err != nil {
		handleError(c, err)
		return
	}

	// Отправляем файл в ответе
	contentType := "application/octet-stream"
	if format == "csvzip" || format == "xlsxzip" {
		contentType = "application/zip"
	}
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Data(http.StatusOK, contentType, fileBuffer.Bytes())
}

func parseDateQuery(c *gin.Context, param string) (time.Time, error) {
	dateStr := c.Query(param)
	if dateStr == "" {
		return time.Time{}, errs.ErrValidationFailed
	}
	return time.Parse("2006-01-02", dateStr)
}

func parseThresholdQuery(c *gin.Context) (float64, error) {
	thresholdStr := c.Query("threshold")
	if thresholdStr == "" {
		return 0, errs.ErrValidationFailed
	}
	return strconv.ParseFloat(thresholdStr, 64)
}
