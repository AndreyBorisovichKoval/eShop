// C:\GoProject\src\eShop\pkg\repository\users.go

package repository

import (
	"eShop/db"
	"eShop/errs"
	"eShop/logger"
	"eShop/models"
	"errors"

	"gorm.io/gorm"
)

// GetUserByUsername получает пользователя из базы данных по его имени пользователя...
func GetUserByUsername(username string) (user models.User, err error) {
	err = db.GetDBConn().Where("username = ?", username).First(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.GetUserByUsername] error getting user by username: %v\n", err)
		return user, translateError(err)
	}
	logger.Info.Printf("[repository.GetUserByUsername] user retrieved successfully by username: %s", username) // Лог успешного получения пользователя
	return user, nil
}

// GetUserByUsernameAndPassword получает пользователя по имени пользователя и паролю...
func GetUserByUsernameAndPassword(username string, password string) (user models.User, err error) {
	err = db.GetDBConn().Where("username = ? AND password = ?", username, password).First(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.GetUserByUsernameAndPassword] error getting user by username and password: %v\n", err)
		return user, translateError(err)
	}
	logger.Info.Printf("[repository.GetUserByUsernameAndPassword] user retrieved successfully by username and password") // Лог успешного получения пользователя
	return user, nil
}

// CreateUser создаёт нового пользователя в базе данных...
func CreateUser(user *models.User) (err error) {
	if err = db.GetDBConn().Create(user).Error; err != nil {
		logger.Error.Printf("[repository.CreateUser] error creating user: %v\n", err)
		return translateError(err)
	}
	logger.Info.Printf("[repository.CreateUser] user created successfully with ID: %d", user.ID) // Лог успешного создания пользователя
	return nil
}

// GetAllUsers получает всех пользователей из базы данных...
func GetAllUsers() (users []models.User, err error) {
	err = db.GetDBConn().Where("is_deleted = ?", false).Find(&users).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllUsers] error getting all users: %v\n", err)
		return nil, translateError(err)
	}
	logger.Info.Printf("[repository.GetAllUsers] users retrieved successfully") // Лог успешного получения всех пользователей
	return users, nil
}

// GetAllDeletedUsers получает всех удалённых пользователей из базы данных...
func GetAllDeletedUsers() (users []models.User, err error) {
	err = db.GetDBConn().Where("is_deleted = ?", true).Find(&users).Error
	if err != nil {
		logger.Error.Printf("[repository.GetDeletedUsers] error getting deleted users: %v\n", err)
		return nil, translateError(err)
	}
	logger.Info.Printf("[repository.GetDeletedUsers] deleted users retrieved successfully") // Лог успешного получения удалённых пользователей
	return users, nil
}

// GetUserByID получает пользователя из базы данных по его ID...
func GetUserByID(id uint) (user models.User, err error) {
	err = db.GetDBConn().Where("id = ? AND is_deleted = ?", id, false).First(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.GetUserByID] error getting user by id: %v\n", err)
		return user, translateError(err)
	}
	logger.Info.Printf("[repository.GetUserByID] user retrieved successfully by ID: %d", id) // Лог успешного получения пользователя
	return user, nil
}

// GetUserIncludingSoftDeleted получает пользователя по ID, включая удалённых...
func GetUserIncludingSoftDeleted(id uint) (user models.User, err error) {
	err = db.GetDBConn().Where("id = ?", id).First(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.GetUserIncludingDeleted] error getting user by id: %v\n", err)
		return user, translateError(err)
	}
	logger.Info.Printf("[repository.GetUserIncludingDeleted] user retrieved successfully including soft deleted by ID: %d", id) // Лог успешного получения пользователя, включая удалённых
	return user, nil
}

// UpdateUserByID обновляет данные пользователя в базе данных...
func UpdateUserByID(user models.User) (err error) {
	err = db.GetDBConn().Save(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.UpdateUserByID] error updating user with id: %v, error: %v\n", user.ID, err)
		return translateError(err)
	}
	logger.Info.Printf("[repository.UpdateUserByID] user updated successfully with ID: %d", user.ID) // Лог успешного обновления пользователя
	return nil
}

// HardDeleteUserByID удаляет пользователя из базы данных...
func HardDeleteUserByID(id uint) error {
	var user models.User
	err := db.GetDBConn().Unscoped().Where("id = ?", id).First(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.HardDeleteUserByID] error finding user by id: %v\n", err)
		return translateError(err)
	}

	err = db.GetDBConn().Unscoped().Delete(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.HardDeleteUserByID] error deleting user by id: %v\n", err)
		return translateError(err)
	}
	logger.Info.Printf("[repository.HardDeleteUserByID] user hard deleted successfully with ID: %d", id) // Лог успешного удаления пользователя
	return nil
}

