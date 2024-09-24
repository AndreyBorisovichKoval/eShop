// C:\GoProject\src\eShop\pkg\controllers\categories.go

package controllers

import (
	"eShop/errs"
	"eShop/logger"
	"eShop/models"
	"eShop/pkg/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateCategory
// @Summary Create a new category
// @Tags categories
// @Description Register a new category (Admin/Manager only)
// @ID create-category
// @Accept json
// @Produce json
// @Param input body models.Category true "Category data"
// @Success 201 {string} string "Category created successfully!!!"
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Failure 403 {object} ErrorResponse "Permission denied"
// @Failure 409 {object} ErrorResponse "Category already exists"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /categories [post]
// @Security ApiKeyAuth
func CreateCategory(c *gin.Context) {
	var category models.Category
	if err := c.BindJSON(&category); err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	logger.Info.Printf("IP: [%s] attempting to create category: %s\n", c.ClientIP(), category.Title)

	if err := service.CreateCategory(category); err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("IP: [%s] successfully created category: %s\n", c.ClientIP(), category.Title)
	c.JSON(http.StatusCreated, gin.H{"message": "Category created successfully!!!"})
}

// UpdateCategoryByID
// @Summary Update category by ID
// @Tags categories
// @Description Update category information by ID (Admin/Manager only)
// @ID update-category-by-id
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param input body models.Category true "Updated category information"
// @Success 200 {object} models.Category "Updated category"
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Failure 403 {object} ErrorResponse "Permission denied"
// @Failure 404 {object} ErrorResponse "Category not found"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /categories/{id} [patch]
// @Security ApiKeyAuth
func UpdateCategoryByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	var updatedCategory models.Category
	if err := c.BindJSON(&updatedCategory); err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	logger.Info.Printf("IP: [%s] requested to update category with ID: %d\n", c.ClientIP(), id)

	category, err := service.UpdateCategoryByID(uint(id), updatedCategory)
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("IP: [%s] successfully updated category with ID: %d\n", c.ClientIP(), id)
	c.JSON(http.StatusOK, category)
}

// SoftDeleteCategoryByID
// @Summary Soft delete category by ID
// @Tags categories
// @Description Soft delete category by ID (Admin/Manager only)
// @ID soft-delete-category-by-id
// @Param id path int true "Category ID"
// @Success 200 {string} string "Category soft deleted successfully!"
// @Failure 400 {object} ErrorResponse "Invalid ID"
// @Failure 403 {object} ErrorResponse "Permission denied"
// @Failure 404 {object} ErrorResponse "Category not found"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /categories/{id}/soft [delete]
// @Security ApiKeyAuth
func SoftDeleteCategoryByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	logger.Info.Printf("IP: [%s] requested to soft delete category with ID: %d\n", c.ClientIP(), id)

	if err := service.SoftDeleteCategoryByID(uint(id)); err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("IP: [%s] successfully soft deleted category with ID: %d\n", c.ClientIP(), id)
	c.JSON(http.StatusOK, gin.H{"message": "Category soft deleted successfully!"})
}

// RestoreCategoryByID
// @Summary Restore soft deleted category by ID
// @Tags categories
// @Description Restore a soft deleted category by ID (Admin/Manager only)
// @ID restore-category-by-id
// @Param id path int true "Category ID"
// @Success 200 {string} string "Category restored successfully!"
// @Failure 400 {object} ErrorResponse "Invalid ID"
// @Failure 403 {object} ErrorResponse "Permission denied"
// @Failure 404 {object} ErrorResponse "Category not found"
// @Failure 409 {object} ErrorResponse "Category not deleted"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /categories/{id}/restore [patch]
// @Security ApiKeyAuth
func RestoreCategoryByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	logger.Info.Printf("IP: [%s] requested to restore category with ID: %d\n", c.ClientIP(), id)

	if err := service.RestoreCategoryByID(uint(id)); err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("IP: [%s] successfully restored category with ID: %d\n", c.ClientIP(), id)
	c.JSON(http.StatusOK, gin.H{"message": "Category restored successfully!"})
}

// HardDeleteCategoryByID
// @Summary Permanently delete category by ID
// @Tags categories
// @Description Permanently delete category by ID (Admin/Manager only)
// @ID hard-delete-category-by-id
// @Param id path int true "Category ID"
// @Success 200 {string} string "Category permanently deleted successfully!"
// @Failure 400 {object} ErrorResponse "Invalid ID"
// @Failure 403 {object} ErrorResponse "Permission denied"
// @Failure 404 {object} ErrorResponse "Category not found"
// @Failure 409 {object} ErrorResponse "Category already deleted"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /categories/{id}/hard [delete]
// @Security ApiKeyAuth
func HardDeleteCategoryByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	logger.Info.Printf("IP: [%s] requested to hard delete category with ID: %d\n", c.ClientIP(), id)

	if err := service.HardDeleteCategoryByID(uint(id)); err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("IP: [%s] successfully hard deleted category with ID: %d\n", c.ClientIP(), id)
	c.JSON(http.StatusOK, gin.H{"message": "Category permanently deleted successfully!"})
}

// GetAllCategories
// @Summary Retrieve all active categories
// @Tags categories
// @Description Get a list of all active categories (Admin/Manager only)
// @ID get-all-categories
// @Produce json
// @Success 200 {array} models.Category "List of active categories"
// @Failure 403 {object} ErrorResponse "Permission denied"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /categories [get]
// @Security ApiKeyAuth
func GetAllCategories(c *gin.Context) {
	logger.Info.Printf("IP: [%s] requested list of all active categories\n", c.ClientIP())

	categories, err := service.GetAllCategories()
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("IP: [%s] successfully retrieved list of active categories\n", c.ClientIP())
	c.JSON(http.StatusOK, categories)
}

// GetAllDeletedCategories
// @Summary Retrieve all deleted categories
// @Tags categories
// @Description Get a list of all soft deleted categories (Admin/Manager only)
// @ID get-all-deleted-categories
// @Produce json
// @Success 200 {array} models.Category "List of deleted categories"
// @Failure 403 {object} ErrorResponse "Permission denied"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /categories/deleted [get]
// @Security ApiKeyAuth
func GetAllDeletedCategories(c *gin.Context) {
	logger.Info.Printf("IP: [%s] requested list of all deleted categories\n", c.ClientIP())

	categories, err := service.GetAllDeletedCategories()
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("IP: [%s] successfully retrieved list of deleted categories\n", c.ClientIP())
	c.JSON(http.StatusOK, categories)
}

// GetCategoryByID
// @Summary Retrieve category by ID
// @Tags categories
// @Description Get category information by ID (Admin/Manager only)
// @ID get-category-by-id
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} models.Category "Category information"
// @Failure 403 {object} ErrorResponse "Permission denied"
// @Failure 404 {object} ErrorResponse "Category not found"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /categories/{id} [get]
// @Security ApiKeyAuth
func GetCategoryByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	logger.Info.Printf("IP: [%s] requested category with ID: %d\n", c.ClientIP(), id)

	category, err := service.GetCategoryByID(uint(id))
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("IP: [%s] successfully retrieved category with ID: %d\n", c.ClientIP(), id)
	c.JSON(http.StatusOK, category)
}
