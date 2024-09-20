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

// GetSalesReport
// @Summary Retrieve sales report for a given period
// @Tags reports
// @Description Get a sales report in JSON, CSV, XLSX, or ZIP format
// @ID get-sales-report
// @Produce json, application/octet-stream, application/zip
// @Param start_date query string true "Start date in format YYYY-MM-DD"
// @Param end_date query string true "End date in format YYYY-MM-DD"
// @Param format query string false "File format (json, csv, xlsx, csv_zip, or xlsxzip)"
// @Success 200 {object} models.SalesReport "Sales report"
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /reports/sales [get]
func GetSalesReport(c *gin.Context) {
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")
	format := c.Query("format")

	if startDateStr == "" || endDateStr == "" {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	// Парсим даты
	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	// Если формат не указан или указан JSON, возвращаем данные в формате JSON
	if format == "json" || format == "" {
		report, err := service.GetSalesReport(startDate, endDate)
		if err != nil {
			handleError(c, err)
			return
		}
		c.JSON(http.StatusOK, report)
		return
	}

	// Генерация файла отчета (CSV, XLSX, ZIP)
	fileBuffer, fileName, err := service.GenerateSalesReportFile(startDate, endDate, format)
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

// GetLowStockReport
// @Summary Retrieve low stock report
// @Tags reports
// @Description Get a low stock report in JSON, CSV, XLSX, or ZIP format
// @ID get-low-stock-report
// @Produce json, application/octet-stream, application/zip
// @Param threshold query float64 true "Minimum stock threshold"
// @Param format query string false "File format (json, csv, xlsx, csv_zip, or xlsxzip)"
// @Success 200 {array} models.LowStockReport "Low stock report"
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /reports/low-stock [get]
func GetLowStockReport(c *gin.Context) {
	thresholdStr := c.Query("threshold")
	format := c.Query("format")

	if thresholdStr == "" {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	threshold, err := strconv.ParseFloat(thresholdStr, 64)
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	// Если формат не указан или указан JSON, возвращаем данные в формате JSON
	if format == "json" || format == "" {
		lowStockReport, err := service.GetLowStockReport(threshold)
		if err != nil {
			handleError(c, err)
			return
		}
		c.JSON(http.StatusOK, lowStockReport)
		return
	}

	// Генерация файла отчета (CSV, XLSX, ZIP)
	fileBuffer, fileName, err := service.GenerateLowStockReportFile(threshold, format)
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

// GetSellerReport
// @Summary Retrieve seller report
// @Tags reports
// @Description Get a seller report in JSON, CSV, XLSX, or ZIP format
// @ID get-seller-report
// @Produce json, application/octet-stream, application/zip
// @Param format query string false "File format (json, csv, xlsx, csv_zip, or xlsxzip)"
// @Success 200 {array} models.SellerReport "Seller report"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /reports/sellers [get]
func GetSellerReport(c *gin.Context) {
	format := c.Query("format")

	// Если формат не указан или указан JSON, возвращаем данные в формате JSON
	if format == "json" || format == "" {
		report, err := service.GetSellerReport()
		if err != nil {
			handleError(c, errs.ErrServerError)
			return
		}
		c.JSON(http.StatusOK, report)
		return
	}

	// Генерация файла отчета (CSV, XLSX, ZIP)
	fileBuffer, fileName, err := service.GenerateSellerReportFile(format)
	if err != nil {
		handleError(c, errs.ErrServerError)
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

// GetSupplierReport
// @Summary Retrieve supplier report
// @Tags reports
// @Description Get a supplier report with total products and total supply value
// @ID get-supplier-report
// @Produce json, application/octet-stream, application/zip
// @Param format query string false "File format (json, csv, xlsx, csv_zip, or xlsxzip)"
// @Success 200 {array} models.SupplierReport "Supplier report"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /reports/suppliers [get]
func GetSupplierReport(c *gin.Context) {
	format := c.Query("format")

	// Если формат не указан или указан JSON, возвращаем данные в формате JSON
	if format == "json" || format == "" {
		report, err := service.GetSupplierReport()
		if err != nil {
			handleError(c, err)
			return
		}
		c.JSON(http.StatusOK, report)
		return
	}

	// Генерация файла отчета (CSV, XLSX, ZIP)
	fileBuffer, fileName, err := service.GenerateSupplierReportFile(format)
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

// GetCategorySalesReport
// @Summary Retrieve category sales report for a given period
// @Tags reports
// @Description Get a category sales report in JSON, CSV, XLSX, or ZIP format
// @ID get-category-sales-report
// @Produce json, application/octet-stream, application/zip
// @Param start_date query string true "Start date in format YYYY-MM-DD"
// @Param end_date query string true "End date in format YYYY-MM-DD"
// @Param format query string false "File format (json, csv, xlsx, csvzip, or xlsxzip)"
// @Success 200 {array} models.CategorySalesReport "Category sales report"
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /reports/category-sales [get]
func GetCategorySalesReport(c *gin.Context) {
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")
	format := c.Query("format")

	if startDateStr == "" || endDateStr == "" {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	// Парсим даты
	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	// Если формат не указан или указан JSON, возвращаем данные в формате JSON
	if format == "json" || format == "" {
		report, err := service.GetCategorySalesReport(startDate, endDate)
		if err != nil {
			handleError(c, err)
			return
		}
		c.JSON(http.StatusOK, report)
		return
	}

	// Генерация файла отчета (CSV, XLSX, ZIP)
	fileBuffer, fileName, err := service.GenerateCategorySalesReportFile(startDate, endDate, format)
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
