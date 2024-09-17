// C:\GoProject\src\eShop\pkg\controllers\errors.go
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
	case errors.Is(err, errs.ErrUsernameUniquenessFailed):
		// Ошибка уникальности или неверного пароля...
		c.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))

	case errors.Is(err, errs.ErrIncorrectUsernameOrPassword):
		// Ошибка неверного имени пользователя или пароля...
		c.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))

	case errors.Is(err, errs.ErrRecordNotFound):
		// Ошибка "Запись не найдена"...
		c.JSON(http.StatusNotFound, newErrorResponse(err.Error()))

	case errors.Is(err, errs.ErrPermissionDenied):
		// Ошибка "Доступ запрещен"...
		c.JSON(http.StatusForbidden, newErrorResponse(err.Error()))

	case errors.Is(err, errs.ErrUserNotFound):
		// Ошибка "Пользователь не найден"...
		c.JSON(http.StatusNotFound, newErrorResponse(err.Error()))

	case errors.Is(err, errs.ErrUsersNotFound):
		// Ошибка "Пользователи не найдены"...
		c.JSON(http.StatusNotFound, newErrorResponse(err.Error()))

	case errors.Is(err, errs.ErrUserAlreadyDeleted):
		// Ошибка "Пользователь уже удалён"...
		c.JSON(http.StatusConflict, newErrorResponse(err.Error()))

	case errors.Is(err, errs.ErrUserNotDeleted):
		// Ошибка "Пользователь не был удалён"...
		c.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))

	case errors.Is(err, errs.ErrUserAlreadyBlocked):
		// Ошибка "Пользователь уже заблокирован"...
		c.JSON(http.StatusConflict, newErrorResponse(err.Error()))

	case errors.Is(err, errs.ErrUserNotBlocked):
		// Ошибка "Пользователь не был заблокирован"...
		c.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))

	case errors.Is(err, errs.ErrUnauthorizedPasswordChange):
		// Ошибка "Попытка смены пароля без прав"...
		c.JSON(http.StatusForbidden, newErrorResponse(err.Error()))

	case errors.Is(err, errs.ErrIncorrectPassword):
		// Ошибка "Неверный старый пароль"...
		c.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))

	case errors.Is(err, errs.ErrSupplierNotFound):
		// Ошибка "Поставщик не найден"...
		c.JSON(http.StatusNotFound, newErrorResponse(err.Error()))

	case errors.Is(err, errs.ErrSupplierAlreadyDeleted):
		// Ошибка "Поставщик уже удалён"...
		c.JSON(http.StatusConflict, newErrorResponse(err.Error()))

	case errors.Is(err, errs.ErrSupplierNotDeleted):
		// Ошибка "Поставщик не был удалён"...
		c.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))

	case errors.Is(err, errs.ErrSupplierAlreadyExists):
		// Ошибка "Поставщик уже существует"...
		c.JSON(http.StatusConflict, newErrorResponse(err.Error()))

	case errors.Is(err, errs.ErrCategoryNotFound):
		// Ошибка "Категория не найдена"...
		c.JSON(http.StatusNotFound, newErrorResponse(err.Error()))

	case errors.Is(err, errs.ErrCategoryAlreadyDeleted):
		// Ошибка "Категория уже удалена"...
		c.JSON(http.StatusConflict, newErrorResponse(err.Error()))

	case errors.Is(err, errs.ErrCategoryNotDeleted):
		// Ошибка "Категория не была удалена"...
		c.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))

	case errors.Is(err, errs.ErrCategoryAlreadyExists):
		// Ошибка "Категория уже существует"...
		c.JSON(http.StatusConflict, newErrorResponse(err.Error()))

	case errors.Is(err, errs.ErrSupplierAlreadyExists):
		// Ошибка "Поставщик уже существует"...
		c.JSON(http.StatusConflict, newErrorResponse(err.Error()))

	case errors.Is(err, errs.ErrCategoryAlreadyExists):
		// Ошибка "Категория уже существует"...
		c.JSON(http.StatusConflict, newErrorResponse(err.Error()))

	case errors.Is(err, errs.ErrUniquenessViolation):
		// Ошибка нарушения уникальности...
		c.JSON(http.StatusConflict, newErrorResponse(err.Error()))

	case errors.Is(err, errs.ErrProductNotFound):
		// Ошибка "Продукт не найден"...
		c.JSON(http.StatusNotFound, newErrorResponse(err.Error()))

	case errors.Is(err, errs.ErrProductAlreadyExists):
		c.JSON(http.StatusConflict, newErrorResponse(err.Error()))

	case errors.Is(err, errs.ErrProductAlreadyDeleted):
		// Ошибка "Продукт уже удалён"...
		c.JSON(http.StatusConflict, newErrorResponse(err.Error()))

	case errors.Is(err, errs.ErrProductNotDeleted):
		// Ошибка "Продукт не был удалён"...
		c.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))

	// /
	case errors.Is(err, errs.ErrOrderNotFound):
		// Ошибка "Заказ не найден"...
		c.JSON(http.StatusNotFound, newErrorResponse(err.Error()))

	case errors.Is(err, errs.ErrOrderItemNotFound):
		// Ошибка "Элемент заказа не найден"...
		c.JSON(http.StatusNotFound, newErrorResponse(err.Error()))

	case errors.Is(err, errs.ErrInsufficientStock):
		// Ошибка "Недостаточно товара на складе"...
		c.JSON(http.StatusConflict, newErrorResponse(err.Error()))

	case errors.Is(err, errs.ErrUnauthorized):
		// Ошибка "Неавторизованный доступ"...
		c.JSON(http.StatusUnauthorized, newErrorResponse(err.Error()))

	// /
	default:
		// Внутренняя ошибка сервера...
		c.JSON(http.StatusInternalServerError, newErrorResponse(err.Error()))
	}
}

// ErrorResponse представляет структуру для обработки сообщений об ошибках...
type ErrorResponse struct {
	Error string `json:"error"` // Описание возникшей ошибки...
}

func newErrorResponse(message string) ErrorResponse {
	return ErrorResponse{
		Error: message,
	}
}