// CreateUserSettings создаёт запись с настройками пользователя в базе данных
func CreateUserSettings(userSettings models.UserSettings) error {
	if err := db.GetDBConn().Create(&userSettings).Error; err != nil {
		logger.Error.Printf("[repository.CreateUserSettings] error creating user settings: %v\n", err)
		return translateError(err)
	}
	logger.Info.Printf("[repository.CreateUserSettings] user settings created successfully for user ID: %d", userSettings.UserID) // Лог успешного создания настроек
	return nil
}

// GetUserSettingsByUserID получает настройки пользователя по его ID
func GetUserSettingsByUserID(userID uint) (models.UserSettings, error) {
	var settings models.UserSettings
	if err := db.GetDBConn().Where("user_id = ?", userID).First(&settings).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Warning.Printf("[repository.GetUserSettingsByUserID] User settings not found for user ID: %d", userID)
			return settings, errs.ErrRecordNotFound
		}
		logger.Error.Printf("[repository.GetUserSettingsByUserID] Error fetching settings for user ID: %d: %v", userID, err)
		return settings, translateError(err)
	}
	logger.Info.Printf("[repository.GetUserSettingsByUserID] user settings retrieved successfully for user ID: %d", userID) // Лог успешного получения настроек пользователя
	return settings, nil
}

// UpdateUserSettings обновляет настройки пользователя в базе данных
func UpdateUserSettings(settings models.UserSettings) error {
	if err := db.GetDBConn().Save(&settings).Error; err != nil {
		logger.Error.Printf("[repository.UpdateUserSettings] Error updating settings for user ID: %d: %v", settings.UserID, err)
		return translateError(err)
	}
	logger.Info.Printf("[repository.UpdateUserSettings] user settings updated successfully for user ID: %d", settings.UserID) // Лог успешного обновления настроек
	return nil
}

// package repository

// import (
// 	"eShop/db"
// 	"eShop/errs"
// 	"eShop/logger"
// 	"eShop/models"
// 	"errors"

// 	"gorm.io/gorm"
// )

// // GetUserByUsername получает пользователя из базы данных по его имени пользователя...
// func GetUserByUsername(username string) (user models.User, err error) {
// 	// Выполняем запрос к базе данных для поиска пользователя по имени пользователя...
// 	err = db.GetDBConn().Where("username = ?", username).First(&user).Error
// 	if err != nil {
// 		// Логируем ошибку, если не удалось получить пользователя...
// 		logger.Error.Printf("[repository.GetUserByUsername] error getting user by username: %v\n", err)
// 		// Преобразуем и возвращаем ошибку с помощью translateError...
// 		return user, translateError(err)
// 	}

// 	// Возвращаем найденного пользователя...
// 	return user, nil
// }

// // GetUserByUsernameAndPassword получает пользователя из базы данных по имени пользователя и паролю...
// func GetUserByUsernameAndPassword(username string, password string) (user models.User, err error) {
// 	// Выполняем запрос к базе данных для поиска пользователя по имени пользователя и паролю...
// 	err = db.GetDBConn().Where("username = ? AND password = ?", username, password).First(&user).Error
// 	if err != nil {
// 		// Логируем ошибку, если не удалось получить пользователя...
// 		logger.Error.Printf("[repository.GetUserByUsernameAndPassword] error getting user by username and password: %v\n", err)
// 		// Преобразуем и возвращаем ошибку с помощью translateError...
// 		return user, translateError(err)
// 	}

// 	// Возвращаем найденного пользователя...
// 	return user, nil
// }

// // // CreateUser создаёт нового пользователя в базе данных...
// // func CreateUser(user models.User) (err error) {
// // 	// Выполняем запрос на создание нового пользователя в базе данных...
// // 	if err = db.GetDBConn().Create(&user).Error; err != nil {
// // 		// Логируем ошибку, если не удалось создать пользователя...
// // 		logger.Error.Printf("[repository.CreateUser] error creating user: %v\n", err)
// // 		// Преобразуем и возвращаем ошибку с помощью translateError...
// // 		return translateError(err)
// // 	}

// // 	// Возвращаем nil, если создание прошло успешно...
// // 	return nil
// // }

// // CreateUser создаёт нового пользователя в базе данных...
// func CreateUser(user *models.User) (err error) {
// 	// Выполняем запрос на создание нового пользователя в базе данных...
// 	if err = db.GetDBConn().Create(user).Error; err != nil {
// 		logger.Error.Printf("[repository.CreateUser] error creating user: %v\n", err)
// 		return translateError(err)
// 	}
// 	// Теперь ID пользователя доступен после создания.
// 	return nil
// }

// // GetAllUsers получает всех пользователей из базы данных...
// func GetAllUsers() (users []models.User, err error) {
// 	// Выполняем запрос к базе данных для получения всех пользователей...
// 	// err = db.GetDBConn().Find(&users).Error
// 	err = db.GetDBConn().Where("is_deleted = ?", false).Find(&users).Error
// 	if err != nil {
// 		// Логируем как ошибку для всех других типов ошибок, кроме отсутствия записей...
// 		logger.Error.Printf("[repository.GetAllUsers] error getting all users: %v\n", err)
// 		return nil, translateError(err)
// 	}

