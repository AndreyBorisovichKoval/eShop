// C:\GoProject\src\eShop\pkg\controllers\users.go

package controllers

import (
	"eShop/errs"
	"eShop/logger"
	"eShop/models"
	"eShop/pkg/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAllUsers
// @Summary Retrieve all users
// @Tags users
// @Description Get a list of all registered users
// @ID get-all-users
// @Produce json
// @Success 200 {array} models.User "List of users"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /users [get]
// @Security ApiKeyAuth
func GetAllUsers(c *gin.Context) {
	// Логируем IP клиента при запросе списка всех пользователей...
	logger.Info.Printf("IP: [%s] requested list of all users\n", c.ClientIP())

	// Вызываем сервис для получения списка всех пользователей...
	users, err := service.GetAllUsers()
	if err != nil {
		// Логируем ошибку при получении списка пользователей...
		logger.Error.Printf("[controllers.GetAllUsers] error getting all users: %v\n", err)
		handleError(c, err)
		return
	}

	// Логируем IP клиента при успешной выдаче списка пользователей...
	logger.Info.Printf("IP: [%s] got list of all users\n", c.ClientIP())

	// Отправляем клиенту ответ со списком пользователей...
	c.JSON(http.StatusOK, users)
}

// GetUserByID
// @Summary Retrieve user by ID
// @Tags users
// @Description Get user information by user ID
// @ID get-user-by-id
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.User "User information"
// @Failure 400 {object} ErrorResponse "Invalid ID"
// @Failure 404 {object} ErrorResponse "User not found"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /users/{id} [get]
// @Security ApiKeyAuth
func GetUserByID(c *gin.Context) {
	// Извлекаем ID пользователя из параметра запроса...
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// Логируем ошибку при некорректном параметре ID...
		logger.Error.Printf("[controllers.GetUserByID] invalid user_id path parameter: %s, IP: [%s]\n", c.Param("id"), c.ClientIP())
		handleError(c, errs.ErrValidationFailed)
		return
	}

	// Логируем IP клиента и запрашиваемый ID пользователя...
	logger.Info.Printf("IP: [%s] requested user with ID: %d\n", c.ClientIP(), id)

	// Вызываем сервис для получения данных пользователя по ID...
	user, err := service.GetUserByID(uint(id))
	if err != nil {
		handleError(c, err)
		return
	}

	// Логируем успешное получение данных пользователя...
	logger.Info.Printf("IP: [%s] got user with ID: %d\n", c.ClientIP(), id)

	// Отправляем клиенту ответ с данными пользователя...
	c.JSON(http.StatusOK, user)
}

// UpdateUserByID
// @Summary Update user by ID
// @Tags users
// @Description Update user information by user ID (Admin only)
// @ID update-user-by-id
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param input body models.User true "Updated user information"
// @Success 200 {object} models.User "Updated user"
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Failure 403 {object} ErrorResponse "Permission denied"
// @Failure 404 {object} ErrorResponse "User not found"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /users/{id} [patch]
// @Security ApiKeyAuth
func UpdateUserByID(c *gin.Context) {
	// Получаем роль пользователя из контекста. Если она не задана, возвращаем ошибку валидации...
	userRole := c.GetString(userRoleCtx)
	// if userRole == "" {
	// 	handleError(c, errs.ErrValidationFailed)
	// 	return
	// }

	// Проверяем, является ли пользователь администратором. Если нет, возвращаем ошибку "Доступ запрещен"...
	if userRole != "Admin" {
		handleError(c, errs.ErrPermissionDenied)
		return
	}

	// Извлекаем ID пользователя из параметра запроса и конвертируем его в число...
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// Логируем ошибку при некорректном параметре ID и возвращаем ошибку валидации...
		logger.Error.Printf("[controllers.UpdateUserByID] invalid user_id path parameter: %s, IP: [%s]\n", c.Param("id"), c.ClientIP())
		handleError(c, errs.ErrValidationFailed)
		return
	}

	// Логируем IP клиента и ID пользователя, для которого запрашивается обновление...
	logger.Info.Printf("IP: [%s] requested to update user with ID: %d\n", c.ClientIP(), id)

	// Извлекаем данные для обновления пользователя из тела запроса. Если данные некорректны, возвращаем ошибку валидации...
	var updatedUser models.User
	if err := c.BindJSON(&updatedUser); err != nil {
		// Логируем ошибку, если данные некорректны...
		logger.Error.Printf("[controllers.UpdateUserByID] invalid user data, IP: [%s]\n", c.ClientIP())
		handleError(c, errs.ErrValidationFailed)
		return
	}

	// Вызываем сервис для обновления данных пользователя. В случае ошибки передаем её клиенту...
	user, err := service.UpdateUserByID(uint(id), updatedUser)
	if err != nil {
		handleError(c, err)
		return
	}

	// Логируем успешное обновление пользователя...
	logger.Info.Printf("IP: [%s] successfully updated user with ID: %d\n", c.ClientIP(), id)

	// Отправляем клиенту ответ с обновлёнными данными пользователя...
	c.JSON(http.StatusOK, user)
}

