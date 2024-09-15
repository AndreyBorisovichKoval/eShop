// C:\GoProject\src\eShop\pkg\repository\suppliers.go

package repository

import (
	"eShop/db"
	"eShop/logger"
	"eShop/models"
)

// CreateSupplier добавляет нового поставщика в базу данных...
func CreateSupplier(supplier models.Supplier) error {
	err := db.GetDBConn().Create(&supplier).Error
	if err != nil {
		logger.Error.Printf("[repository.CreateSupplier] error creating supplier: %v\n", err)
		return translateError(err)
	}
	return nil
}

// UpdateSupplierByID обновляет данные поставщика в базе данных...
func UpdateSupplierByID(supplier models.Supplier) error {
	err := db.GetDBConn().Save(&supplier).Error
	if err != nil {
		logger.Error.Printf("[repository.UpdateSupplierByID] error updating supplier: %v\n", err)
		return translateError(err)
	}
	return nil
}

// SoftDeleteSupplierByID помечает поставщика как удалённого...
func SoftDeleteSupplierByID(id uint) error {
	err := db.GetDBConn().Model(&models.Supplier{}).Where("id = ?", id).Update("is_deleted", true).Error
	if err != nil {
		logger.Error.Printf("[repository.SoftDeleteSupplierByID] error soft deleting supplier: %v\n", err)
		return translateError(err)
	}
	return nil
}

// RestoreSupplierByID восстанавливает удалённого поставщика...
func RestoreSupplierByID(id uint) error {
	err := db.GetDBConn().Model(&models.Supplier{}).Where("id = ?", id).Update("is_deleted", false).Error
	if err != nil {
		logger.Error.Printf("[repository.RestoreSupplierByID] error restoring supplier: %v\n", err)
		return translateError(err)
	}
	return nil
}

// HardDeleteSupplierByID удаляет поставщика из базы данных...
func HardDeleteSupplierByID(id uint) error {
	err := db.GetDBConn().Unscoped().Where("id = ?", id).Delete(&models.Supplier{}).Error
	if err != nil {
		logger.Error.Printf("[repository.HardDeleteSupplierByID] error hard deleting supplier: %v\n", err)
		return translateError(err)
	}
	return nil
}

// GetAllSuppliers получает список всех поставщиков из базы данных...
func GetAllSuppliers() (suppliers []models.Supplier, err error) {
	err = db.GetDBConn().Where("is_deleted = ?", false).Find(&suppliers).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllSuppliers] error getting all suppliers: %v\n", err)
		return nil, translateError(err)
	}
	return suppliers, nil
}

// GetAllDeletedSuppliers получает всех удалённых поставщиков из базы данных...
func GetAllDeletedSuppliers() (suppliers []models.Supplier, err error) {
	// Выполняем запрос к базе данных для получения всех удалённых поставщиков...
	err = db.GetDBConn().Where("is_deleted = ?", true).Find(&suppliers).Error
	if err != nil {
		// Логируем ошибку при получении списка удалённых поставщиков...
		logger.Error.Printf("[repository.GetAllDeletedSuppliers] error getting deleted suppliers: %v\n", err)
		return nil, translateError(err)
	}

	// Возвращаем список поставщиков...
	return suppliers, nil
}

// GetSupplierByID получает поставщика по ID...
func GetSupplierByID(id uint) (supplier models.Supplier, err error) {
	err = db.GetDBConn().Where("id = ? AND is_deleted = ?", id, false).First(&supplier).Error
	if err != nil {
		logger.Error.Printf("[repository.GetSupplierByID] error getting supplier by id: %v\n", err)
		return supplier, translateError(err)
	}
	return supplier, nil
}
