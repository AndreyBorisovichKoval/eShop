// C:\GoProject\src\eShop\errs\errs.go

package errs

import "errors"

var (
	ErrUsernameUniquenessFailed    = errors.New("ErrUsernameUniquenessFailed")
	ErrOperationNotFound           = errors.New("ErrOperationNotFound")
	ErrIncorrectUsernameOrPassword = errors.New("ErrIncorrectUsernameOrPassword")
	ErrRecordNotFound              = errors.New("ErrRecordNotFound")
	ErrSomethingWentWrong          = errors.New("ErrSomethingWentWrong")
)
