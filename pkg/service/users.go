// C:\GoProject\src\eShop\pkg\service\users.go

package service

import (
	"eShop/errs"
	"eShop/models"
	"eShop/pkg/repository"
	"eShop/utils"
	"errors"
)

// GetAllUsers получает список всех пользователей...
func GetAllUsers() (users []models.User, err error) {
	// Получаем всех пользователей из репозитория...
	users, err = repository.GetAllUsers()
	if err != nil {
		// Если не найдено ни одного пользователя, возвращаем кастомную ошибку...
		if errors.Is(err, errs.ErrRecordNotFound) {
			return nil, errs.ErrUsersNotFound
		}
		// Возвращаем любую другую ошибку, если она возникла...
		return nil, err
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

// UpdateUserByID обновляет данные пользователя по его ID, проверяя его существование...
func UpdateUserByID(id uint, updatedUser models.User) (user models.User, err error) {
	// Получаем пользователя из репозитория по ID...
	user, err = repository.UpdateUserByID(id, updatedUser)
	if err != nil {
		// Если пользователь не найден, возвращаем ошибку, что пользователь не существует...
		if errors.Is(err, errs.ErrRecordNotFound) {
			return user, errs.ErrUserNotFound
		}
		// Возвращаем любую другую ошибку, если она возникла...
		return user, err
	}
	// Возвращаем обновлённые данные пользователя...
	return user, nil
}

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
