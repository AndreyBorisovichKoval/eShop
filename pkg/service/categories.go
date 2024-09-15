// C:\GoProject\src\eShop\pkg\service\categories.go

package service

import (
	"eShop/errs"
	"eShop/logger"
	"eShop/models"
	"eShop/pkg/repository"
	"errors"
	"time"
)

// CreateCategory создает новую категорию
func CreateCategory(category models.Category) error {
	// Проверяем, что обязательные поля заполнены
	if category.Title == "" {
		return errs.ErrValidationFailed // Возвращаем ошибку валидации, если поле Title пусто
	}

	// Проверяем, существует ли уже категория с таким именем
	existingCategory, err := repository.GetCategoryByTitle(category.Title)
	if err != nil && err != errs.ErrRecordNotFound {
		return err
	}

	if existingCategory.ID > 0 {
		return errs.ErrCategoryAlreadyExists // Категория уже существует
	}

	// Создаем новую категорию через репозиторий
	if err := repository.CreateCategory(category); err != nil {
		return err
	}

	return nil
}

// UpdateCategoryByID обновляет данные категории по её ID
func UpdateCategoryByID(id uint, updatedCategory models.Category) (category models.Category, err error) {
	// Получаем существующую категорию
	category, err = repository.GetCategoryByID(id)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			return category, errs.ErrCategoryNotFound
		}
		return category, err
	}

	// Обновляем только изменённые поля
	if updatedCategory.Title != "" {
		category.Title = updatedCategory.Title
	}
	if updatedCategory.Description != "" {
		category.Description = updatedCategory.Description
	}

	// Используем функцию обновления в репозитории
	err = repository.UpdateCategoryByID(category)
	if err != nil {
		return category, err
	}

	return category, nil
}

// SoftDeleteCategoryByID помечает категорию как удалённую
func SoftDeleteCategoryByID(id uint) error {
	// Получаем существующую категорию
	category, err := repository.GetCategoryByID(id)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			return errs.ErrCategoryNotFound
		}
		return err
	}

	// Проверяем, не была ли категория уже удалена
	if category.IsDeleted {
		return errs.ErrCategoryAlreadyDeleted
	}

	// Помечаем категорию как удалённую
	category.IsDeleted = true
	currentTime := time.Now()
	category.DeletedAt = &currentTime

	// Используем общую функцию обновления для сохранения изменений
	if err := repository.UpdateCategoryByID(category); err != nil {
		return err
	}

	return nil
}

// RestoreCategoryByID восстанавливает мягко удалённую категорию
func RestoreCategoryByID(id uint) error {
	// Получаем категорию, включая мягко удалённые
	category, err := repository.GetCategoryIncludingSoftDeleted(id)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			return errs.ErrCategoryNotFound
		}
		return err
	}

	// Проверяем, была ли категория удалена
	if !category.IsDeleted {
		return errs.ErrCategoryNotDeleted
	}

	// Восстанавливаем категорию
	category.IsDeleted = false
	category.DeletedAt = nil

	// Используем общую функцию обновления для сохранения изменений
	if err := repository.UpdateCategoryByID(category); err != nil {
		return err
	}

	return nil
}

// HardDeleteCategoryByID выполняет жёсткое удаление категории
func HardDeleteCategoryByID(id uint) error {
	// Проверяем, существует ли категория, включая мягко удалённые
	category, err := repository.GetCategoryIncludingSoftDeleted(id)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			logger.Warning.Printf("[service.HardDeleteCategoryByID] category with ID: %v not found", id)
			return errs.ErrCategoryNotFound
		}
		return err
	}

	// Проверяем, не была ли категория уже жёстко удалена
	if category.IsDeleted {
		logger.Warning.Printf("[service.HardDeleteCategoryByID] category with ID: %v is already deleted", id)
		return errs.ErrCategoryAlreadyDeleted
	}

	// Выполняем жёсткое удаление
	if err := repository.HardDeleteCategoryByID(category.ID); err != nil {
		return err
	}

	return nil
}

// GetAllCategories получает все активные категории
func GetAllCategories() (categories []models.Category, err error) {
	categories, err = repository.GetAllActiveCategories()
	if err != nil {
		return nil, err
	}

	// Если категорий нет, логируем предупреждение
	if len(categories) == 0 {
		logger.Warning.Printf("[service.GetAllCategories] no categories found")
	}

	return categories, nil
}

// GetAllDeletedCategories получает все мягко удалённые категории
func GetAllDeletedCategories() (categories []models.Category, err error) {
	categories, err = repository.GetAllDeletedCategories()
	if err != nil {
		return nil, err
	}

	// Если удалённых категорий нет, логируем предупреждение
	if len(categories) == 0 {
		logger.Warning.Printf("[service.GetAllDeletedCategories] no deleted categories found")
	}

	return categories, nil
}

// GetCategoryByID получает категорию по её ID
func GetCategoryByID(id uint) (category models.Category, err error) {
	// Получаем категорию через репозиторий
	category, err = repository.GetCategoryByID(id)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			logger.Warning.Printf("[service.GetCategoryByID] category with ID %d not found", id)
			return category, errs.ErrCategoryNotFound
		}
		logger.Error.Printf("[service.GetCategoryByID] error getting category by ID: %v\n", err)
		return category, err
	}
	return category, nil
}
