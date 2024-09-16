// C:\GoProject\src\eShop\pkg\service\taxes.go

package service

import (
	"eShop/errs"
	"eShop/logger"
	"eShop/models"
	"eShop/pkg/repository"
	"errors"
)

// GetAllTaxes получает текущие ставки налогов
func GetAllTaxes() (taxes []models.Taxes, err error) {
	taxes, err = repository.GetAllTaxes()
	if err != nil {
		logger.Error.Printf("Error getting taxes: %v", err)
		return nil, errs.ErrSomethingWentWrong
	}

	if len(taxes) == 0 {
		logger.Warning.Printf("No taxes found")
	}

	return taxes, nil
}

// UpdateTaxByID обновляет ставку налога по его ID
func UpdateTaxByID(id uint, updatedTax models.Taxes) (tax models.Taxes, err error) {
	// Получаем существующую налоговую ставку
	tax, err = repository.GetTaxByID(id)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			return tax, errs.ErrRecordNotFound
		}
		return tax, err
	}

	// Обновляем только процентную ставку, если она указана
	if updatedTax.Rate > 0 {
		tax.Rate = updatedTax.Rate
	}

	// Используем функцию обновления в репозитории
	err = repository.UpdateTax(tax)
	if err != nil {
		return tax, err
	}

	return tax, nil
}
