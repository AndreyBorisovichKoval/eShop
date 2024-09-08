// C:\GoProject\src\eShop\pkg\repository\gorm.go

package repository

import (
	"eShop/errs"
	"eShop/logger"
	"errors"

	"gorm.io/gorm"
)

func translateError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		logger.Warning.Printf("Record not found error: %v", err)
		return errs.ErrRecordNotFound
	}

	// Добавить логирование для других ошибок по мере необходимости
	logger.Error.Printf("Unhandled error: %v", err)

	return err
}
