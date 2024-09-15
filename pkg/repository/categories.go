// C:\GoProject\src\eShop\pkg\repository\categories.go

package repository

import (
	"eShop/db"
	"eShop/errs"
	"eShop/logger"
	"eShop/models"

	"gorm.io/gorm"
)

// CreateCategory создает новую категорию в базе данных
func CreateCategory(category models.Category) error {
	if err := db.GetDBConn().Create(&category).Error; err != nil {
		logger.Error.Printf("[repository.CreateCategory] error creating category: %v\n", err)
		return translateError(err)
	}
	return nil
}

// GetCategoryByTitle получает категорию по имени
func GetCategoryByTitle(title string) (category models.Category, err error) {
	err = db.GetDBConn().Where("title = ?", title).First(&category).Error
	if err != nil {
		logger.Error.Printf("[repository.GetCategoryByTitle] error getting category: %v\n", err)
		return category, translateError(err)
	}
	return category, nil
}

// GetCategoryByID получает категорию по её ID (только активные)
func GetCategoryByID(id uint) (category models.Category, err error) {
	err = db.GetDBConn().Where("id = ? AND is_deleted = ?", id, false).First(&category).Error
	if err != nil {
		logger.Error.Printf("[repository.GetCategoryByID] error getting category by id: %v\n", err)
		return category, translateError(err)
	}
	return category, nil
}

// GetCategoryIncludingSoftDeleted получает категорию по ID, включая мягко удалённые
func GetCategoryIncludingSoftDeleted(id uint) (category models.Category, err error) {
	err = db.GetDBConn().Unscoped().Where("id = ?", id).First(&category).Error
	if err != nil {
		logger.Error.Printf("[repository.GetCategoryIncludingSoftDeleted] error getting category: %v\n", err)
		return category, translateError(err)
	}
	return category, nil
}

// UpdateCategoryByID обновляет данные категории в базе данных
func UpdateCategoryByID(category models.Category) error {
	if err := db.GetDBConn().Save(&category).Error; err != nil {
		logger.Error.Printf("[repository.UpdateCategoryByID] error updating category: %v\n", err)
		return translateError(err)
	}
	return nil
}

// HardDeleteCategoryByID выполняет жёсткое удаление категории
func HardDeleteCategoryByID(id uint) error {
	var category models.Category

	// Проверяем существование категории перед удалением
	if err := db.GetDBConn().Unscoped().First(&category, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			logger.Warning.Printf("[repository.HardDeleteCategoryByID] category with ID: %v not found", id)
			return errs.ErrCategoryNotFound
		}
		logger.Error.Printf("[repository.HardDeleteCategoryByID] error retrieving category with ID: %v, error: %v", id, err)
		return translateError(err)
	}

	// Жёсткое удаление категории
	if err := db.GetDBConn().Unscoped().Delete(&category).Error; err != nil {
		logger.Error.Printf("[repository.HardDeleteCategoryByID] error hard deleting category with ID: %v, error: %v", id, err)
		return translateError(err)
	}

	return nil
}

// GetAllActiveCategories получает все активные категории (не удалённые)
func GetAllActiveCategories() (categories []models.Category, err error) {
	err = db.GetDBConn().Where("is_deleted = ?", false).Find(&categories).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllActiveCategories] error getting all active categories: %v\n", err)
		return nil, translateError(err)
	}
	return categories, nil
}

// GetAllDeletedCategories получает все мягко удалённые категории
func GetAllDeletedCategories() (categories []models.Category, err error) {
	err = db.GetDBConn().Where("is_deleted = ?", true).Find(&categories).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllDeletedCategories] error getting all deleted categories: %v\n", err)
		return nil, translateError(err)
	}
	return categories, nil
}
