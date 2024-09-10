// C:\GoProject\src\eShop\pkg\repository\users.go

package repository

import (
	"eShop/db"
	"eShop/logger"
	"eShop/models"
)

// GetAllUsers получает всех пользователей из базы данных...
func GetAllUsers() (users []models.User, err error) {
	// Выполняем запрос к базе данных для получения всех пользователей...
	err = db.GetDBConn().Find(&users).Error
	if err != nil {
		// Логируем ошибку, если не удалось получить список пользователей...
		logger.Error.Printf("[repository.GetAllUsers] error getting all users: %v\n", err)
		// Преобразуем и возвращаем ошибку с помощью translateError...
		return nil, translateError(err) // Пропускаем ошибку через translateError...
	}
	// Возвращаем список пользователей...
	return users, nil
}

// GetUserByID получает пользователя из базы данных по его ID...
func GetUserByID(id uint) (user models.User, err error) {
	// Выполняем запрос к базе данных для поиска пользователя по ID...
	err = db.GetDBConn().Where("id = ?", id).First(&user).Error
	if err != nil {
		// Логируем ошибку, если не удалось получить пользователя...
		logger.Error.Printf("[repository.GetUserByID] error getting user by id: %v\n", err)
		// Преобразуем и возвращаем ошибку с помощью translateError...
		return user, translateError(err)
	}
	// Возвращаем найденного пользователя...
	return user, nil
}

// UpdateUserByID обновляет данные пользователя в базе данных по его ID...
func UpdateUserByID(id uint, updatedUser models.User) (user models.User, err error) {
	// Получаем пользователя по ID...
	err = db.GetDBConn().Where("id = ?", id).First(&user).Error
	if err != nil {
		// Логируем ошибку при поиске пользователя...
		logger.Error.Printf("[repository.UpdateUserByID] error getting user by id: %v\n", err)
		return user, translateError(err)
	}

	// Обновляем данные пользователя...
	user.FullName = updatedUser.FullName
	user.Email = updatedUser.Email
	user.Username = updatedUser.Username
	user.IsBlocked = updatedUser.IsBlocked

	// Сохраняем обновленные данные в базе данных...
	err = db.GetDBConn().Save(&user).Error
	if err != nil {
		// Логируем ошибку при обновлении пользователя...
		logger.Error.Printf("[repository.UpdateUserByID] error updating user with id: %d, error: %v\n", user.ID, err)
		return user, translateError(err)
	}

	return user, nil
}

// GetUserByUsername получает пользователя из базы данных по его имени пользователя...
func GetUserByUsername(username string) (user models.User, err error) {
	// Выполняем запрос к базе данных для поиска пользователя по имени пользователя...
	err = db.GetDBConn().Where("username = ?", username).First(&user).Error
	if err != nil {
		// Логируем ошибку, если не удалось получить пользователя...
		logger.Error.Printf("[repository.GetUserByUsername] error getting user by username: %v\n", err)
		// Преобразуем и возвращаем ошибку с помощью translateError...
		return user, translateError(err)
	}

	// Возвращаем найденного пользователя...
	return user, nil
}

// GetUserByUsernameAndPassword получает пользователя из базы данных по имени пользователя и паролю...
func GetUserByUsernameAndPassword(username string, password string) (user models.User, err error) {
	// Выполняем запрос к базе данных для поиска пользователя по имени пользователя и паролю...
	err = db.GetDBConn().Where("username = ? AND password = ?", username, password).First(&user).Error
	if err != nil {
		// Логируем ошибку, если не удалось получить пользователя...
		logger.Error.Printf("[repository.GetUserByUsernameAndPassword] error getting user by username and password: %v\n", err)
		// Преобразуем и возвращаем ошибку с помощью translateError...
		return user, translateError(err)
	}

	// Возвращаем найденного пользователя...
	return user, nil
}

// CreateUser создаёт нового пользователя в базе данных...
func CreateUser(user models.User) (err error) {
	// Выполняем запрос на создание нового пользователя в базе данных...
	if err = db.GetDBConn().Create(&user).Error; err != nil {
		// Логируем ошибку, если не удалось создать пользователя...
		logger.Error.Printf("[repository.CreateUser] error creating user: %v\n", err)
		// Преобразуем и возвращаем ошибку с помощью translateError...
		return translateError(err)
	}

	// Возвращаем nil, если создание прошло успешно...
	return nil
}
