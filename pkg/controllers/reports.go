// C:\GoProject\src\eShop\pkg\controllers\reports.go

package controllers

import (
	"eShop/errs"
	"eShop/logger"
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
// @Param format query string false "File format (json, csv, xlsx, csv_zip, or xlsx_zip)"
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
	if format == "csv_zip" || format == "xlsx_zip" {
		contentType = "application/zip"
	}
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Data(http.StatusOK, contentType, fileBuffer.Bytes())
}

// // GetSalesReport
// // @Summary Retrieve sales report for a given period
// // @Tags reports
// // @Description Get a sales report in JSON, CSV, or XLSX format
// // @ID get-sales-report
// // @Produce json
// // @Param start_date query string true "Start date in format YYYY-MM-DD"
// // @Param end_date query string true "End date in format YYYY-MM-DD"
// // @Param format query string false "File format (json, csv, or xlsx)"
// // @Success 200 {object} models.SalesReport "Sales report"
// // @Failure 400 {object} ErrorResponse "Invalid input"
// // @Failure 500 {object} ErrorResponse "Server error"
// // @Router /reports/sales [get]
// func GetSalesReport(c *gin.Context) {
// 	startDateStr := c.Query("start_date")
// 	endDateStr := c.Query("end_date")
// 	format := c.Query("format")

// 	if startDateStr == "" || endDateStr == "" {
// 		handleError(c, errs.ErrValidationFailed)
// 		return
// 	}

// 	// Парсим даты
// 	startDate, err := time.Parse("2006-01-02", startDateStr)
// 	if err != nil {
// 		logger.Error.Printf("[controllers.GetSalesReport] error parsing start_date: %v\n", err)
// 		handleError(c, errs.ErrValidationFailed)
// 		return
// 	}

// 	endDate, err := time.Parse("2006-01-02", endDateStr)
// 	if err != nil {
// 		logger.Error.Printf("[controllers.GetSalesReport] error parsing end_date: %v\n", err)
// 		handleError(c, errs.ErrValidationFailed)
// 		return
// 	}

// 	// Генерация отчета в зависимости от формата
// 	if format == "json" || format == "" {
// 		// Возвращаем отчет в формате JSON
// 		report, err := service.GetSalesReport(startDate, endDate)
// 		if err != nil {
// 			handleError(c, err)
// 			return
// 		}
// 		c.JSON(http.StatusOK, report)
// 		return
// 	}

// 	// Генерация файла отчета (CSV или XLSX)
// 	fileBuffer, fileName, err := service.GenerateSalesReportFile(startDate, endDate, format)
// 	if err != nil {
// 		handleError(c, err)
// 		return
// 	}

// 	// Отправляем файл в ответе
// 	c.Header("Content-Disposition", "attachment; filename="+fileName)
// 	c.Data(http.StatusOK, "application/octet-stream", fileBuffer.Bytes())
// }

// GetLowStockReport
// @Summary Получить отчет по товарам с низким запасом
// @Tags reports
// @Description Отчет по товарам, у которых уровень запаса ниже порога
// @ID get-low-stock-report
// @Produce json
// @Param threshold query float64 true "Минимальный порог запаса"
// @Success 200 {array} models.LowStockReport "Список товаров с низким запасом"
// @Failure 400 {object} ErrorResponse "Ошибка ввода"
// @Failure 500 {object} ErrorResponse "Ошибка сервера"
// @Router /reports/low-stock [get]
// @Security ApiKeyAuth
func GetLowStockReport(c *gin.Context) {
	thresholdStr := c.Query("threshold")
	if thresholdStr == "" {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	threshold, err := strconv.ParseFloat(thresholdStr, 64)
	if err != nil {
		logger.Error.Printf("[controllers.GetLowStockReport] error parsing threshold: %v\n", err)
		handleError(c, errs.ErrValidationFailed)
		return
	}

	lowStockReport, err := service.GetLowStockReport(threshold)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, lowStockReport)
}

// GetSellerReport
// @Summary Получить отчет по продавцам
// @Tags reports
// @Description Отчет по количеству заказов и выручке каждого продавца
// @ID get-seller-report
// @Produce json
// @Success 200 {array} models.SellerReport "Список продавцов с количеством заказов и общей выручкой"
// @Failure 500 {object} ErrorResponse "Ошибка сервера"
// @Router /reports/sellers [get]
func GetSellerReport(c *gin.Context) {
	report, err := service.GetSellerReport()
	if err != nil {
		logger.Error.Printf("[controllers.GetSellerReport] error generating seller report: %v\n", err)
		handleError(c, errs.ErrServerError)
		return
	}

	c.JSON(http.StatusOK, report)
}

// GetSupplierReport
// @Summary Получить отчет по поставщикам
// @Tags reports
// @Description Отчет по поставщикам: количество товаров и общая стоимость поставок
// @ID get-supplier-report
// @Produce json
// @Success 200 {array} models.SupplierReport "Список поставщиков с количеством товаров и общей стоимостью"
// @Failure 500 {object} ErrorResponse "Ошибка сервера"
// @Router /reports/suppliers [get]
func GetSupplierReport(c *gin.Context) {
	report, err := service.GetSupplierReport()
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, report)
}

// /

// GetCategorySalesReport
// @Summary Получить отчет по категориям товаров
// @Tags reports
// @Description Отчет по категориям товаров с выручкой за указанный период
// @ID get-category-sales-report
// @Produce json
// @Param start_date query string true "Дата начала в формате YYYY-MM-DD"
// @Param end_date query string true "Дата окончания в формате YYYY-MM-DD"
// @Success 200 {array} models.CategorySalesReport "Список категорий с выручкой"
// @Failure 400 {object} ErrorResponse "Ошибка ввода"
// @Failure 500 {object} ErrorResponse "Ошибка сервера"
// @Router /reports/category-sales [get]
func GetCategorySalesReport(c *gin.Context) {
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	if startDateStr == "" || endDateStr == "" {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	// Парсим даты
	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		logger.Error.Printf("[controllers.GetCategorySalesReport] error parsing start_date: %v\n", err)
		handleError(c, errs.ErrValidationFailed)
		return
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		logger.Error.Printf("[controllers.GetCategorySalesReport] error parsing end_date: %v\n", err)
		handleError(c, errs.ErrValidationFailed)
		return
	}

	// Получаем отчет из сервиса
	report, err := service.GetCategorySalesReport(startDate, endDate)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, report)
}
