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

func GetSellerReport() ([]models.SellerReport, error) {
	return repository.GetSellerReport()
}

func GetSupplierReport() ([]models.SupplierReport, error) {
	return repository.GetSupplierReport()
}

func GetCategorySalesReport(startDate, endDate time.Time) ([]models.CategorySalesReport, error) {
	return repository.GetCategorySalesReport(startDate, endDate)
}

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

// GetLowStockReport получает отчет о товарах с низким запасом.
// Возвращает список товаров, у которых запас меньше или равен указанному порогу.
func GetLowStockReport(threshold float64) ([]models.LowStockReport, error) {
	return repository.GetLowStockProducts(threshold)
}

// GenerateLowStockReportFile генерирует отчет в формате CSV или XLSX, а также возвращает ZIP-архив при необходимости.
func GenerateLowStockReportFile(threshold float64, format string) (*bytes.Buffer, string, error) {
	report, err := GetLowStockReport(threshold)
	if err != nil {
		return nil, "", err
	}

	var buf bytes.Buffer
	var fileName string

	switch format {
	case "csv", "csvzip":
		// Создаем временный файл для CSV
		tmpFile, err := os.CreateTemp("", "low_stock_report_*.csv")
		if err != nil {
			return nil, "", err
		}
		defer tmpFile.Close()
		defer os.Remove(tmpFile.Name()) // Удаляем файл после завершения

		writer := csv.NewWriter(tmpFile)

		// Заголовки CSV
		writer.Write([]string{"Product ID", "Title", "Stock"})

		// Данные по товарам
		for _, item := range report {
			row := []string{
				fmt.Sprintf("%d", item.ProductID),
				item.Title,
				fmt.Sprintf("%.2f", item.Stock),
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
			fileName = "low_stock_report.zip"
			zipWriter := zip.NewWriter(&buf)
			defer zipWriter.Close()

			// Добавляем CSV файл в архив
			csvFileInZip, err := zipWriter.Create("low_stock_report.csv")
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
			fileName = "low_stock_report.csv"
		}

	case "xlsx", "xlsxzip":
		excelFile := excelize.NewFile()
		sheetName := "LowStockReport"
		excelFile.NewSheet(sheetName)

		// Заголовки Excel
		excelFile.SetCellValue(sheetName, "A1", "Product ID")
		excelFile.SetCellValue(sheetName, "B1", "Title")
		excelFile.SetCellValue(sheetName, "C1", "Stock")

		// Данные по товарам
		for i, item := range report {
			rowIndex := i + 2
			excelFile.SetCellValue(sheetName, fmt.Sprintf("A%d", rowIndex), item.ProductID)
			excelFile.SetCellValue(sheetName, fmt.Sprintf("B%d", rowIndex), item.Title)
			excelFile.SetCellValue(sheetName, fmt.Sprintf("C%d", rowIndex), item.Stock)
		}

		if format == "xlsxzip" {
			fileName = "low_stock_report.zip"
			zipWriter := zip.NewWriter(&buf)
			defer zipWriter.Close()

			// Создаем XLSX файл в буфере
			var xlsxBuffer bytes.Buffer
			if err := excelFile.Write(&xlsxBuffer); err != nil {
				return nil, "", err
			}

			// Добавляем XLSX файл в архив
			xlsxFileInZip, err := zipWriter.Create("low_stock_report.xlsx")
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
			fileName = "low_stock_report.xlsx"
		}

	default:
		return nil, "", fmt.Errorf("unsupported format: %s", format)
	}

	return &buf, fileName, nil
}

// GenerateSellerReportFile генерирует отчет в формате CSV или XLSX, а также возвращает ZIP-архив при необходимости.
func GenerateSellerReportFile(format string) (*bytes.Buffer, string, error) {
	report, err := GetSellerReport()
	if err != nil {
		return nil, "", err
	}

	var buf bytes.Buffer
	var fileName string

	switch format {
	case "csv", "csvzip":
		// Создаем временный файл для CSV
		tmpFile, err := os.CreateTemp("", "seller_report_*.csv")
		if err != nil {
			return nil, "", err
		}
		defer tmpFile.Close()
		defer os.Remove(tmpFile.Name()) // Удаляем файл после завершения

		writer := csv.NewWriter(tmpFile)

		// Заголовки CSV
		writer.Write([]string{"Seller ID", "Seller Name", "Order Count", "Total Revenue"})

		// Данные по продавцам
		for _, item := range report {
			row := []string{
				fmt.Sprintf("%d", item.SellerID),
				item.SellerName,
				fmt.Sprintf("%d", item.OrderCount),
				fmt.Sprintf("%.2f", item.TotalRevenue),
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
			fileName = "seller_report.zip"
			zipWriter := zip.NewWriter(&buf)
			defer zipWriter.Close()

			// Добавляем CSV файл в архив
			csvFileInZip, err := zipWriter.Create("seller_report.csv")
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
			fileName = "seller_report.csv"
		}

	case "xlsx", "xlsxzip":
		excelFile := excelize.NewFile()
		sheetName := "SellerReport"
		excelFile.NewSheet(sheetName)

		// Заголовки Excel
		excelFile.SetCellValue(sheetName, "A1", "Seller ID")
		excelFile.SetCellValue(sheetName, "B1", "Seller Name")
		excelFile.SetCellValue(sheetName, "C1", "Order Count")
		excelFile.SetCellValue(sheetName, "D1", "Total Revenue")

		// Данные по продавцам
		for i, item := range report {
			rowIndex := i + 2
			excelFile.SetCellValue(sheetName, fmt.Sprintf("A%d", rowIndex), item.SellerID)
			excelFile.SetCellValue(sheetName, fmt.Sprintf("B%d", rowIndex), item.SellerName)
			excelFile.SetCellValue(sheetName, fmt.Sprintf("C%d", rowIndex), item.OrderCount)
			excelFile.SetCellValue(sheetName, fmt.Sprintf("D%d", rowIndex), item.TotalRevenue)
		}

		if format == "xlsxzip" {
			fileName = "seller_report.zip"
			zipWriter := zip.NewWriter(&buf)
			defer zipWriter.Close()

			// Создаем XLSX файл в буфере
			var xlsxBuffer bytes.Buffer
			if err := excelFile.Write(&xlsxBuffer); err != nil {
				return nil, "", err
			}

			// Добавляем XLSX файл в архив
			xlsxFileInZip, err := zipWriter.Create("seller_report.xlsx")
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
			fileName = "seller_report.xlsx"
		}

	default:
		return nil, "", fmt.Errorf("unsupported format: %s", format)
	}

	return &buf, fileName, nil
}
