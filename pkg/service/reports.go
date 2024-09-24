// C:\GoProject\src\eShop\pkg\service\reports.go

package service

import (
	"archive/zip"
	"bytes"
	"eShop/logger"
	"eShop/models"
	"eShop/pkg/repository"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/xuri/excelize/v2"
)

// GenerateReport генерирует отчет в зависимости от типа.
func GenerateReport(reportType string, startDate, endDate time.Time, threshold float64) (interface{}, error) {
	switch reportType {
	case "sales":
		return repository.GetSalesReport(startDate, endDate)
	case "low-stock":
		return repository.GetLowStockProducts(threshold)
	case "seller":
		return repository.GetSellerReport()
	case "supplier":
		return repository.GetSupplierReport()
	case "category-sales":
		return repository.GetCategorySalesReport(startDate, endDate)
	default:
		return nil, fmt.Errorf("unsupported report type: %s", reportType)
	}
}

// GenerateReportFile генерирует файл отчета (CSV, XLSX, ZIP) для выбранного типа отчета.
func GenerateReportFile(reportType string, startDate, endDate time.Time, threshold float64, format string) (*bytes.Buffer, string, error) {
	// Получаем данные для отчета
	report, err := GenerateReport(reportType, startDate, endDate, threshold)
	if err != nil {
		return nil, "", err
	}

	// Создаем файл отчета
	switch reportType {
	case "sales":
		return generateSalesReportFile(report.(models.SalesReport), format)
	case "low-stock":
		return generateLowStockReportFile(report.([]models.LowStockReport), format)
	case "seller":
		return generateSellerReportFile(report.([]models.SellerReport), format)
	case "supplier":
		return generateSupplierReportFile(report.([]models.SupplierReport), format)
	case "category-sales":
		return generateCategorySalesReportFile(report.([]models.CategorySalesReport), format)
	default:
		return nil, "", fmt.Errorf("unsupported report type: %s", reportType)
	}
}

func generateCSVOrXLSXFile(report interface{}, headers []string, rows [][]string, format string) (*bytes.Buffer, string, error) {
	var buf bytes.Buffer
	var fileName string

	switch format {
	case "csv", "csvzip":
		// Создаем временный файл для CSV
		tmpFile, err := os.CreateTemp("", "report_*.csv")
		if err != nil {
			return nil, "", err
		}
		defer tmpFile.Close()
		defer os.Remove(tmpFile.Name()) // Удаляем файл после завершения

		writer := csv.NewWriter(tmpFile)
		// Записываем заголовки
		writer.Write(headers)

		// Записываем строки данных
		for _, row := range rows {
			writer.Write(row)
		}

		writer.Flush()
		if err := writer.Error(); err != nil {
			return nil, "", err
		}

		if format == "csvzip" {
			fileName = "report.zip"
			zipWriter := zip.NewWriter(&buf)
			defer zipWriter.Close()

			csvFileInZip, err := zipWriter.Create("report.csv")
			if err != nil {
				return nil, "", err
			}

			tmpFile.Seek(0, io.SeekStart)
			_, err = io.Copy(csvFileInZip, tmpFile)
			if err != nil {
				return nil, "", err
			}

			zipWriter.Close()
		} else {
			tmpFile.Seek(0, io.SeekStart)
			_, err = io.Copy(&buf, tmpFile)
			if err != nil {
				return nil, "", err
			}
			fileName = "report.csv"
		}

	case "xlsx", "xlsxzip":
		excelFile := excelize.NewFile()
		sheetName := "Report"
		excelFile.NewSheet(sheetName)

		// // Удаляем лист "Sheet1", созданный по умолчанию
		// excelFile.DeleteSheet("Sheet1")
		// Удаляем лист "Sheet1", созданный по умолчанию
		if err := excelFile.DeleteSheet("Sheet1"); err != nil {
			logger.Error.Printf("Error deleting default sheet: %v", err)
			// return nil, "", fmt.Errorf("failed to delete default sheet: %v", err) // Не возвращаем ошибку, продолжаем выполнение...
		}

		// Заголовки
		for i, header := range headers {
			col := string('A' + i)
			excelFile.SetCellValue(sheetName, fmt.Sprintf("%s1", col), header)
		}

		// Данные
		for i, row := range rows {
			for j, cell := range row {
				col := string('A' + j)
				excelFile.SetCellValue(sheetName, fmt.Sprintf("%s%d", col, i+2), cell)
			}
		}

		if format == "xlsxzip" {
			fileName = "report.zip"
			zipWriter := zip.NewWriter(&buf)
			defer zipWriter.Close()

			var xlsxBuffer bytes.Buffer
			if err := excelFile.Write(&xlsxBuffer); err != nil {
				return nil, "", err
			}

			xlsxFileInZip, err := zipWriter.Create("report.xlsx")
			if err != nil {
				return nil, "", err
			}
			_, err = io.Copy(xlsxFileInZip, &xlsxBuffer)
			if err != nil {
				return nil, "", err
			}

			zipWriter.Close()
		} else {
			if err := excelFile.Write(&buf); err != nil {
				return nil, "", err
			}
			fileName = "report.xlsx"
		}

	default:
		return nil, "", fmt.Errorf("unsupported format: %s", format)
	}

	return &buf, fileName, nil
}

