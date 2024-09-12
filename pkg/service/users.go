// C:\GoProject\src\eShop\pkg\service\users.go

package service

import (
	"eShop/errs"
	"eShop/logger"
	"eShop/models"
	"eShop/pkg/repository"
	"eShop/utils"
	"errors"
	"time"
)

// CreateUser проверяет уникальность пользователя, генерирует хеш пароля и сохраняет пользователя...
func CreateUser(user models.User) error {
	// 1. Проверяем уникальность имени пользователя...
	userFromDB, err := repository.GetUserByUsername(user.Username)
	if err != nil && !errors.Is(err, errs.ErrRecordNotFound) {
		// Если возникает ошибка, отличная от "запись не найдена", возвращаем её...
		return err
	}

	// Если пользователь с таким именем уже существует, возвращаем ошибку уникальности...
	if userFromDB.ID > 0 {
		return errs.ErrUsernameUniquenessFailed
	}

	// 2. Генерируем хеш пароля пользователя...
	user.Password = utils.GenerateHash(user.Password)

	// 3. Сохраняем пользователя через репозиторий...
	err = repository.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

// GetAllUsers получает список всех пользователей...
func GetAllUsers() (users []models.User, err error) {
	// Получаем всех пользователей из репозитория...
	users, err = repository.GetAllUsers()
	if err != nil {
		return nil, err
	}

	// Если пользователей нет, логируем предупреждение и возвращаем пустой массив...
	if len(users) == 0 {
		logger.Warning.Printf("[repository.GetAllUsers] no users found")
	}

	// Возвращаем список пользователей...
	return users, nil
}

// GetUserByID получает данные пользователя по его ID...
func GetUserByID(id uint) (user models.User, err error) {
	// Получаем пользователя из репозитория по ID...
	user, err = repository.GetUserByID(id)
	if err != nil {
		// Если пользователь не найден, возвращаем кастомную ошибку...
		if errors.Is(err, errs.ErrRecordNotFound) {
			return user, errs.ErrUserNotFound
		}
		// Возвращаем любую другую ошибку, если она возникла...
		return user, err
	}

	// Возвращаем найденного пользователя...
	return user, nil
}

// UpdateUserByID обновляет данные пользователя по его ID через репозиторий...
func UpdateUserByID(id uint, updatedUser models.User) (user models.User, err error) {
	// Получаем пользователя из репозитория по ID...
	user, err = repository.GetUserByID(id)
	if err != nil {
		// Если пользователь не найден, возвращаем кастомную ошибку...
		if errors.Is(err, errs.ErrRecordNotFound) {
			logger.Warning.Printf("[service.UpdateUserByID] no user found with id: %v", id)
			return user, errs.ErrUserNotFound
		}
		// Логируем как ошибку и возвращаем любую другую ошибку...
		logger.Error.Printf("[service.UpdateUserByID] error getting user by id: %v\n", err)
		return user, err
	}

	// Обновляем только переданные поля...
	if updatedUser.FullName != "" {
		user.FullName = updatedUser.FullName
	}
	if updatedUser.Email != "" {
		user.Email = updatedUser.Email
	}
	if updatedUser.Username != "" {
		user.Username = updatedUser.Username
	}
	if updatedUser.Password != "" {
		// Генерируем хеш пароля только если новый пароль передан...
		user.Password = utils.GenerateHash(updatedUser.Password)
	}
	if updatedUser.Role != "" {
		user.Role = updatedUser.Role
	}

	// Обновляем флаг блокировки и удаления...
	user.IsBlocked = updatedUser.IsBlocked
	user.IsDeleted = updatedUser.IsDeleted

	// Сохраняем обновлённые данные через репозиторий...
	err = repository.UpdateUserByID(user)
	if err != nil {
		// Логируем ошибку при сохранении данных...
		logger.Error.Printf("[service.UpdateUserByID] error saving updated user with id: %v\n", id)
		return user, err
	}

	// Возвращаем обновлённого пользователя...
	return user, nil
}

// SoftDeleteUserByID помечает пользователя как удалённого...
func SoftDeleteUserByID(id uint) (err error) {
	// Получаем пользователя по ID...
	user, err := repository.GetUserByID(id)
	if err != nil {
		// Проверяем, если это ошибка записи не найдена и возвращаем кастомную ошибку...
		if errors.Is(err, errs.ErrRecordNotFound) {
			logger.Warning.Printf("[service.UpdateUserByID] no user found with id: %v", id)
			return errs.ErrUserNotFound
		}
		// Логируем и возвращаем другие ошибки...
		logger.Error.Printf("[service.UpdateUserByID] error getting user by id: %v\n", err)
		return err
	}

	// Проверяем, был ли пользователь уже удалён...
	if user.IsDeleted {
		// Возвращаем ошибку, если пользователь уже удалён...
		return errs.ErrUserAlreadyDeleted
	}

	// Помечаем как удалённого...
	user.IsDeleted = true
	currentTime := time.Now()
	user.DeletedAt = &currentTime

	// Используем существующую функцию обновления...
	err = repository.UpdateUserByID(user)
	if err != nil {
		logger.Error.Printf("[service.UpdateUserByID] error saving updated user with id: %v\n", id)
		return err
	}

	return nil
}

// RestoreUserByID восстанавливает пользователя...
func RestoreUserByID(id uint) (err error) {
	// Получаем пользователя по ID...
	user, err := repository.GetUserIncludingSoftDeleted(id)
	if err != nil {
		// Если пользователь не найден, возвращаем кастомную ошибку...
		if errors.Is(err, errs.ErrRecordNotFound) {
			logger.Warning.Printf("[service.RestoreUserByID] no user found with id: %v", id)
			return errs.ErrUserNotFound
		}
		// Логируем и возвращаем другие ошибки...
		logger.Error.Printf("[service.RestoreUserByID] error getting user by id: %v\n", err)
		return err
	}

	// Проверяем, был ли пользователь не удалён...
	if !user.IsDeleted {
		// Возвращаем ошибку, если пользователь не был удалён...
		return errs.ErrUserNotDeleted
	}

	// Восстанавливаем пользователя...
	user.IsDeleted = false
	user.DeletedAt = nil

	// Используем существующую функцию обновления...
	err = repository.UpdateUserByID(user)
	if err != nil {
		logger.Error.Printf("[service.RestoreUserByID] error saving updated user with id: %v\n", id)
		return err
	}

	return nil
}

// GetDeletedUsers возвращает список всех удалённых пользователей...
func GetAllDeletedUsers() (users []models.User, err error) {
	users, err = repository.GetAllDeletedUsers()
	if err != nil {
		return nil, err
	}

	// Возвращаем список удалённых пользователей...
	return users, nil
}

// HardDeleteUserByID удаляет пользователя из базы данных...
func HardDeleteUserByID(id uint) error {
	// Вызываем репозиторий для выполнения жёсткого удаления...
	err := repository.HardDeleteUserByID(id)
	if err != nil {
		// Логируем и возвращаем ошибку, если что-то пошло не так...
		logger.Error.Printf("[service.HardDeleteUserByID] error hard deleting user with id: %v\n", id)
		return err
	}
	return nil
}

// BlockUserByID блокирует пользователя по его ID...
func BlockUserByID(id uint) (err error) {
	// Получаем пользователя из репозитория...
	user, err := repository.GetUserByID(id)
	if err != nil {
		// Проверяем, если запись не найдена...
		if errors.Is(err, errs.ErrRecordNotFound) {
			logger.Warning.Printf("[service.BlockUserByID] no user found with id: %v", id)
			return errs.ErrUserNotFound
		}
		// Логируем и возвращаем ошибку...
		logger.Error.Printf("[service.BlockUserByID] error getting user by id: %v\n", err)
		return err
	}

	// Проверяем, заблокирован ли пользователь...
	if user.IsBlocked {
		return errs.ErrUserAlreadyBlocked
	}

	// Помечаем пользователя как заблокированного и сохраняем время блокировки...
	user.IsBlocked = true
	currentTime := time.Now()
	user.BlockedAt = &currentTime

	// Сохраняем изменения в репозитории...
	err = repository.UpdateUserByID(user)
	if err != nil {
		logger.Error.Printf("[service.BlockUserByID] error updating user with id: %v\n", id)
		return err
	}

	return nil
}

// UnblockUserByID разблокирует пользователя...
func UnblockUserByID(id uint) (err error) {
	// Получаем пользователя по ID...
	user, err := repository.GetUserByID(id)
	if err != nil {
		if errors.Is(err, errs.ErrUserNotFound) {
			logger.Warning.Printf("[service.UnblockUserByID] no user found with id: %v", id)
			return errs.ErrUserNotFound
		}
		logger.Error.Printf("[service.UnblockUserByID] error getting user by id: %v\n", err)
		return err
	}

	// Проверяем, заблокирован ли пользователь...
	if !user.IsBlocked {
		return errs.ErrUserNotBlocked
	}

	// Разблокируем пользователя...
	user.IsBlocked = false
	user.BlockedAt = nil

	// Обновляем данные пользователя...
	err = repository.UpdateUserByID(user)
	if err != nil {
		logger.Error.Printf("[service.UnblockUserByID] error saving updated user with id: %v\n", id)
		return err
	}

	return nil
}

// ResetPassword сбрасывает пароль пользователя (доступно только администратору)...
func ResetPassword(userID uint, newPassword string) error {
	user, err := repository.GetUserByID(userID)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			return errs.ErrUserNotFound
		}
		return err
	}

	// Хешируем новый пароль...
	user.Password = utils.GenerateHash(newPassword)

	return repository.UpdateUserByID(user)
}

// ChangeOwnPassword позволяет пользователю изменить свой пароль...
func ChangeOwnPassword(userID uint, oldPassword, newPassword string) error {
	user, err := repository.GetUserByID(userID)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			return errs.ErrUserNotFound
		}
		return err
	}

	// Проверяем старый пароль...
	if !utils.CheckPasswordHash(oldPassword, user.Password) {
		return errs.ErrIncorrectPassword
	}

	// Хешируем новый пароль и сохраняем его...
	user.Password = utils.GenerateHash(newPassword)
	return repository.UpdateUserByID(user)
}
