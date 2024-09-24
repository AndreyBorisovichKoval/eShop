// C:\GoProject\src\eShop\errs\errs.go

package errs

import "errors"

// Определение кастомных ошибок, используемых в приложении...

// Ошибки аутентификации
var (
	ErrEmptyAuthHeader       = errors.New("empty auth header")       // Ошибка: пустой заголовок аутентификации...
	ErrInvalidAuthHeader     = errors.New("invalid auth header")     // Ошибка: некорректный заголовок аутентификации...
	ErrTokenParsingFailed    = errors.New("token parsing failed")    // Ошибка: не удалось обработать токен...
	ErrUserNotAuthenticated  = errors.New("user not authenticated")  // Ошибка: пользователь не аутентифицирован...
	ErrPasswordResetRequired = errors.New("password reset required") // Ошибка: необходимо сменить пароль...
)

// Ошибки разрешений
var (
	ErrPermissionDenied                      = errors.New("access denied")                          // Ошибка: доступ запрещен...
	ErrPermissionDeniedOnlyForAdmin          = errors.New("access denied: admins only")             // Ошибка: доступ запрещен только для администраторов...
	ErrPermissionDeniedOnlyForAdminOrManager = errors.New("access denied: admins or managers only") // Ошибка: доступ запрещен только для администраторов и менеджеров...
)

// Ошибки валидации
var (
	ErrValidationFailed            = errors.New("validation failed")              // Ошибка: валидация данных не пройдена...
	ErrUsernameUniquenessFailed    = errors.New("username must be unique")        // Ошибка: имя пользователя должно быть уникальным...
	ErrIncorrectUsernameOrPassword = errors.New("incorrect username or password") // Ошибка: неверное имя пользователя или пароль...
)

// Ошибки пользователей
var (
	ErrUserBlocked                = errors.New("user is blocked")              // Ошибка: пользователь заблокирован...
	ErrRecordNotFound             = errors.New("record not found")             // Ошибка: запись не найдена...
	ErrUserNotFound               = errors.New("user not found")               // Ошибка: пользователь не найден...
	ErrUsersNotFound              = errors.New("users not found")              // Ошибка: пользователи не найдены...
	ErrUserAlreadyDeleted         = errors.New("user already deleted")         // Ошибка: пользователь уже был удалён ранее...
	ErrUserNotDeleted             = errors.New("user not deleted")             // Ошибка: пользователь не был удалён...
	ErrSomethingWentWrong         = errors.New("something went wrong")         // Ошибка: что-то пошло не так...
	ErrUserAlreadyBlocked         = errors.New("user already blocked")         // Ошибка: пользователь уже заблокирован...
	ErrUserNotBlocked             = errors.New("user not blocked")             // Ошибка: пользователь не был заблокирован...
	ErrUnauthorizedPasswordChange = errors.New("unauthorized password change") // Ошибка: попытка смены пароля без прав...
	ErrIncorrectPassword          = errors.New("incorrect old password")       // Ошибка: неверный старый пароль...
)

// Ошибки для сущности Поставщика
var (
	ErrSupplierAlreadyExists  = errors.New("supplier already exists")  // Ошибка: поставщик с таким именем или email уже существует...
	ErrSupplierNotFound       = errors.New("supplier not found")       // Ошибка: поставщик не найден...
	ErrSupplierAlreadyDeleted = errors.New("supplier already deleted") // Ошибка: поставщик уже удалён...
	ErrSupplierNotDeleted     = errors.New("supplier not deleted")     // Ошибка: поставщик не был удалён...
)

// Ошибки для сущности Категории
var (
	ErrCategoryAlreadyExists  = errors.New("category already exists")  // Ошибка: категория с таким названием уже существует...
	ErrCategoryNotFound       = errors.New("category not found")       // Ошибка: категория не найдена...
	ErrCategoryAlreadyDeleted = errors.New("category already deleted") // Ошибка: категория уже удалена...
	ErrCategoryNotDeleted     = errors.New("category not deleted")     // Ошибка: категория не была удалена...
)

// Ошибка нарушения уникальности записи
var (
	ErrUniquenessViolation = errors.New("uniqueness violation") // Ошибка: нарушение уникальности записи...
)

// Ошибки для сущности Продукта
var (
	ErrProductAlreadyExists  = errors.New("product already exists")  // Ошибка: продукт с таким штрих-кодом уже существует...
	ErrProductNotFound       = errors.New("product not found")       // Ошибка: продукт не найден...
	ErrProductAlreadyDeleted = errors.New("product already deleted") // Ошибка: продукт уже удалён...
	ErrProductNotDeleted     = errors.New("product not deleted")     // Ошибка: продукт не был удалён...
)

// Ошибки для сущности Заказов
var (
	ErrOrderNotFound     = errors.New("order not found")      // Ошибка: заказ не найден...
	ErrOrderItemNotFound = errors.New("order item not found") // Ошибка: элемент заказа не найден...
	ErrInsufficientStock = errors.New("insufficient stock")   // Ошибка: недостаточно товара на складе...
	ErrOrderAlreadyPaid  = errors.New("order already paid")   // Ошибка: заказ уже оплачен...
)

// Ошибки при работе с оплачиваемыми заказами
var (
	ErrProductNotWeightBased     = errors.New("product not weight based")              // Ошибка: продукт не основан на весе...
	ErrCannotDeletePaidOrder     = errors.New("cannot delete a paid order")            // Ошибка: нельзя удалить оплаченный заказ...
	ErrCannotDeletePaidOrderItem = errors.New("cannot delete items from a paid order") // Ошибка: нельзя удалить товар из оплаченного заказа...
	ErrCannotAddToPaidOrder      = errors.New("cannot add products to a paid order")   // Ошибка: нельзя добавлять в оплаченный заказ...
)

// Общие ошибки
var (
	ErrUnauthorized = errors.New("unauthorized access") // Ошибка: неавторизованный доступ...
	ErrServerError  = errors.New("server error")        // Ошибка: ошибка сервера...
)
