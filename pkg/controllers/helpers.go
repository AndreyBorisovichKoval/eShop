// C:\GoProject\src\eShop\pkg\controllers\helpers.go
package controllers

import (
	"eShop/errs"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// handleError обрабатывает все ошибки, возникающие в процессе выполнения
func handleError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, errs.ErrUsernameUniquenessFailed),
		errors.Is(err, errs.ErrIncorrectUsernameOrPassword):
		// Ошибка уникальности или неверного пароля
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	case errors.Is(err, errs.ErrRecordNotFound):
		// Ошибка "Запись не найдена"
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})

	case errors.Is(err, errs.ErrPermissionDenied):
		// Ошибка "Доступ запрещен"
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})

	default:
		// Здесь просто возвращаем внутреннюю ошибку клиенту
		c.JSON(http.StatusInternalServerError, gin.H{"error": errs.ErrSomethingWentWrong.Error()})
	}
}