// SoftDeleteUserByID
// @Summary Soft delete user by ID
// @Tags users
// @Description Soft delete user by ID (Admin only)
// @ID soft-delete-user-by-id
// @Param id path int true "User ID"
// @Success 200 {string} string "User soft deleted successfully!"
// @Failure 400 {object} ErrorResponse "Invalid ID"
// @Failure 403 {object} ErrorResponse "Permission denied"
// @Failure 404 {object} ErrorResponse "User not found"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /users/{id}/soft [delete]
// @Security ApiKeyAuth
func SoftDeleteUserByID(c *gin.Context) {
	// Получаем роль пользователя из контекста. Если она не задана, возвращаем ошибку валидации...
	userRole := c.GetString(userRoleCtx)
	// if userRole == "" {
	// 	handleError(c, errs.ErrValidationFailed)
	// 	return
	// }

	// Проверяем, является ли пользователь администратором. Если нет, возвращаем ошибку "Доступ запрещен"...
	if userRole != "Admin" {
		handleError(c, errs.ErrPermissionDenied)
		return
	}

	// Извлекаем ID пользователя...
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	// Логируем запрос...
	logger.Info.Printf("IP: [%s] requested to soft delete user with ID: %d\n", c.ClientIP(), id)

	// Вызываем сервис для софт удаления пользователя...
	err = service.SoftDeleteUserByID(uint(id))
	if err != nil {
		handleError(c, err)
		return
	}

	// Логируем успешное удаление...
	logger.Info.Printf("IP: [%s] successfully soft deleted user with ID: %d\n", c.ClientIP(), id)

	// Отправляем клиенту ответ...
	c.JSON(http.StatusOK, gin.H{"message": "User soft deleted successfully!"})
}

// RestoreUserByID
// @Summary Restore user by ID
// @Tags users
// @Description Restore a soft deleted user by ID (Admin only)
// @ID restore-user-by-id
// @Param id path int true "User ID"
// @Success 200 {string} string "User restored successfully!"
// @Failure 400 {object} ErrorResponse "Invalid ID"
// @Failure 403 {object} ErrorResponse "Permission denied"
// @Failure 404 {object} ErrorResponse "User not found"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /users/{id}/restore [patch]
// @Security ApiKeyAuth
func RestoreUserByID(c *gin.Context) {
	// Получаем роль пользователя из контекста. Если она не задана, возвращаем ошибку валидации...
	userRole := c.GetString(userRoleCtx)
	// if userRole == "" {
	// 	handleError(c, errs.ErrValidationFailed)
	// 	return
	// }

	// Проверяем, является ли пользователь администратором. Если нет, возвращаем ошибку "Доступ запрещен"...
	if userRole != "Admin" {
		handleError(c, errs.ErrPermissionDenied)
		return
	}

	// Извлекаем ID пользователя...
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	// Логируем запрос...
	logger.Info.Printf("IP: [%s] requested to restore user with ID: %d\n", c.ClientIP(), id)

	// Вызываем сервис для восстановления пользователя...
	err = service.RestoreUserByID(uint(id))
	if err != nil {
		handleError(c, err)
		return
	}

	// Логируем успешное восстановление...
	logger.Info.Printf("IP: [%s] successfully restored user with ID: %d\n", c.ClientIP(), id)

	// Отправляем клиенту ответ...
	c.JSON(http.StatusOK, gin.H{"message": "User restored successfully!"})
}

