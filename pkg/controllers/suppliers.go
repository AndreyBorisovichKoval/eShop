// C:\GoProject\src\eShop\pkg\controllers\suppliers.go

package controllers

import (
	"eShop/errs"
	"eShop/models"
	"eShop/pkg/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateSupplier
// @Summary Create a new supplier
// @Tags suppliers
// @Description Add a new supplier to the system
// @ID create-supplier
// @Accept json
// @Produce json
// @Param input body models.Supplier true "Supplier information"
// @Success 201 {object} models.Supplier "Created supplier"
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /suppliers [post]
// @Security ApiKeyAuth
func CreateSupplier(c *gin.Context) {
	var supplier models.Supplier
	if err := c.BindJSON(&supplier); err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	if err := service.CreateSupplier(supplier); err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, supplier)
}

// GetAllSuppliers
// @Summary Get all suppliers
// @Tags suppliers
// @Description Retrieve all suppliers from the system
// @ID get-all-suppliers
// @Produce json
// @Success 200 {array} models.Supplier "List of suppliers"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /suppliers [get]
// @Security ApiKeyAuth
func GetAllSuppliers(c *gin.Context) {
	suppliers, err := service.GetAllSuppliers()
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, suppliers)
}

// GetSupplierByID
// @Summary Get supplier by ID
// @Tags suppliers
// @Description Retrieve supplier information by ID
// @ID get-supplier-by-id
// @Produce json
// @Param id path int true "Supplier ID"
// @Success 200 {object} models.Supplier "Supplier information"
// @Failure 400 {object} ErrorResponse "Invalid ID"
// @Failure 404 {object} ErrorResponse "Supplier not found"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /suppliers/{id} [get]
// @Security ApiKeyAuth
func GetSupplierByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	supplier, err := service.GetSupplierByID(uint(id))
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, supplier)
}

// UpdateSupplierByID
// @Summary Update supplier by ID
// @Tags suppliers
// @Description Update supplier information by ID
// @ID update-supplier-by-id
// @Accept json
// @Produce json
// @Param id path int true "Supplier ID"
// @Param input body models.Supplier true "Updated supplier information"
// @Success 200 {object} models.Supplier "Updated supplier"
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Failure 404 {object} ErrorResponse "Supplier not found"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /suppliers/{id} [patch]
// @Security ApiKeyAuth
func UpdateSupplierByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	var supplier models.Supplier
	if err := c.BindJSON(&supplier); err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	updatedSupplier, err := service.UpdateSupplierByID(uint(id), supplier)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, updatedSupplier)
}

// SoftDeleteSupplierByID
// @Summary Soft delete supplier by ID
// @Tags suppliers
// @Description Mark a supplier as deleted (soft delete)
// @ID soft-delete-supplier-by-id
// @Param id path int true "Supplier ID"
// @Success 200 {string} string "Supplier soft deleted successfully!"
// @Failure 400 {object} ErrorResponse "Invalid ID"
// @Failure 404 {object} ErrorResponse "Supplier not found"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /suppliers/{id}/soft [delete]
// @Security ApiKeyAuth
func SoftDeleteSupplierByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	if err := service.SoftDeleteSupplierByID(uint(id)); err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Supplier soft deleted successfully!"})
}

// RestoreSupplierByID
// @Summary Restore supplier by ID
// @Tags suppliers
// @Description Restore a soft deleted supplier by ID
// @ID restore-supplier-by-id
// @Param id path int true "Supplier ID"
// @Success 200 {string} string "Supplier restored successfully!"
// @Failure 400 {object} ErrorResponse "Invalid ID"
// @Failure 404 {object} ErrorResponse "Supplier not found"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /suppliers/{id}/restore [patch]
// @Security ApiKeyAuth
func RestoreSupplierByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	if err := service.RestoreSupplierByID(uint(id)); err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Supplier restored successfully!"})
}

// HardDeleteSupplierByID
// @Summary Hard delete supplier by ID
// @Tags suppliers
// @Description Permanently delete supplier by ID
// @ID hard-delete-supplier-by-id
// @Param id path int true "Supplier ID"
// @Success 200 {string} string "Supplier hard deleted successfully!"
// @Failure 400 {object} ErrorResponse "Invalid ID"
// @Failure 404 {object} ErrorResponse "Supplier not found"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /suppliers/{id}/hard [delete]
// @Security ApiKeyAuth
func HardDeleteSupplierByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	if err := service.HardDeleteSupplierByID(uint(id)); err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Supplier hard deleted successfully!"})
}
