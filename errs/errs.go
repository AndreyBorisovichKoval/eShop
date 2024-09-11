// C:\GoProject\src\eShop\errs\errs.go

package errs

import "errors"

// Определение кастомных ошибок, используемых в приложении...
var (
	ErrPermissionDenied            = errors.New("ErrPermissionDenied")            // Ошибка: доступ запрещен...
	ErrValidationFailed            = errors.New("ErrValidationFailed")            // Ошибка: валидация данных не пройдена...
	ErrUsernameUniquenessFailed    = errors.New("ErrUsernameUniquenessFailed")    // Ошибка: имя пользователя должно быть уникальным...
	ErrIncorrectUsernameOrPassword = errors.New("ErrIncorrectUsernameOrPassword") // Ошибка: неверное имя пользователя или пароль...
	ErrRecordNotFound              = errors.New("ErrRecordNotFound")              // Ошибка: запись не найдена...
	ErrUserNotFound                = errors.New("ErrUserNotFound")                // Ошибка: пользователь не найден...
	ErrUsersNotFound               = errors.New("ErrUsersNotFound")               // Ошибка: пользователи не найдены...
	ErrUserAlreadyDeleted          = errors.New("ErrUserAlreadyDeleted")          // Ошибка: пользователь уже был удалён ранее...
	ErrUserNotDeleted              = errors.New("ErrUserNotDeleted")              // Ошибка: пользователь не был удалён...
	ErrSomethingWentWrong          = errors.New("ErrSomethingWentWrong")          // Ошибка: что-то пошло не так...
)
