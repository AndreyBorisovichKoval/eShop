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
	ErrUserAlreadyBlocked          = errors.New("ErrUserAlreadyBlocked")          // Ошибка: пользователь уже заблокирован...
	ErrUserNotBlocked              = errors.New("ErrUserNotBlocked")              // Ошибка: пользователь не был заблокирован...
	ErrUnauthorizedPasswordChange  = errors.New("ErrUnauthorizedPasswordChange")  // Ошибка: попытка смены пароля без прав...
	ErrIncorrectPassword           = errors.New("ErrIncorrectPassword")           // Ошибка: неверный старый пароль...

	// Ошибки для сущности Поставщика...
	ErrSupplierAlreadyExists  = errors.New("ErrSupplierAlreadyExists")  // Ошибка: поставщик с таким именем или email уже существует...
	ErrSupplierNotFound       = errors.New("ErrSupplierNotFound")       // Ошибка: поставщик не найден...
	ErrSupplierAlreadyDeleted = errors.New("ErrSupplierAlreadyDeleted") // Ошибка: поставщик уже удалён...
	ErrSupplierNotDeleted     = errors.New("ErrSupplierNotDeleted")     // Ошибка: поставщик не был удалён...

	// Ошибки для сущности Категории...
	ErrCategoryAlreadyExists  = errors.New("ErrCategoryAlreadyExists")  // Ошибка: категория с таким названием уже существует...
	ErrCategoryNotFound       = errors.New("ErrCategoryNotFound")       // Ошибка: категория не найдена...
	ErrCategoryAlreadyDeleted = errors.New("ErrCategoryAlreadyDeleted") // Ошибка: категория уже удалена...
	ErrCategoryNotDeleted     = errors.New("ErrCategoryNotDeleted")     // Ошибка: категория не была удалена...

	// Ошибка нарушения уникальности записи...
	ErrUniquenessViolation = errors.New("ErrUniquenessViolation") // Ошибка: нарушение уникальности записи...

	ErrProductAlreadyExists = errors.New("ErrProductAlreadyExists") // Ошибка: продукт с таким штрих-кодом уже существует...
	ErrProductNotFound      = errors.New("ErrProductNotFound")      // Ошибка: продукт не найден...

	ErrProductAlreadyDeleted = errors.New("ErrProductAlreadyDeleted") // Ошибка: продукт уже удалён...
	ErrProductNotDeleted     = errors.New("ErrProductNotDeleted")     // Ошибка: продукт не был удалён...

	// Ошибки для сущности Заказов...
	ErrOrderNotFound     = errors.New("ErrOrderNotFound")     // Ошибка: заказ не найден...
	ErrOrderItemNotFound = errors.New("ErrOrderItemNotFound") // Ошибка: элемент заказа не найден...
	ErrInsufficientStock = errors.New("ErrInsufficientStock") // Ошибка: недостаточно товара на складе...
	ErrOrderAlreadyPaid  = errors.New("ErrOrderAlreadyPaid")  // Ошибка: заказ уже оплачен...

	ErrUnauthorized = errors.New("ErrUnauthorized") // Ошибка: неавторизованный доступ...
)
