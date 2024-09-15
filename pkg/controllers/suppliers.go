// C:\GoProject\src\eShop\pkg\controllers\suppliers.go

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

// CreateSupplier
// @Summary Create a new supplier
// @Tags suppliers
// @Description Register a new supplier (Admin only)
// @ID create-supplier
// @Accept json
// @Produce json
// @Param input body models.Supplier true "Supplier data"
// @Success 201 {string} string "Supplier created successfully!!!"
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Failure 403 {object} ErrorResponse "Permission denied"
// @Failure 409 {object} ErrorResponse "Supplier already exists"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /suppliers [post]
// @Security ApiKeyAuth
func CreateSupplier(c *gin.Context) {
	userRole := c.GetString(userRoleCtx)

	if userRole != "Admin" && userRole != "Manager" {
		handleError(c, errs.ErrPermissionDenied)
		return
	}

	var supplier models.Supplier
	if err := c.BindJSON(&supplier); err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	logger.Info.Printf("IP: [%s] attempting to create supplier: %s\n", c.ClientIP(), supplier.Name)

	if err := service.CreateSupplier(supplier); err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("IP: [%s] successfully created supplier: %s\n", c.ClientIP(), supplier.Name)
	c.JSON(http.StatusCreated, gin.H{"message": "Supplier created successfully!!!"})
}