func generateSalesReportFile(report models.SalesReport, format string) (*bytes.Buffer, string, error) {
	headers := []string{"Product ID", "Title", "Quantity", "Total"}
	rows := [][]string{}

	for _, item := range report.TopSelling {
		rows = append(rows, []string{
			fmt.Sprintf("%d", item.ProductID),
			item.Title,
			fmt.Sprintf("%.2f", item.Quantity),
			fmt.Sprintf("%.2f", item.Total),
		})
	}

	return generateCSVOrXLSXFile(report, headers, rows, format)
}

func generateLowStockReportFile(report []models.LowStockReport, format string) (*bytes.Buffer, string, error) {
	headers := []string{"Product ID", "Title", "Stock"}
	rows := [][]string{}

	for _, item := range report {
		rows = append(rows, []string{
			fmt.Sprintf("%d", item.ProductID),
			item.Title,
			fmt.Sprintf("%.2f", item.Stock),
		})
	}

	return generateCSVOrXLSXFile(report, headers, rows, format)
}

func generateSellerReportFile(report []models.SellerReport, format string) (*bytes.Buffer, string, error) {
	headers := []string{"Seller ID", "Seller Name", "Order Count", "Total Revenue"}
	rows := [][]string{}

	for _, item := range report {
		rows = append(rows, []string{
			fmt.Sprintf("%d", item.SellerID),
			item.SellerName,
			fmt.Sprintf("%d", item.OrderCount),
			fmt.Sprintf("%.2f", item.TotalRevenue),
		})
	}

	return generateCSVOrXLSXFile(report, headers, rows, format)
}

func generateSupplierReportFile(report []models.SupplierReport, format string) (*bytes.Buffer, string, error) {
	headers := []string{"Supplier ID", "Supplier Name", "Product Count", "Total Supplies"}
	rows := [][]string{}

	for _, item := range report {
		rows = append(rows, []string{
			fmt.Sprintf("%d", item.SupplierID),
			item.SupplierName,
			fmt.Sprintf("%d", item.ProductCount),
			fmt.Sprintf("%.2f", item.TotalSupplies),
		})
	}

	return generateCSVOrXLSXFile(report, headers, rows, format)
}

func generateCategorySalesReportFile(report []models.CategorySalesReport, format string) (*bytes.Buffer, string, error) {
	headers := []string{"Category ID", "Category Name", "Total Sales"}
	rows := [][]string{}

	for _, item := range report {
		rows = append(rows, []string{
			fmt.Sprintf("%d", item.CategoryID),
			item.CategoryName,
			fmt.Sprintf("%.2f", item.TotalSales),
		})
	}

	return generateCSVOrXLSXFile(report, headers, rows, format)
}
