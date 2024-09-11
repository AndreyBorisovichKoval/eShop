// C:\GoProject\src\eShop\pkg\repository\users.go

package repository

import (
	"eShop/db"
	"eShop/logger"
	"eShop/models"
)

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

// GetAllUsers получает всех пользователей из базы данных...
func GetAllUsers() (users []models.User, err error) {
	// Выполняем запрос к базе данных для получения всех пользователей...
	// err = db.GetDBConn().Find(&users).Error
	err = db.GetDBConn().Where("is_deleted = ?", false).Find(&users).Error
	if err != nil {
		// Логируем как ошибку для всех других типов ошибок, кроме отсутствия записей...
		logger.Error.Printf("[repository.GetAllUsers] error getting all users: %v\n", err)
		return nil, translateError(err)
	}

	// Возвращаем список пользователей (или пустой массив, если их нет)...
	return users, nil
}

// GetUserByID получает пользователя из базы данных по его ID...
func GetUserByID(id uint) (user models.User, err error) {
	// Выполняем запрос к базе данных для поиска пользователя по ID...
	err = db.GetDBConn().Where("id = ? AND is_deleted = ?", id, false).First(&user).Error
	if err != nil {
		// Логируем ошибку, если не удалось получить пользователя...
		logger.Error.Printf("[repository.GetUserByID] error getting user by id: %v\n", err)
		// Преобразуем и возвращаем ошибку с помощью translateError...
		return user, translateError(err)
	}
	// Возвращаем найденного пользователя...
	return user, nil
}

// GetUserIncludingSoftDeleted получает пользователя из базы данных по ID, включая удалённых...
func GetUserIncludingSoftDeleted(id uint) (user models.User, err error) {
	// Выполняем запрос к базе данных для поиска пользователя по ID...
	err = db.GetDBConn().Where("id = ?", id).First(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.GetUserIncludingDeleted] error getting user by id: %v\n", err)
		return user, translateError(err)
	}
	return user, nil
}

// UpdateUserByID обновляет данные пользователя в базе данных...
func UpdateUserByID(user models.User) (err error) {
	// Сохраняем изменения пользователя в базе данных...
	err = db.GetDBConn().Save(&user).Error
	if err != nil {
		// Логируем ошибку при сохранении данных пользователя...
		logger.Error.Printf("[repository.UpdateUserByID] error updating user with id: %v, error: %v\n", user.ID, err)
		return translateError(err)
	}

	// Возвращаем nil при успешном обновлении...
	return nil
}

// GetDeletedUsers получает всех удалённых пользователей из базы данных...
func GetAllDeletedUsers() (users []models.User, err error) {
	err = db.GetDBConn().Where("is_deleted = ?", true).Find(&users).Error
	if err != nil {
		logger.Error.Printf("[repository.GetDeletedUsers] error getting deleted users: %v\n", err)
		return nil, translateError(err)
	}

	return users, nil
}

// // HardDeleteUserByID полностью удаляет пользователя из базы данных...
// func HardDeleteUserByID(user models.User) (err error) {
// 	// Выполняем реальное удаление записи из базы данных...
// 	err = db.GetDBConn().Unscoped().Delete(&user).Error
// 	if err != nil {
// 		logger.Error.Printf("[repository.HardDeleteUserByID] error hard deleting user with id: %v, error: %v\n", user.ID, err)
// 		return translateError(err)
// 	}

// 	return nil
// }

// HardDeleteUserByID удаляет пользователя из базы данных...
func HardDeleteUserByID(id uint) error {
	// Находим пользователя с учётом удалённых записей...
	var user models.User
	err := db.GetDBConn().Unscoped().Where("id = ?", id).First(&user).Error
	if err != nil {
		// Логируем ошибку, если не удалось найти пользователя...
		logger.Error.Printf("[repository.HardDeleteUserByID] error finding user by id: %v\n", err)
		return translateError(err)
	}

	// Выполняем жёсткое удаление пользователя...
	err = db.GetDBConn().Unscoped().Delete(&user).Error
	if err != nil {
		// Логируем ошибку при удалении пользователя...
		logger.Error.Printf("[repository.HardDeleteUserByID] error deleting user by id: %v\n", err)
		return translateError(err)
	}

	return nil
}