// GetAllDeletedUsers
// @Summary Retrieve all deleted users
// @Tags users
// @Description Get a list of all soft deleted users (Admin only)
// @ID get-all-deleted-users
// @Produce json
// @Success 200 {array} models.User "List of deleted users"
// @Failure 403 {object} ErrorResponse "Permission denied"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /users/deleted [get]
// @Security ApiKeyAuth
func GetAllDeletedUsers(c *gin.Context) {
	// Получаем роль пользователя из контекста...
	userRole := c.GetString(userRoleCtx)
	// if userRole == "" {
	// 	handleError(c, errs.ErrValidationFailed)
	// 	return
	// }

	// Проверяем, является ли пользователь администратором...
	if userRole != "Admin" {
		handleError(c, errs.ErrPermissionDenied)
		return
	}

	// Логируем запрос...
	logger.Info.Printf("IP: [%s] requested list of all deleted users\n", c.ClientIP())

	// Вызываем сервис для получения списка удалённых пользователей...
	users, err := service.GetAllDeletedUsers()
	if err != nil {
		handleError(c, err)
		return
	}

	// Логируем успешное получение списка...
	logger.Info.Printf("IP: [%s] successfully retrieved list of deleted users\n", c.ClientIP())

	// Возвращаем список удалённых пользователей клиенту...
	c.JSON(http.StatusOK, users)
}

// HardDeleteUserByID
// @Summary Hard delete user by ID
// @Tags users
// @Description Permanently delete user by ID (Admin only)
// @ID hard-delete-user-by-id
// @Param id path int true "User ID"
// @Success 200 {string} string "User hard deleted successfully!"
// @Failure 400 {object} ErrorResponse "Invalid ID"
// @Failure 403 {object} ErrorResponse "Permission denied"
// @Failure 404 {object} ErrorResponse "User not found"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /users/{id}/hard [delete]
// @Security ApiKeyAuth
func HardDeleteUserByID(c *gin.Context) {
	// Получаем роль пользователя из контекста...
	userRole := c.GetString(userRoleCtx)
	// if userRole == "" {
	// 	handleError(c, errs.ErrValidationFailed)
	// 	return
	// }

	// Проверяем, является ли пользователь администратором...
	if userRole != "Admin" {
		handleError(c, errs.ErrPermissionDenied)
		return
	}

	// Извлекаем ID пользователя...
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	// Логируем запрос...
	logger.Info.Printf("IP: [%s] requested to hard delete user with ID: %d\n", c.ClientIP(), id)

	// Вызываем сервис для реального удаления пользователя...
	err = service.HardDeleteUserByID(uint(id))
	if err != nil {
		handleError(c, err)
		return
	}

	// Логируем успешное удаление...
	logger.Info.Printf("IP: [%s] successfully hard deleted user with ID: %d\n", c.ClientIP(), id)

	// Отправляем клиенту ответ...
	c.JSON(http.StatusOK, gin.H{"message": "User hard deleted successfully!"})
}

