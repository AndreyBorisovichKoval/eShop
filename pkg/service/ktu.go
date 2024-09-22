package service

import (
	"eShop/errs"
	"eShop/logger"
	"eShop/pkg/repository"
	"fmt"
)

// CalculateKTU вычисляет КТУ для всех сотрудников на основе продаж за указанный месяц
func CalculateKTU(year, month int) (map[string]float64, error) {
	// Получаем данные по продажам всех пользователей за указанный месяц
	salesData, err := repository.GetSalesDataByMonth(year, month)
	if err != nil {
		logger.Error.Printf("[service.CalculateKTU] Error getting sales data: %v\n", err)
		return nil, errs.ErrServerError
	}

	// Рассчитываем общую сумму продаж
	var totalSales float64
	for _, sales := range salesData {
		totalSales += sales.TotalSales
	}

	if totalSales == 0 {
		return nil, fmt.Errorf("no sales found for the specified period")
	}

	// Рассчитываем КТУ для каждого сотрудника
	ktuResult := make(map[string]float64)
	for _, sales := range salesData {
		ktuResult[sales.FullName] = sales.TotalSales / totalSales
	}

	return ktuResult, nil
}
