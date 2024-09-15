// C:\GoProject\src\eShop\pkg\service\suppliers.go

package service

import (
	"eShop/errs"
	"eShop/logger"
	"eShop/models"
	"eShop/pkg/repository"
	"errors"
	"time"
)

// CreateSupplier создает нового поставщика
func CreateSupplier(supplier models.Supplier) error {
	// Проверяем, что хотя бы одно из полей (Title или Email) не пустое
	if supplier.Title == "" && supplier.Email == "" {
		return errs.ErrValidationFailed // Возвращаем ошибку валидации, если оба поля пусты
	}

	// Проверяем, существует ли уже поставщик с таким именем или email
	existingSupplier, err := repository.GetSupplierByTitleOrEmail(supplier.Title, supplier.Email)
	if err != nil && err != errs.ErrRecordNotFound {
		return err
	}

	if existingSupplier.ID > 0 {
		return errs.ErrSupplierAlreadyExists // Поставщик уже существует
	}

	// Создаем нового поставщика через репозиторий
	if err := repository.CreateSupplier(supplier); err != nil {
		return err
	}

	return nil
}

// UpdateSupplierByID обновляет данные поставщика по его ID
func UpdateSupplierByID(id uint, updatedSupplier models.Supplier) (supplier models.Supplier, err error) {
	supplier, err = repository.GetSupplierByID(id)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			return supplier, errs.ErrSupplierNotFound
		}
		return supplier, err
	}

	// Обновляем только изменённые поля
	if updatedSupplier.Title != "" {
		supplier.Title = updatedSupplier.Title
	}
	if updatedSupplier.Email != "" {
		supplier.Email = updatedSupplier.Email
	}
	if updatedSupplier.Phone != "" {
		supplier.Phone = updatedSupplier.Phone
	}

	// Используем функцию обновления в репозитории
	err = repository.UpdateSupplierByID(supplier)
	if err != nil {
		return supplier, err
	}

	return supplier, nil
}

// SoftDeleteSupplierByID помечает поставщика как удалённого
func SoftDeleteSupplierByID(id uint) error {
	supplier, err := repository.GetSupplierByID(id)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			return errs.ErrSupplierNotFound
		}
		return err
	}

	// Проверяем, не был ли поставщик уже удалён
	if supplier.IsDeleted {
		return errs.ErrSupplierAlreadyDeleted
	}

	// Помечаем как удалённого
	supplier.IsDeleted = true
	currentTime := time.Now()
	supplier.DeletedAt = &currentTime

	// Используем общую функцию обновления для сохранения изменений
	if err := repository.UpdateSupplierByID(supplier); err != nil {
		return err
	}

	return nil
}

// RestoreSupplierByID восстанавливает мягко удалённого поставщика
func RestoreSupplierByID(id uint) error {
	supplier, err := repository.GetSupplierIncludingSoftDeleted(id)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			return errs.ErrSupplierNotFound
		}
		return err
	}

	// Проверяем, был ли поставщик действительно удалён
	if !supplier.IsDeleted {
		return errs.ErrSupplierNotDeleted
	}

	// Восстанавливаем поставщика
	supplier.IsDeleted = false
	supplier.DeletedAt = nil

	// Используем общую функцию обновления для сохранения изменений
	if err := repository.UpdateSupplierByID(supplier); err != nil {
		return err
	}

	return nil
}

// HardDeleteSupplierByID выполняет жёсткое удаление поставщика
func HardDeleteSupplierByID(id uint) error {
	// Используем существующую функцию, которая включает мягко удалённых
	supplier, err := repository.GetSupplierIncludingSoftDeleted(id)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			logger.Warning.Printf("[service.HardDeleteSupplierByID] supplier with ID: %v not found", id)
			return errs.ErrSupplierNotFound
		}
		return err
	}

	// Проверяем, был ли поставщик уже удалён
	if supplier.IsDeleted {
		logger.Warning.Printf("[service.HardDeleteSupplierByID] supplier with ID: %v is already deleted", id)
		return errs.ErrSupplierAlreadyDeleted
	}

	// Выполняем жёсткое удаление
	if err := repository.HardDeleteSupplierByID(supplier.ID); err != nil {
		return err
	}

	return nil
}

// GetAllSuppliers получает всех активных поставщиков
func GetAllSuppliers() (suppliers []models.Supplier, err error) {
	suppliers, err = repository.GetAllActiveSuppliers()
	if err != nil {
		return nil, err
	}

	// Если поставщиков нет, возвращаем пустой массив
	if len(suppliers) == 0 {
		logger.Warning.Printf("[service.GetAllSuppliers] no suppliers found")
	}

	return suppliers, nil
}

// GetAllDeletedSuppliers получает всех мягко удалённых поставщиков
func GetAllDeletedSuppliers() (suppliers []models.Supplier, err error) {
	suppliers, err = repository.GetAllDeletedSuppliers()
	if err != nil {
		return nil, err
	}

	// Если удалённых поставщиков нет, возвращаем пустой массив
	if len(suppliers) == 0 {
		logger.Warning.Printf("[service.GetAllDeletedSuppliers] no deleted suppliers found")
	}

	return suppliers, nil
}

// GetSupplierByID получает поставщика по его ID
func GetSupplierByID(id uint) (supplier models.Supplier, err error) {
	// Получаем поставщика через репозиторий
	supplier, err = repository.GetSupplierByID(id)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			logger.Warning.Printf("[service.GetSupplierByID] supplier with ID %d not found", id)
			return supplier, errs.ErrSupplierNotFound
		}
		logger.Error.Printf("[service.GetSupplierByID] error getting supplier by ID: %v\n", err)
		return supplier, err
	}
	return supplier, nil
}
