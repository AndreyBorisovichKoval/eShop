// C:\GoProject\src\eShop\pkg\service\reports.go

package service

import (
	"bytes"
	"eShop/models"
	"eShop/pkg/repository"
	"encoding/csv"
	"fmt"
	"time"

	"github.com/xuri/excelize/v2"
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

// // GenerateSalesReportCSV создает CSV-файл из отчёта о продажах
// func GenerateSalesReportCSV(report models.SalesReport) ([]byte, error) {
// 	var buf bytes.Buffer
// 	writer := csv.NewWriter(&buf)

// 	// Заголовки
// 	err := writer.Write([]string{"Product ID", "Product Name", "Quantity Sold", "Total Sales"})
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Данные
// 	for _, item := range report.TopSelling {
// 		err := writer.Write([]string{
// 			strconv.Itoa(int(item.ProductID)),
// 			item.Title,
// 			strconv.FormatFloat(item.Quantity, 'f', 2, 64),
// 			strconv.FormatFloat(item.Total, 'f', 2, 64),
// 		})
// 		if err != nil {
// 			return nil, err
// 		}
// 	}

// 	writer.Flush()
// 	return buf.Bytes(), nil
// }

// // GenerateSalesReportXLSX создает XLSX-файл из отчёта о продажах
// func GenerateSalesReportXLSX(report models.SalesReport) ([]byte, error) {
// 	file := xlsx.NewFile()
// 	sheet, err := file.AddSheet("Sales Report")
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Заголовки
// 	row := sheet.AddRow()
// 	row.AddCell().Value = "Product ID"
// 	row.AddCell().Value = "Product Name"
// 	row.AddCell().Value = "Quantity Sold"
// 	row.AddCell().Value = "Total Sales"

// 	// Данные
// 	for _, item := range report.TopSelling {
// 		row := sheet.AddRow()
// 		row.AddCell().Value = strconv.Itoa(int(item.ProductID))
// 		row.AddCell().Value = item.Title
// 		row.AddCell().Value = strconv.FormatFloat(item.Quantity, 'f', 2, 64)
// 		row.AddCell().Value = strconv.FormatFloat(item.Total, 'f', 2, 64)
// 	}

// 	var buf bytes.Buffer
// 	err = file.Write(&buf)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return buf.Bytes(), nil
// }

// // GenerateSalesReportFile генерирует отчет в формате CSV или XLSX.
// func GenerateSalesReportFile(startDate, endDate time.Time, format string) (*bytes.Buffer, string, error) {
// 	report, err := GetSalesReport(startDate, endDate)
// 	if err != nil {
// 		return nil, "", err
// 	}

// 	var buf bytes.Buffer
// 	var fileName string

// 	switch format {
// 	case "csv":
// 		writer := csv.NewWriter(&buf)
// 		defer writer.Flush()

// 		// Заголовки CSV
// 		writer.Write([]string{"Product ID", "Title", "Quantity", "Total"})

// 		// Данные по продажам
// 		for _, item := range report.TopSelling {
// 			row := []string{
// 				fmt.Sprintf("%d", item.ProductID),
// 				item.Title,
// 				fmt.Sprintf("%.2f", item.Quantity),
// 				fmt.Sprintf("%.2f", item.Total),
// 			}
// 			writer.Write(row)
// 		}
// 		fileName = "sales_report.csv"

// 	case "xlsx":
// 		excelFile := excelize.NewFile()
// 		sheetName := "SalesReport"
// 		excelFile.NewSheet(sheetName)

// 		// Заголовки Excel
// 		excelFile.SetCellValue(sheetName, "A1", "Product ID")
// 		excelFile.SetCellValue(sheetName, "B1", "Title")
// 		excelFile.SetCellValue(sheetName, "C1", "Quantity")
// 		excelFile.SetCellValue(sheetName, "D1", "Total")

// 		// Данные по продажам
// 		for i, item := range report.TopSelling {
// 			rowIndex := i + 2
// 			excelFile.SetCellValue(sheetName, fmt.Sprintf("A%d", rowIndex), item.ProductID)
// 			excelFile.SetCellValue(sheetName, fmt.Sprintf("B%d", rowIndex), item.Title)
// 			excelFile.SetCellValue(sheetName, fmt.Sprintf("C%d", rowIndex), item.Quantity)
// 			excelFile.SetCellValue(sheetName, fmt.Sprintf("D%d", rowIndex), item.Total)
// 		}

// 		// Сохраняем в буфер
// 		if err := excelFile.Write(&buf); err != nil {
// 			return nil, "", err
// 		}
// 		fileName = "sales_report.xlsx"

// 	default:
// 		return nil, "", fmt.Errorf("unsupported format: %s", format)
// 	}

// 	return &buf, fileName, nil
// }

// GenerateSalesReportFile генерирует отчет в формате CSV или XLSX.
func GenerateSalesReportFile(startDate, endDate time.Time, format string) (*bytes.Buffer, string, error) {
	report, err := GetSalesReport(startDate, endDate)
	if err != nil {
		return nil, "", err
	}

	var buf bytes.Buffer
	var fileName string

	switch format {
	case "csv":
		writer := csv.NewWriter(&buf)
		defer writer.Flush()

		// Заголовки CSV
		writer.Write([]string{"Product ID", "Title", "Quantity", "Total"})

		// Данные по продажам
		for _, item := range report.TopSelling {
			row := []string{
				fmt.Sprintf("%d", item.ProductID),
				item.Title,
				fmt.Sprintf("%.2f", item.Quantity),
				fmt.Sprintf("%.2f", item.Total),
			}
			writer.Write(row)
		}
		fileName = "sales_report.csv"

	case "xlsx":
		excelFile := excelize.NewFile()
		sheetName := "SalesReport"
		excelFile.NewSheet(sheetName)

		// Заголовки Excel
		excelFile.SetCellValue(sheetName, "A1", "Product ID")
		excelFile.SetCellValue(sheetName, "B1", "Title")
		excelFile.SetCellValue(sheetName, "C1", "Quantity")
		excelFile.SetCellValue(sheetName, "D1", "Total")

		// Данные по продажам
		for i, item := range report.TopSelling {
			rowIndex := i + 2
			excelFile.SetCellValue(sheetName, fmt.Sprintf("A%d", rowIndex), item.ProductID)
			excelFile.SetCellValue(sheetName, fmt.Sprintf("B%d", rowIndex), item.Title)
			excelFile.SetCellValue(sheetName, fmt.Sprintf("C%d", rowIndex), item.Quantity)
			excelFile.SetCellValue(sheetName, fmt.Sprintf("D%d", rowIndex), item.Total)
		}

		// Сохраняем в буфер
		if err := excelFile.Write(&buf); err != nil {
			return nil, "", err
		}
		fileName = "sales_report.xlsx"

	default:
		return nil, "", fmt.Errorf("unsupported format: %s", format)
	}

	return &buf, fileName, nil
}
