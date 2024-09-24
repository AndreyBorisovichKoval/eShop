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
	// Ошибки аутентификации
	case errors.Is(err, errs.ErrEmptyAuthHeader):
		c.JSON(http.StatusUnauthorized, newErrorResponse(err.Error()))
	case errors.Is(err, errs.ErrInvalidAuthHeader):
		c.JSON(http.StatusUnauthorized, newErrorResponse(err.Error()))
	case errors.Is(err, errs.ErrTokenParsingFailed):
		c.JSON(http.StatusUnauthorized, newErrorResponse(err.Error()))
	case errors.Is(err, errs.ErrUserNotAuthenticated):
		c.JSON(http.StatusUnauthorized, newErrorResponse(err.Error()))
	case errors.Is(err, errs.ErrPasswordResetRequired):
		c.JSON(http.StatusForbidden, newErrorResponse(err.Error()))

	// Ошибки разрешений
	case errors.Is(err, errs.ErrPermissionDenied):
		c.JSON(http.StatusForbidden, newErrorResponse(err.Error()))
	case errors.Is(err, errs.ErrPermissionDeniedOnlyForAdmin):
		c.JSON(http.StatusForbidden, newErrorResponse(err.Error()))
	case errors.Is(err, errs.ErrPermissionDeniedOnlyForAdminOrManager):
		c.JSON(http.StatusForbidden, newErrorResponse(err.Error()))
	case errors.Is(err, errs.ErrUserBlocked):
		c.JSON(http.StatusForbidden, newErrorResponse(err.Error()))

	// Ошибки валидации
	case errors.Is(err, errs.ErrValidationFailed):
		c.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
	case errors.Is(err, errs.ErrUsernameUniquenessFailed):
		c.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
	case errors.Is(err, errs.ErrIncorrectUsernameOrPassword):
		c.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
	case errors.Is(err, errs.ErrIncorrectPassword):
		c.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))

	// Ошибки пользователей
	case errors.Is(err, errs.ErrRecordNotFound):
		c.JSON(http.StatusNotFound, newErrorResponse(err.Error()))
	case errors.Is(err, errs.ErrUserNotFound):
		c.JSON(http.StatusNotFound, newErrorResponse(err.Error()))
	case errors.Is(err, errs.ErrUsersNotFound):
		c.JSON(http.StatusNotFound, newErrorResponse(err.Error()))
	case errors.Is(err, errs.ErrUserAlreadyDeleted):
		c.JSON(http.StatusConflict, newErrorResponse(err.Error()))
	case errors.Is(err, errs.ErrUserNotDeleted):
		c.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
	case errors.Is(err, errs.ErrUserAlreadyBlocked):
		c.JSON(http.StatusConflict, newErrorResponse(err.Error()))
	case errors.Is(err, errs.ErrUserNotBlocked):
		c.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))

	// Ошибки поставщиков
	case errors.Is(err, errs.ErrSupplierNotFound):
		c.JSON(http.StatusNotFound, newErrorResponse(err.Error()))
	case errors.Is(err, errs.ErrSupplierAlreadyExists):
		c.JSON(http.StatusConflict, newErrorResponse(err.Error()))
	case errors.Is(err, errs.ErrSupplierAlreadyDeleted):
		c.JSON(http.StatusConflict, newErrorResponse(err.Error()))
	case errors.Is(err, errs.ErrSupplierNotDeleted):
		c.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))

	// Ошибки категорий
	case errors.Is(err, errs.ErrCategoryNotFound):
		c.JSON(http.StatusNotFound, newErrorResponse(err.Error()))
	case errors.Is(err, errs.ErrCategoryAlreadyExists):
		c.JSON(http.StatusConflict, newErrorResponse(err.Error()))
	case errors.Is(err, errs.ErrCategoryAlreadyDeleted):
		c.JSON(http.StatusConflict, newErrorResponse(err.Error()))
	case errors.Is(err, errs.ErrCategoryNotDeleted):
		c.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))

	// Ошибки продуктов
	case errors.Is(err, errs.ErrProductNotFound):
		c.JSON(http.StatusNotFound, newErrorResponse(err.Error()))
	case errors.Is(err, errs.ErrProductAlreadyExists):
		c.JSON(http.StatusConflict, newErrorResponse(err.Error()))
	case errors.Is(err, errs.ErrProductAlreadyDeleted):
		c.JSON(http.StatusConflict, newErrorResponse(err.Error()))
	case errors.Is(err, errs.ErrProductNotDeleted):
		c.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
	case errors.Is(err, errs.ErrProductNotWeightBased):
		c.JSON(http.StatusConflict, newErrorResponse(err.Error()))

	// Ошибки заказов
	case errors.Is(err, errs.ErrOrderNotFound):
		c.JSON(http.StatusNotFound, newErrorResponse(err.Error()))
	case errors.Is(err, errs.ErrOrderItemNotFound):
		c.JSON(http.StatusNotFound, newErrorResponse(err.Error()))
	case errors.Is(err, errs.ErrInsufficientStock):
		c.JSON(http.StatusConflict, newErrorResponse(err.Error()))
	case errors.Is(err, errs.ErrOrderAlreadyPaid):
		c.JSON(http.StatusConflict, newErrorResponse(err.Error()))

	// Общие ошибки
	case errors.Is(err, errs.ErrUnauthorized):
		c.JSON(http.StatusUnauthorized, newErrorResponse(err.Error()))
	case errors.Is(err, errs.ErrServerError):
		c.JSON(http.StatusInternalServerError, newErrorResponse(err.Error()))

	// Ошибки, связанные с удалением заказов
	case errors.Is(err, errs.ErrCannotDeletePaidOrder):
		c.JSON(http.StatusConflict, newErrorResponse(err.Error()))
	case errors.Is(err, errs.ErrCannotDeletePaidOrderItem):
		c.JSON(http.StatusConflict, newErrorResponse(err.Error()))
	case errors.Is(err, errs.ErrCannotAddToPaidOrder):
		c.JSON(http.StatusConflict, newErrorResponse(err.Error()))

	// Внутренняя ошибка сервера
	default:
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