// BlockUserByID
// @Summary Block user by ID
// @Tags users
// @Description Block a user by ID (Admin only)
// @ID block-user-by-id
// @Param id path int true "User ID"
// @Success 200 {string} string "User blocked successfully!"
// @Failure 400 {object} ErrorResponse "Invalid ID"
// @Failure 403 {object} ErrorResponse "Permission denied"
// @Failure 404 {object} ErrorResponse "User not found"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /users/{id}/block [patch]
// @Security ApiKeyAuth
func BlockUserByID(c *gin.Context) {
	// Получаем роль пользователя из контекста...
	userRole := c.GetString(userRoleCtx)
	// if userRole == "" {
	// 	handleError(c, errs.ErrValidationFailed)
	// 	return
	// }

	// Проверяем, является ли пользователь администратором...
	if userRole != "Admin" {
		handleError(c, errs.ErrPermissionDenied)
		return
	}

	// Извлекаем ID пользователя из параметра запроса...
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	// Логируем попытку блокировки...
	logger.Info.Printf("IP: [%s] requested to block user with ID: %d\n", c.ClientIP(), id)

	// Вызываем сервис для блокировки пользователя...
	err = service.BlockUserByID(uint(id))
	if err != nil {
		handleError(c, err)
		return
	}

	// Логируем успешную блокировку...
	logger.Info.Printf("IP: [%s] successfully blocked user with ID: %d\n", c.ClientIP(), id)

	// Отправляем клиенту ответ...
	c.JSON(http.StatusOK, gin.H{"message": "User blocked successfully!"})
}

// UnblockUserByID
// @Summary Unblock user by ID
// @Tags users
// @Description Unblock a user by ID (Admin only)
// @ID unblock-user-by-id
// @Param id path int true "User ID"
// @Success 200 {string} string "User unblocked successfully!"
// @Failure 400 {object} ErrorResponse "Invalid ID"
// @Failure 403 {object} ErrorResponse "Permission denied"
// @Failure 404 {object} ErrorResponse "User not found"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /users/{id}/unblock [patch]
// @Security ApiKeyAuth
func UnblockUserByID(c *gin.Context) {
	// Получаем роль пользователя из контекста...
	userRole := c.GetString(userRoleCtx)
	// if userRole == "" {
	// 	handleError(c, errs.ErrValidationFailed)
	// 	return
	// }

	// Проверяем, является ли пользователь администратором...
	if userRole != "Admin" {
		handleError(c, errs.ErrPermissionDenied)
		return
	}

	// Извлекаем ID пользователя...
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	// Логируем запрос на разблокировку пользователя...
	logger.Info.Printf("IP: [%s] requested to unblock user with ID: %d\n", c.ClientIP(), id)

	// Вызываем сервис для разблокировки пользователя...
	err = service.UnblockUserByID(uint(id))
	if err != nil {
		handleError(c, err)
		return
	}

	// Логируем успешную разблокировку...
	logger.Info.Printf("IP: [%s] successfully unblocked user with ID: %d\n", c.ClientIP(), id)

	// Отправляем клиенту ответ...
	c.JSON(http.StatusOK, gin.H{"message": "User unblocked successfully!"})
}

// ResetPassword
// @Summary Reset user password
// @Tags users
// @Description Reset a user's password (Admin only)
// @ID reset-password
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param input body models.SwagUser true "New password"
// @Success 200 {string} string "Password reset successfully!"
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Failure 403 {object} ErrorResponse "Permission denied"
// @Failure 404 {object} ErrorResponse "User not found"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /users/{id}/reset-password [patch]
// @Security ApiKeyAuth
func ResetPassword(c *gin.Context) {
	userRole := c.GetString(userRoleCtx)
	if userRole != "Admin" {
		handleError(c, errs.ErrPermissionDenied)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	var passwordData struct {
		NewPassword string `json:"new_password"`
	}

	if err := c.BindJSON(&passwordData); err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	err = service.ResetPassword(uint(id), passwordData.NewPassword)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password reset successfully!"})
}

// ChangeOwnPassword
// @Summary Change own password
// @Tags users
// @Description Change the logged-in user's password
// @ID change-own-password
// @Accept json
// @Produce json
// @Param input body models.SwagUser true "Old and new passwords"
// @Success 200 {string} string "Password changed successfully!"
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Failure 403 {object} ErrorResponse "Unauthorized password change"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /users/password [patch]
// @Security ApiKeyAuth
func ChangeOwnPassword(c *gin.Context) {
	userID := c.GetUint(userIDCtx)

	var passwordData struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}

	if err := c.BindJSON(&passwordData); err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	err := service.ChangeOwnPassword(userID, passwordData.OldPassword, passwordData.NewPassword)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password changed successfully!"})
}
