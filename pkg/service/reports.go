// C:\GoProject\src\eShop\pkg\service\reports.go

package service

import (
	"eShop/models"
	"eShop/pkg/repository"
	"time"
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

// /

func GetCategorySalesReport(startDate, endDate time.Time) ([]models.CategorySalesReport, error) {
	return repository.GetCategorySalesReport(startDate, endDate)
}