// 	// Возвращаем список пользователей (или пустой массив, если их нет)...
// 	return users, nil
// }

// // GetDeletedUsers получает всех удалённых пользователей из базы данных...
// func GetAllDeletedUsers() (users []models.User, err error) {
// 	err = db.GetDBConn().Where("is_deleted = ?", true).Find(&users).Error
// 	if err != nil {
// 		logger.Error.Printf("[repository.GetDeletedUsers] error getting deleted users: %v\n", err)
// 		return nil, translateError(err)
// 	}

// 	return users, nil
// }

// // GetUserByID получает пользователя из базы данных по его ID...
// func GetUserByID(id uint) (user models.User, err error) {
// 	// Выполняем запрос к базе данных для поиска пользователя по ID...
// 	err = db.GetDBConn().Where("id = ? AND is_deleted = ?", id, false).First(&user).Error
// 	if err != nil {
// 		// Логируем ошибку, если не удалось получить пользователя...
// 		logger.Error.Printf("[repository.GetUserByID] error getting user by id: %v\n", err)
// 		// Преобразуем и возвращаем ошибку с помощью translateError...
// 		return user, translateError(err)
// 	}
// 	// Возвращаем найденного пользователя...
// 	return user, nil
// }

// // GetUserIncludingSoftDeleted получает пользователя из базы данных по ID, включая удалённых...
// func GetUserIncludingSoftDeleted(id uint) (user models.User, err error) {
// 	// Выполняем запрос к базе данных для поиска пользователя по ID...
// 	err = db.GetDBConn().Where("id = ?", id).First(&user).Error
// 	if err != nil {
// 		logger.Error.Printf("[repository.GetUserIncludingDeleted] error getting user by id: %v\n", err)
// 		return user, translateError(err)
// 	}
// 	return user, nil
// }

// // UpdateUserByID обновляет данные пользователя в базе данных...
// func UpdateUserByID(user models.User) (err error) {
// 	// Сохраняем изменения пользователя в базе данных...
// 	err = db.GetDBConn().Save(&user).Error
// 	if err != nil {
// 		// Логируем ошибку при сохранении данных пользователя...
// 		logger.Error.Printf("[repository.UpdateUserByID] error updating user with id: %v, error: %v\n", user.ID, err)
// 		return translateError(err)
// 	}

// 	// Возвращаем nil при успешном обновлении...
// 	return nil
// }

// // HardDeleteUserByID удаляет пользователя из базы данных...
// func HardDeleteUserByID(id uint) error {
// 	// Находим пользователя с учётом удалённых записей...
// 	var user models.User
// 	err := db.GetDBConn().Unscoped().Where("id = ?", id).First(&user).Error
// 	if err != nil {
// 		// Логируем ошибку, если не удалось найти пользователя...
// 		logger.Error.Printf("[repository.HardDeleteUserByID] error finding user by id: %v\n", err)
// 		return translateError(err)
// 	}

// 	// Выполняем жёсткое удаление пользователя...
// 	err = db.GetDBConn().Unscoped().Delete(&user).Error
// 	if err != nil {
// 		// Логируем ошибку при удалении пользователя...
// 		logger.Error.Printf("[repository.HardDeleteUserByID] error deleting user by id: %v\n", err)
// 		return translateError(err)
// 	}

// 	return nil
// }

// // CreateUserSettings создаёт запись с настройками пользователя в базе данных
// func CreateUserSettings(userSettings models.UserSettings) error {
// 	if err := db.GetDBConn().Create(&userSettings).Error; err != nil {
// 		logger.Error.Printf("[repository.CreateUserSettings] error creating user settings: %v\n", err)
// 		return translateError(err)
// 	}
// 	return nil
// }

// // GetUserSettingsByUserID получает настройки пользователя по его ID
// func GetUserSettingsByUserID(userID uint) (models.UserSettings, error) {
// 	var settings models.UserSettings
// 	if err := db.GetDBConn().Where("user_id = ?", userID).First(&settings).Error; err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			logger.Warning.Printf("[repository.GetUserSettingsByUserID] User settings not found for user ID: %d", userID)
// 			return settings, errs.ErrRecordNotFound
// 		}
// 		logger.Error.Printf("[repository.GetUserSettingsByUserID] Error fetching settings for user ID: %d: %v", userID, err)
// 		return settings, translateError(err)
// 	}
// 	return settings, nil
// }

// // UpdateUserSettings обновляет настройки пользователя в базе данных
// func UpdateUserSettings(settings models.UserSettings) error {
// 	if err := db.GetDBConn().Save(&settings).Error; err != nil {
// 		logger.Error.Printf("[repository.UpdateUserSettings] Error updating settings for user ID: %d: %v", settings.UserID, err)
// 		return translateError(err)
// 	}
// 	return nil
// }
