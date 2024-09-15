// C:\GoProject\src\eShop\pkg\repository\suppliers.go

package repository

import (
	"eShop/db"
	"eShop/logger"
	"eShop/models"
)

// CreateSupplier добавляет нового поставщика в базу данных...
func CreateSupplier(supplier models.Supplier) error {
	if err := db.GetDBConn().Create(&supplier).Error; err != nil {
		logger.Error.Printf("[repository.CreateSupplier] error creating supplier: %v\n", err)
		return translateError(err)
	}
	return nil
}

// GetAllSuppliers получает список всех поставщиков из базы данных...
func GetAllSuppliers() ([]models.Supplier, error) {
	var suppliers []models.Supplier
	if err := db.GetDBConn().Where("is_deleted = ?", false).Find(&suppliers).Error; err != nil {
		logger.Error.Printf("[repository.GetAllSuppliers] error getting suppliers: %v\n", err)
		return nil, translateError(err)
	}
	return suppliers, nil
}

// GetSupplierByID получает поставщика по ID...
func GetSupplierByID(id uint) (models.Supplier, error) {
	var supplier models.Supplier
	if err := db.GetDBConn().Where("id = ? AND is_deleted = ?", id, false).First(&supplier).Error; err != nil {
		logger.Error.Printf("[repository.GetSupplierByID] error getting supplier by id: %v\n", err)
		return supplier, translateError(err)
	}
	return supplier, nil
}

// GetSupplierIncludingSoftDeleted получает поставщика по ID, включая удалённых...
func GetSupplierIncludingSoftDeleted(id uint) (models.Supplier, error) {
	var supplier models.Supplier
	if err := db.GetDBConn().Unscoped().Where("id = ?", id).First(&supplier).Error; err != nil {
		logger.Error.Printf("[repository.GetSupplierIncludingSoftDeleted] error getting supplier by id: %v\n", err)
		return supplier, translateError(err)
	}
	return supplier, nil
}

// UpdateSupplierByID обновляет данные поставщика в базе данных...
func UpdateSupplierByID(supplier models.Supplier) error {
	if err := db.GetDBConn().Save(&supplier).Error; err != nil {
		logger.Error.Printf("[repository.UpdateSupplierByID] error updating supplier: %v\n", err)
		return translateError(err)
	}
	return nil
}

// HardDeleteSupplierByID удаляет поставщика из базы данных...
func HardDeleteSupplierByID(id uint) error {
	if err := db.GetDBConn().Unscoped().Where("id = ?", id).Delete(&models.Supplier{}).Error; err != nil {
		logger.Error.Printf("[repository.HardDeleteSupplierByID] error deleting supplier by id: %v\n", err)
		return translateError(err)
	}
	return nil
}
