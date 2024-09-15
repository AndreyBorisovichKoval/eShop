// C:\GoProject\src\eShop\pkg\service\suppliers.go

package service

import (
	"eShop/errs"
	"eShop/models"
	"eShop/pkg/repository"
	"errors"
)

func CreateSupplier(supplier models.Supplier) error {
	return repository.CreateSupplier(supplier)
}

func GetAllSuppliers() ([]models.Supplier, error) {
	return repository.GetAllSuppliers()
}

func GetSupplierByID(id uint) (models.Supplier, error) {
	supplier, err := repository.GetSupplierByID(id)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			return supplier, errs.ErrSupplierNotFound
		}
		return supplier, err
	}
	return supplier, nil
}

func UpdateSupplierByID(id uint, updatedSupplier models.Supplier) (models.Supplier, error) {
	supplier, err := repository.GetSupplierByID(id)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			return supplier, errs.ErrSupplierNotFound
		}
		return supplier, err
	}

	// Обновляем данные поставщика
	if updatedSupplier.Name != "" {
		supplier.Name = updatedSupplier.Name
	}
	if updatedSupplier.Email != "" {
		supplier.Email = updatedSupplier.Email
	}
	if updatedSupplier.Phone != "" {
		supplier.Phone = updatedSupplier.Phone
	}

	if err := repository.UpdateSupplierByID(supplier); err != nil {
		return supplier, err
	}
	return supplier, nil
}

func SoftDeleteSupplierByID(id uint) error {
	supplier, err := repository.GetSupplierByID(id)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			return errs.ErrSupplierNotFound
		}
		return err
	}

	if supplier.IsDeleted {
		return errs.ErrSupplierAlreadyDeleted
	}

	supplier.IsDeleted = true
	return repository.UpdateSupplierByID(supplier)
}

func RestoreSupplierByID(id uint) error {
	supplier, err := repository.GetSupplierIncludingSoftDeleted(id)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			return errs.ErrSupplierNotFound
		}
		return err
	}

	if !supplier.IsDeleted {
		return errs.ErrSupplierNotDeleted
	}

	supplier.IsDeleted = false
	return repository.UpdateSupplierByID(supplier)
}

func HardDeleteSupplierByID(id uint) error {
	return repository.HardDeleteSupplierByID(id)
}
