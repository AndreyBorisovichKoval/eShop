// C:\GoProject\src\eShop\pkg\service\reports.go

package service

import (
	"archive/zip"
	"bytes"
	"eShop/models"
	"eShop/pkg/repository"
	"encoding/csv"
	"fmt"
	"io"
	"os"
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

// // GenerateSalesReportFile генерирует отчет в формате CSV, XLSX или ZIP.
// func GenerateSalesReportFile(startDate, endDate time.Time, format string) (*bytes.Buffer, string, error) {
// 	report, err := GetSalesReport(startDate, endDate)
// 	if err != nil {
// 		return nil, "", err
// 	}

// 	var buf bytes.Buffer
// 	var fileName string

// 	switch format {
// 	case "csv":
// 		// Генерация CSV файла
// 		writer := csv.NewWriter(&buf)
// 		defer writer.Flush()

// 		writer.Write([]string{"Product ID", "Title", "Quantity", "Total"})
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
// 		// Генерация XLSX файла
// 		excelFile := excelize.NewFile()
// 		sheetName := "SalesReport"
// 		excelFile.NewSheet(sheetName)

// 		excelFile.SetCellValue(sheetName, "A1", "Product ID")
// 		excelFile.SetCellValue(sheetName, "B1", "Title")
// 		excelFile.SetCellValue(sheetName, "C1", "Quantity")
// 		excelFile.SetCellValue(sheetName, "D1", "Total")

// 		for i, item := range report.TopSelling {
// 			rowIndex := i + 2
// 			excelFile.SetCellValue(sheetName, fmt.Sprintf("A%d", rowIndex), item.ProductID)
// 			excelFile.SetCellValue(sheetName, fmt.Sprintf("B%d", rowIndex), item.Title)
// 			excelFile.SetCellValue(sheetName, fmt.Sprintf("C%d", rowIndex), item.Quantity)
// 			excelFile.SetCellValue(sheetName, fmt.Sprintf("D%d", rowIndex), item.Total)
// 		}

// 		if err := excelFile.Write(&buf); err != nil {
// 			return nil, "", err
// 		}
// 		fileName = "sales_report.xlsx"

// 	case "csv_zip", "xlsx_zip":
// 		// Генерация ZIP файла
// 		zipBuf := new(bytes.Buffer)
// 		zipWriter := zip.NewWriter(zipBuf)

// 		var innerFileName string
// 		var innerBuf bytes.Buffer

// 		if format == "csv_zip" {
// 			innerFileName = "sales_report.csv"
// 			writer := csv.NewWriter(&innerBuf)
// 			defer writer.Flush()

// 			writer.Write([]string{"Product ID", "Title", "Quantity", "Total"})
// 			for _, item := range report.TopSelling {
// 				row := []string{
// 					fmt.Sprintf("%d", item.ProductID),
// 					item.Title,
// 					fmt.Sprintf("%.2f", item.Quantity),
// 					fmt.Sprintf("%.2f", item.Total),
// 				}
// 				writer.Write(row)
// 			}
// 		} else {
// 			innerFileName = "sales_report.xlsx"
// 			excelFile := excelize.NewFile()
// 			sheetName := "SalesReport"
// 			excelFile.NewSheet(sheetName)

// 			excelFile.SetCellValue(sheetName, "A1", "Product ID")
// 			excelFile.SetCellValue(sheetName, "B1", "Title")
// 			excelFile.SetCellValue(sheetName, "C1", "Quantity")
// 			excelFile.SetCellValue(sheetName, "D1", "Total")

// 			for i, item := range report.TopSelling {
// 				rowIndex := i + 2
// 				excelFile.SetCellValue(sheetName, fmt.Sprintf("A%d", rowIndex), item.ProductID)
// 				excelFile.SetCellValue(sheetName, fmt.Sprintf("B%d", rowIndex), item.Title)
// 				excelFile.SetCellValue(sheetName, fmt.Sprintf("C%d", rowIndex), item.Quantity)
// 				excelFile.SetCellValue(sheetName, fmt.Sprintf("D%d", rowIndex), item.Total)
// 			}

// 			if err := excelFile.Write(&innerBuf); err != nil {
// 				return nil, "", err
// 			}
// 		}

// 		// Добавляем файл в архив
// 		zipFile, err := zipWriter.Create(innerFileName)
// 		if err != nil {
// 			return nil, "", err
// 		}

// 		_, err = zipFile.Write(innerBuf.Bytes())
// 		if err != nil {
// 			return nil, "", err
// 		}

// 		// Закрываем архив
// 		if err := zipWriter.Close(); err != nil {
// 			return nil, "", err
// 		}

// 		buf = *zipBuf
// 		fileName = "sales_report.zip"

// 	default:
// 		return nil, "", fmt.Errorf("unsupported format: %s", format)
// 	}

// 	return &buf, fileName, nil
// }

// // GenerateSalesReportFile генерирует отчет в формате CSV, XLSX или ZIP.
// func GenerateSalesReportFile(startDate, endDate time.Time, format string) (*bytes.Buffer, string, error) {
// 	report, err := GetSalesReport(startDate, endDate)
// 	if err != nil {
// 		return nil, "", err
// 	}

// 	var buf bytes.Buffer
// 	var fileName string

// 	switch format {
// 	case "csv":
// 		// Генерация CSV файла
// 		writer := csv.NewWriter(&buf)
// 		// defer writer.Flush()

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

// 		// Обязательно очищаем буфер перед архивированием
// 		writer.Flush()

// 		// Проверяем на ошибки записи
// 		if err := writer.Error(); err != nil {
// 			return nil, "", err
// 		}

// 		fileName = "sales_report.csv"

// 	case "xlsx":
// 		// Генерация XLSX файла
// 		excelFile := excelize.NewFile()
// 		sheetName := "SalesReport"
// 		excelFile.NewSheet(sheetName)

// 		excelFile.SetCellValue(sheetName, "A1", "Product ID")
// 		excelFile.SetCellValue(sheetName, "B1", "Title")
// 		excelFile.SetCellValue(sheetName, "C1", "Quantity")
// 		excelFile.SetCellValue(sheetName, "D1", "Total")

// 		for i, item := range report.TopSelling {
// 			rowIndex := i + 2
// 			excelFile.SetCellValue(sheetName, fmt.Sprintf("A%d", rowIndex), item.ProductID)
// 			excelFile.SetCellValue(sheetName, fmt.Sprintf("B%d", rowIndex), item.Title)
// 			excelFile.SetCellValue(sheetName, fmt.Sprintf("C%d", rowIndex), item.Quantity)
// 			excelFile.SetCellValue(sheetName, fmt.Sprintf("D%d", rowIndex), item.Total)
// 		}

// 		if err := excelFile.Write(&buf); err != nil {
// 			return nil, "", err
// 		}
// 		fileName = "sales_report.xlsx"

// 	case "csv_zip", "xlsx_zip":
// 		// Генерация ZIP файла
// 		zipBuf := new(bytes.Buffer)
// 		zipWriter := zip.NewWriter(zipBuf)

// 		var innerFileName string
// 		var innerBuf bytes.Buffer

// 		if format == "csv_zip" {
// 			innerFileName = "sales_report.csv"
// 			writer := csv.NewWriter(&innerBuf)
// 			defer writer.Flush()

// 			writer.Write([]string{"Product ID", "Title", "Quantity", "Total"})
// 			for _, item := range report.TopSelling {
// 				row := []string{
// 					fmt.Sprintf("%d", item.ProductID),
// 					item.Title,
// 					fmt.Sprintf("%.2f", item.Quantity),
// 					fmt.Sprintf("%.2f", item.Total),
// 				}
// 				writer.Write(row)
// 			}
// 		} else {
// 			innerFileName = "sales_report.xlsx"
// 			excelFile := excelize.NewFile()
// 			sheetName := "SalesReport"
// 			excelFile.NewSheet(sheetName)

// 			excelFile.SetCellValue(sheetName, "A1", "Product ID")
// 			excelFile.SetCellValue(sheetName, "B1", "Title")
// 			excelFile.SetCellValue(sheetName, "C1", "Quantity")
// 			excelFile.SetCellValue(sheetName, "D1", "Total")

// 			for i, item := range report.TopSelling {
// 				rowIndex := i + 2
// 				excelFile.SetCellValue(sheetName, fmt.Sprintf("A%d", rowIndex), item.ProductID)
// 				excelFile.SetCellValue(sheetName, fmt.Sprintf("B%d", rowIndex), item.Title)
// 				excelFile.SetCellValue(sheetName, fmt.Sprintf("C%d", rowIndex), item.Quantity)
// 				excelFile.SetCellValue(sheetName, fmt.Sprintf("D%d", rowIndex), item.Total)
// 			}

// 			if err := excelFile.Write(&innerBuf); err != nil {
// 				return nil, "", err
// 			}
// 		}

// 		// Добавляем файл в архив
// 		zipFile, err := zipWriter.Create(innerFileName)
// 		if err != nil {
// 			return nil, "", err
// 		}

// 		_, err = zipFile.Write(innerBuf.Bytes())
// 		if err != nil {
// 			return nil, "", err
// 		}

// 		// Закрываем архив
// 		if err := zipWriter.Close(); err != nil {
// 			return nil, "", err
// 		}

// 		buf = *zipBuf
// 		fileName = "sales_report.zip"

// 	default:
// 		return nil, "", fmt.Errorf("unsupported format: %s", format)
// 	}

// 	return &buf, fileName, nil
// }

// GenerateSalesReportFile генерирует отчет в формате CSV или XLSX, а также возвращает ZIP-архив при необходимости.
func GenerateSalesReportFile(startDate, endDate time.Time, format string) (*bytes.Buffer, string, error) {
	report, err := GetSalesReport(startDate, endDate)
	if err != nil {
		return nil, "", err
	}

	var buf bytes.Buffer
	var fileName string

	switch format {
	case "csv", "csvzip":
		// Создаем временный файл для CSV
		tmpFile, err := os.CreateTemp("", "sales_report_*.csv")
		if err != nil {
			return nil, "", err
		}
		defer tmpFile.Close()
		defer os.Remove(tmpFile.Name()) // Удаляем файл после завершения

		writer := csv.NewWriter(tmpFile)

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

		// Обязательно очищаем буфер
		writer.Flush()

		// Проверяем на ошибки записи
		if err := writer.Error(); err != nil {
			return nil, "", err
		}

		// Если запрашивается ZIP
		if format == "csvzip" {
			fileName = "sales_report.zip"
			zipWriter := zip.NewWriter(&buf)
			defer zipWriter.Close()

			// Добавляем CSV файл в архив
			csvFileInZip, err := zipWriter.Create("sales_report.csv")
			if err != nil {
				return nil, "", err
			}

			// Открываем временный файл для чтения
			tmpFile.Seek(0, io.SeekStart)
			_, err = io.Copy(csvFileInZip, tmpFile)
			if err != nil {
				return nil, "", err
			}

			zipWriter.Close()
		} else {
			// Читаем временный файл в буфер
			tmpFile.Seek(0, io.SeekStart)
			_, err = io.Copy(&buf, tmpFile)
			if err != nil {
				return nil, "", err
			}
			fileName = "sales_report.csv"
		}

	case "xlsx", "xlsxzip":
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

		if format == "xlsxzip" {
			fileName = "sales_report.zip"
			zipWriter := zip.NewWriter(&buf)
			defer zipWriter.Close()

			// Создаем XLSX файл в буфере
			var xlsxBuffer bytes.Buffer
			if err := excelFile.Write(&xlsxBuffer); err != nil {
				return nil, "", err
			}

			// Добавляем XLSX файл в архив
			xlsxFileInZip, err := zipWriter.Create("sales_report.xlsx")
			if err != nil {
				return nil, "", err
			}
			_, err = io.Copy(xlsxFileInZip, &xlsxBuffer)
			if err != nil {
				return nil, "", err
			}

			zipWriter.Close()
		} else {
			// Сохраняем в буфер для XLSX
			if err := excelFile.Write(&buf); err != nil {
				return nil, "", err
			}
			fileName = "sales_report.xlsx"
		}

	default:
		return nil, "", fmt.Errorf("unsupported format: %s", format)
	}

	return &buf, fileName, nil
}
