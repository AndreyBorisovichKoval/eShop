// C:\GoProject\src\eShop\pkg\controllers\helpers.go
package controllers

import (
	"eShop/errs"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// handleError обрабатывает все ошибки, возникающие в процессе выполнения...
// Добавляет статус код к ним и сразу возвращает клиенту...
func handleError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, errs.ErrUsernameUniquenessFailed),
		errors.Is(err, errs.ErrIncorrectUsernameOrPassword):
		// Ошибка уникальности или неверного пароля...
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	case errors.Is(err, errs.ErrRecordNotFound):
		// Ошибка "Запись не найдена"...
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})

	case errors.Is(err, errs.ErrPermissionDenied):
		// Ошибка "Доступ запрещен"...
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})

	case errors.Is(err, errs.ErrUserNotFound):
		// Ошибка "Пользователь не найден"...
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})

	case errors.Is(err, errs.ErrUsersNotFound):
		// Ошибка "Пользователи не найдены"...
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})

	case errors.Is(err, errs.ErrUserAlreadyDeleted):
		// Ошибка "Пользователь уже удалён"...
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})

	case errors.Is(err, errs.ErrUserNotDeleted):
		// Ошибка "Пользователь не был удалён"...
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	case errors.Is(err, errs.ErrUserAlreadyBlocked):
		// Ошибка "Пользователь уже заблокирован"...
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})

	case errors.Is(err, errs.ErrUserNotBlocked):
		// Ошибка "Пользователь не был заблокирован"...
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	case errors.Is(err, errs.ErrUnauthorizedPasswordChange):
		// Ошибка "Попытка смены пароля без прав"...
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})

	case errors.Is(err, errs.ErrIncorrectPassword):
		// Ошибка "Неверный старый пароль"...
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	default:
		// Внутренняя ошибка сервера...
		c.JSON(http.StatusInternalServerError, gin.H{"error": errs.ErrSomethingWentWrong.Error()})
	}
}
