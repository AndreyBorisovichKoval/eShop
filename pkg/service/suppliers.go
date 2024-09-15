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
	// Проверяем, что хотя бы одно из полей (Name или Email) не пустое
	if supplier.Name == "" && supplier.Email == "" {
		return errs.ErrValidationFailed // Возвращаем ошибку валидации, если оба поля пусты
	}

	// Проверяем, существует ли уже поставщик с таким именем или email
	existingSupplier, err := repository.GetSupplierByNameOrEmail(supplier.Name, supplier.Email)
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
	if updatedSupplier.Name != "" {
		supplier.Name = updatedSupplier.Name
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

// /

// HardDeleteSupplierByID удаляет поставщика...
func HardDeleteSupplierByID(id uint) error {
	return repository.HardDeleteSupplierByID(id)
}

// GetAllSuppliers возвращает список всех поставщиков...
func GetAllSuppliers() (suppliers []models.Supplier, err error) {
	suppliers, err = repository.GetAllSuppliers()
	if err != nil {
		logger.Error.Printf("[service.GetAllSuppliers] error: %v\n", err)
		return nil, err
	}
	if len(suppliers) == 0 {
		logger.Warning.Printf("[service.GetAllSuppliers] no suppliers found")
	}
	return suppliers, nil
}

// GetAllDeletedSuppliers получает список всех удалённых поставщиков...
func GetAllDeletedSuppliers() (suppliers []models.Supplier, err error) {
	// Получаем всех удалённых поставщиков из репозитория...
	suppliers, err = repository.GetAllDeletedSuppliers()
	if err != nil {
		return nil, err
	}

	// Возвращаем список удалённых поставщиков...
	return suppliers, nil
}

// GetSupplierByID возвращает данные поставщика по ID...
func GetSupplierByID(id uint) (supplier models.Supplier, err error) {
	supplier, err = repository.GetSupplierByID(id)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			return supplier, errs.ErrSupplierNotFound
		}
		return supplier, err
	}
	return supplier, nil
}
