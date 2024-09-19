// C:\GoProject\src\eShop\pkg\service\reports.go

package service

import (
	"bytes"
	"eShop/models"
	"eShop/pkg/repository"
	"encoding/csv"
	"strconv"
	"time"

	"github.com/tealeg/xlsx"
)

// GetSalesReport получает отчет о продажах за указанный период.
// Возвращает информацию о сумме продаж, общем количестве проданных товаров и топ-продуктах.
func GetSalesReport(startDate, endDate time.Time) (models.SalesReport, error) {
	report, err := repository.GetSalesReport(startDate, endDate)
	if err != nil {
		return report, err
	}
	return report, nil
}

// GetLowStockReport получает отчет о товарах с низким запасом.
// Возвращает список товаров, у которых запас меньше или равен указанному порогу.
func GetLowStockReport(threshold float64) ([]models.LowStockReport, error) {
	return repository.GetLowStockProducts(threshold)
}

func GetSellerReport() ([]models.SellerReport, error) {
	return repository.GetSellerReport()
}

func GetSupplierReport() ([]models.SupplierReport, error) {
	return repository.GetSupplierReport()
}

func GetCategorySalesReport(startDate, endDate time.Time) ([]models.CategorySalesReport, error) {
	return repository.GetCategorySalesReport(startDate, endDate)
}

// /

// GenerateSalesReportCSV создает CSV-файл из отчёта о продажах
func GenerateSalesReportCSV(report models.SalesReport) ([]byte, error) {
	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)

	// Заголовки
	err := writer.Write([]string{"Product ID", "Product Name", "Quantity Sold", "Total Sales"})
	if err != nil {
		return nil, err
	}

	// Данные
	for _, item := range report.TopSelling {
		err := writer.Write([]string{
			strconv.Itoa(int(item.ProductID)),
			item.Title,
			strconv.FormatFloat(item.Quantity, 'f', 2, 64),
			strconv.FormatFloat(item.Total, 'f', 2, 64),
		})
		if err != nil {
			return nil, err
		}
	}

	writer.Flush()
	return buf.Bytes(), nil
}

// GenerateSalesReportXLSX создает XLSX-файл из отчёта о продажах
func GenerateSalesReportXLSX(report models.SalesReport) ([]byte, error) {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sales Report")
	if err != nil {
		return nil, err
	}

	// Заголовки
	row := sheet.AddRow()
	row.AddCell().Value = "Product ID"
	row.AddCell().Value = "Product Name"
	row.AddCell().Value = "Quantity Sold"
	row.AddCell().Value = "Total Sales"

	// Данные
	for _, item := range report.TopSelling {
		row := sheet.AddRow()
		row.AddCell().Value = strconv.Itoa(int(item.ProductID))
		row.AddCell().Value = item.Title
		row.AddCell().Value = strconv.FormatFloat(item.Quantity, 'f', 2, 64)
		row.AddCell().Value = strconv.FormatFloat(item.Total, 'f', 2, 64)
	}

	var buf bytes.Buffer
	err = file.Write(&buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
