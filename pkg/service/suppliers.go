// C:\GoProject\src\eShop\pkg\service\suppliers.go

package service

import (
	"eShop/errs"
	"eShop/logger"
	"eShop/models"
	"eShop/pkg/repository"
	"errors"
)

// CreateSupplier создаёт нового поставщика...
func CreateSupplier(supplier models.Supplier) error {
	return repository.CreateSupplier(supplier)
}

// UpdateSupplierByID обновляет данные поставщика по ID...
func UpdateSupplierByID(id uint, updatedSupplier models.Supplier) (supplier models.Supplier, err error) {
	supplier, err = repository.GetSupplierByID(id)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			return supplier, errs.ErrSupplierNotFound
		}
		return supplier, err
	}

	// Обновляем только переданные поля...
	if updatedSupplier.Name != "" {
		supplier.Name = updatedSupplier.Name
	}
	if updatedSupplier.Email != "" {
		supplier.Email = updatedSupplier.Email
	}
	if updatedSupplier.Phone != "" {
		supplier.Phone = updatedSupplier.Phone
	}

	err = repository.UpdateSupplierByID(supplier)
	if err != nil {
		logger.Error.Printf("[service.UpdateSupplierByID] error updating supplier: %v\n", err)
		return supplier, err
	}
	return supplier, nil
}

// SoftDeleteSupplierByID помечает поставщика как удалённого...
func SoftDeleteSupplierByID(id uint) error {
	return repository.SoftDeleteSupplierByID(id)
}

// RestoreSupplierByID восстанавливает удалённого поставщика...
func RestoreSupplierByID(id uint) error {
	return repository.RestoreSupplierByID(id)
}

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