// UpdateSupplierByID
// @Summary Update supplier by ID
// @Tags suppliers
// @Description Update supplier information by ID (Admin/Manager only)
// @ID update-supplier-by-id
// @Accept json
// @Produce json
// @Param id path int true "Supplier ID"
// @Param input body models.Supplier true "Updated supplier information"
// @Success 200 {object} models.Supplier "Updated supplier"
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Failure 403 {object} ErrorResponse "Permission denied"
// @Failure 404 {object} ErrorResponse "Supplier not found"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /suppliers/{id} [patch]
// @Security ApiKeyAuth
func UpdateSupplierByID(c *gin.Context) {
	userRole := c.GetString(userRoleCtx)

	if userRole != "Admin" && userRole != "Manager" {
		handleError(c, errs.ErrPermissionDenied)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	var updatedSupplier models.Supplier
	if err := c.BindJSON(&updatedSupplier); err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	logger.Info.Printf("IP: [%s] requested to update supplier with ID: %d\n", c.ClientIP(), id)

	supplier, err := service.UpdateSupplierByID(uint(id), updatedSupplier)
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("IP: [%s] successfully updated supplier with ID: %d\n", c.ClientIP(), id)
	c.JSON(http.StatusOK, supplier)
}

// SoftDeleteSupplierByID
// @Summary Soft delete supplier by ID
// @Tags suppliers
// @Description Soft delete supplier by ID (Admin/Manager only)
// @ID soft-delete-supplier-by-id
// @Param id path int true "Supplier ID"
// @Success 200 {string} string "Supplier soft deleted successfully!"
// @Failure 400 {object} ErrorResponse "Invalid ID"
// @Failure 403 {object} ErrorResponse "Permission denied"
// @Failure 404 {object} ErrorResponse "Supplier not found"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /suppliers/{id}/soft [delete]
// @Security ApiKeyAuth
func SoftDeleteSupplierByID(c *gin.Context) {
	userRole := c.GetString(userRoleCtx)

	if userRole != "Admin" && userRole != "Manager" {
		handleError(c, errs.ErrPermissionDenied)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	logger.Info.Printf("IP: [%s] requested to soft delete supplier with ID: %d\n", c.ClientIP(), id)

	if err := service.SoftDeleteSupplierByID(uint(id)); err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("IP: [%s] successfully soft deleted supplier with ID: %d\n", c.ClientIP(), id)
	c.JSON(http.StatusOK, gin.H{"message": "Supplier soft deleted successfully!"})
}

// RestoreSupplierByID
// @Summary Restore soft deleted supplier by ID
// @Tags suppliers
// @Description Restore a soft deleted supplier by ID (Admin/Manager only)
// @ID restore-supplier-by-id
// @Param id path int true "Supplier ID"
// @Success 200 {string} string "Supplier restored successfully!"
// @Failure 400 {object} ErrorResponse "Invalid ID"
// @Failure 403 {object} ErrorResponse "Permission denied"
// @Failure 404 {object} ErrorResponse "Supplier not found"
// @Failure 409 {object} ErrorResponse "Supplier not deleted"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /suppliers/{id}/restore [patch]
// @Security ApiKeyAuth
func RestoreSupplierByID(c *gin.Context) {
	userRole := c.GetString(userRoleCtx)

	if userRole != "Admin" && userRole != "Manager" {
		handleError(c, errs.ErrPermissionDenied)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	logger.Info.Printf("IP: [%s] requested to restore supplier with ID: %d\n", c.ClientIP(), id)

	if err := service.RestoreSupplierByID(uint(id)); err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("IP: [%s] successfully restored supplier with ID: %d\n", c.ClientIP(), id)
	c.JSON(http.StatusOK, gin.H{"message": "Supplier restored successfully!"})
}

// /

// HardDeleteSupplierByID
// @Summary Hard delete supplier by ID
// @Tags suppliers
// @Description Permanently delete supplier by ID
// @ID hard-delete-supplier-by-id
// @Param id path int true "Supplier ID"
// @Success 200 {string} string "Supplier hard deleted successfully!"
// @Failure 400 {object} ErrorResponse "Invalid ID"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /suppliers/{id}/hard [delete]
func HardDeleteSupplierByID(c *gin.Context) {
	// Получаем роль пользователя из контекста...
	userRole := c.GetString(userRoleCtx)

	// Проверяем доступ (только Admin или Manager)...
	if userRole != "Admin" && userRole != "Manager" {
		handleError(c, errs.ErrPermissionDenied)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	// Логируем запрос на удаление поставщика...
	logger.Info.Printf("IP: [%s] requested hard delete for supplier with ID: %d\n", c.ClientIP(), id)

	if err := service.HardDeleteSupplierByID(uint(id)); err != nil {
		handleError(c, err)
		return
	}

	// Логируем успешное удаление...
	logger.Info.Printf("IP: [%s] successfully hard deleted supplier with ID: %d\n", c.ClientIP(), id)
	c.JSON(http.StatusOK, gin.H{"message": "Supplier hard deleted successfully!"})
}

// GetAllSuppliers
// @Summary Retrieve all suppliers
// @Tags suppliers
// @Description Get a list of all registered suppliers
// @ID get-all-suppliers
// @Produce json
// @Success 200 {array} models.Supplier "List of suppliers"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /suppliers [get]
func GetAllSuppliers(c *gin.Context) {
	// Логируем запрос на получение списка всех поставщиков...
	logger.Info.Printf("IP: [%s] requested list of all suppliers\n", c.ClientIP())

	// Получаем список поставщиков через сервис...
	suppliers, err := service.GetAllSuppliers()
	if err != nil {
		handleError(c, err)
		return
	}

	// Логируем успешный ответ...
	logger.Info.Printf("IP: [%s] successfully retrieved list of suppliers\n", c.ClientIP())
	c.JSON(http.StatusOK, suppliers)
}

// GetAllDeletedSuppliers
// @Summary Retrieve all deleted suppliers
// @Tags suppliers
// @Description Get a list of all soft deleted suppliers (Admin only)
// @ID get-all-deleted-suppliers
// @Produce json
// @Success 200 {array} models.Supplier "List of deleted suppliers"
// @Failure 403 {object} ErrorResponse "Permission denied"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /suppliers/deleted [get]
// @Security ApiKeyAuth
func GetAllDeletedSuppliers(c *gin.Context) {
	// Получаем роль пользователя из контекста
	userRole := c.GetString(userRoleCtx)

	// Проверяем доступ (только Admin или Manager)
	if userRole != "Admin" && userRole != "Manager" {
		handleError(c, errs.ErrPermissionDenied)
		return
	}

	// Логируем запрос...
	logger.Info.Printf("IP: [%s] requested list of all deleted suppliers\n", c.ClientIP())

	// Вызываем сервис для получения списка удалённых поставщиков...
	suppliers, err := service.GetAllDeletedSuppliers()
	if err != nil {
		handleError(c, err)
		return
	}

	// Логируем успешное получение списка...
	logger.Info.Printf("IP: [%s] successfully retrieved list of deleted suppliers\n", c.ClientIP())

	// Возвращаем список удалённых поставщиков клиенту...
	c.JSON(http.StatusOK, suppliers)
}

// GetSupplierByID
// @Summary Retrieve supplier by ID
// @Tags suppliers
// @Description Get supplier information by ID
// @ID get-supplier-by-id
// @Produce json
// @Param id path int true "Supplier ID"
// @Success 200 {object} models.Supplier "Supplier information"
// @Failure 400 {object} ErrorResponse "Invalid ID"
// @Failure 404 {object} ErrorResponse "Supplier not found"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /suppliers/{id} [get]
func GetSupplierByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// Логируем ошибку в случае некорректного ID...
		logger.Error.Printf("IP: [%s] invalid supplier_id: %s\n", c.ClientIP(), c.Param("id"))
		handleError(c, errs.ErrValidationFailed)
		return
	}

	// Логируем запрос на получение поставщика по ID...
	logger.Info.Printf("IP: [%s] requested supplier with ID: %d\n", c.ClientIP(), id)

	// Получаем поставщика через сервис...
	supplier, err := service.GetSupplierByID(uint(id))
	if err != nil {
		handleError(c, err)
		return
	}

	// Логируем успешный ответ...
	logger.Info.Printf("IP: [%s] successfully retrieved supplier with ID: %d\n", c.ClientIP(), id)
	c.JSON(http.StatusOK, supplier)
}
