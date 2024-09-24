// C:\GoProject\src\eShop\pkg\repository\suppliers.go

package repository

import (
	"eShop/db"
	"eShop/logger"
	"eShop/models"
)

// CreateSupplier создает нового поставщика в базе данных
func CreateSupplier(supplier models.Supplier) error {
	if err := db.GetDBConn().Create(&supplier).Error; err != nil {
		logger.Error.Printf("[repository.CreateSupplier] error creating supplier: %v\n", err)
		return translateError(err)
	}
	logger.Info.Printf("[repository.CreateSupplier] supplier created successfully with ID: %d\n", supplier.ID) // Лог успешного создания поставщика
	return nil
}

// GetSupplierByTitleOrEmail получает поставщика по имени или email
func GetSupplierByTitleOrEmail(title, email string) (supplier models.Supplier, err error) {
	err = db.GetDBConn().Where("title = ? OR email = ?", title, email).First(&supplier).Error
	if err != nil {
		logger.Error.Printf("[repository.GetSupplierByTitleOrEmail] error getting supplier: %v\n", err)
		return supplier, translateError(err)
	}
	logger.Info.Printf("[repository.GetSupplierByTitleOrEmail] supplier retrieved successfully with title: %s or email: %s\n", title, email) // Лог успешного получения поставщика
	return supplier, nil
}

// UpdateSupplierByID обновляет данные поставщика в базе данных
func UpdateSupplierByID(supplier models.Supplier) error {
	if err := db.GetDBConn().Save(&supplier).Error; err != nil {
		logger.Error.Printf("[repository.UpdateSupplierByID] error updating supplier: %v\n", err)
		return translateError(err)
	}
	logger.Info.Printf("[repository.UpdateSupplierByID] supplier updated successfully with ID: %d\n", supplier.ID) // Лог успешного обновления поставщика
	return nil
}

// GetSupplierByID получает поставщика по его ID
func GetSupplierByID(id uint) (supplier models.Supplier, err error) {
	err = db.GetDBConn().Where("id = ? AND is_deleted = ?", id, false).First(&supplier).Error
	if err != nil {
		logger.Error.Printf("[repository.GetSupplierByID] error getting supplier by id: %v\n", err)
		return supplier, translateError(err)
	}
	logger.Info.Printf("[repository.GetSupplierByID] supplier retrieved successfully with ID: %d\n", id) // Лог успешного получения поставщика
	return supplier, nil
}

// GetSupplierIncludingSoftDeleted получает поставщика, включая мягко удалённых
func GetSupplierIncludingSoftDeleted(id uint) (supplier models.Supplier, err error) {
	err = db.GetDBConn().Unscoped().Where("id = ?", id).First(&supplier).Error
	if err != nil {
		logger.Error.Printf("[repository.GetSupplierIncludingSoftDeleted] error getting supplier: %v\n", err)
		return supplier, translateError(err)
	}
	logger.Info.Printf("[repository.GetSupplierIncludingSoftDeleted] supplier retrieved successfully including soft deleted with ID: %d\n", id) // Лог успешного получения мягко удалённого поставщика
	return supplier, nil
}

// HardDeleteSupplierByID выполняет жёсткое удаление поставщика
func HardDeleteSupplierByID(id uint) error {
	var supplier models.Supplier

	if err := db.GetDBConn().Unscoped().Delete(&supplier, id).Error; err != nil {
		logger.Error.Printf("[repository.HardDeleteSupplierByID] error hard deleting supplier with ID: %v, error: %v", id, err)
		return translateError(err)
	}
	logger.Info.Printf("[repository.HardDeleteSupplierByID] supplier hard deleted successfully with ID: %d\n", id) // Лог успешного жёсткого удаления
	return nil
}

// GetAllActiveSuppliers получает всех активных поставщиков (не удалённых)
func GetAllActiveSuppliers() (suppliers []models.Supplier, err error) {
	err = db.GetDBConn().Where("is_deleted = ?", false).Find(&suppliers).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllActiveSuppliers] error getting all active suppliers: %v\n", err)
		return nil, translateError(err)
	}
	logger.Info.Printf("[repository.GetAllActiveSuppliers] active suppliers retrieved successfully\n") // Лог успешного получения активных поставщиков
	return suppliers, nil
}

// GetAllDeletedSuppliers получает всех мягко удалённых поставщиков
func GetAllDeletedSuppliers() (suppliers []models.Supplier, err error) {
	err = db.GetDBConn().Where("is_deleted = ?", true).Find(&suppliers).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllDeletedSuppliers] error getting all deleted suppliers: %v\n", err)
		return nil, translateError(err)
	}
	logger.Info.Printf("[repository.GetAllDeletedSuppliers] deleted suppliers retrieved successfully\n") // Лог успешного получения удалённых поставщиков
	return suppliers, nil
}

