// C:\GoProject\src\eShop\pkg\controllers\middleware.go

package controllers

import (
	"eShop/errs"
	"eShop/logger"

	// "eShop/pkg/middleware"
	"eShop/pkg/service"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userIDCtx           = "userID"
	userRoleCtx         = "userRole"
)

// Middleware для проверки аутентификации и наличия роли...
func checkUserAuthentication(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)

	if header == "" {
		handleError(c, errs.ErrEmptyAuthHeader) // Используем кастомную ошибку
		c.Abort()
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		handleError(c, errs.ErrInvalidAuthHeader) // Используем кастомную ошибку
		c.Abort()
		return
	}

	accessToken := headerParts[1]
	claims, err := service.ParseToken(accessToken)
	if err != nil {
		handleError(c, errs.ErrTokenParsingFailed) // Используем кастомную ошибку
		c.Abort()
		return
	}

	if claims.Role == "" {
		handleError(c, errs.ErrUserNotAuthenticated) // Используем кастомную ошибку
		c.Abort()
		return
	}

	c.Set(userIDCtx, claims.UserID)
	c.Set(userRoleCtx, claims.Role)

	// Логируем запрос пользователя, вызвав всего одну строку
	// Логируем запрос пользователя, вызвав всего одну строку
	// Логируем запрос пользователя, вызвав всего одну строку
	// _ = service.LogUserRequest(claims.UserID, c.Request.URL.Path, c.Request.Method, c.ClientIP())
	_ = service.LogUserRequest(c, claims.UserID)

	c.Next()
}

// checkUserBlocked проверяет, заблокирован ли пользователь...
func checkUserBlocked(c *gin.Context) {
	userID, exists := c.Get(userIDCtx)
	if !exists {
		handleError(c, errs.ErrUserNotAuthenticated) // Используем кастомную ошибку
		c.Abort()
		return
	}

	user, err := service.GetUserByID(userID.(uint))
	if err != nil {
		handleError(c, err)
		c.Abort()
		return
	}

	if user.IsBlocked {
		logger.Warning.Printf("Blocked user attempting to access: User ID: %d", userID)
		handleError(c, errs.ErrUserBlocked) // Используем новую кастомную ошибку
		c.Abort()
		return
	}

	c.Next()
}

// CheckPasswordResetRequired проверяет, нужно ли пользователю сменить пароль...
func CheckPasswordResetRequired(c *gin.Context) {
	userID, exists := c.Get(userIDCtx)
	if !exists {
		handleError(c, errs.ErrUserNotAuthenticated) // Используем кастомную ошибку
		c.Abort()
		return
	}

	user, err := service.GetUserByID(userID.(uint))
	if err != nil {
		handleError(c, err)
		c.Abort()
		return
	}

	if user.PasswordResetRequired {
		handleError(c, errs.ErrPasswordResetRequired) // Используем кастомную ошибку
		c.Abort()
		return
	}

	c.Next()
}

// Middleware для проверки роли администратора
func checkAdminRole(c *gin.Context) {
	userRole := c.GetString(userRoleCtx)          // Используем твой ключ для получения роли
	logger.Info.Printf("User role: %s", userRole) // Логируем роль для отладки
	if userRole != "Admin" {
		handleError(c, errs.ErrPermissionDeniedOnlyForAdmin) // Используем кастомную ошибку
		c.Abort()
		return
	}
	c.Next()
}

// Middleware для проверки роли администратора или менеджера
func checkManagerOrAdminRole(c *gin.Context) {
	userRole := c.GetString(userRoleCtx)          // Используем твой ключ для получения роли
	logger.Info.Printf("User role: %s", userRole) // Логируем роль для отладки
	if userRole != "Manager" && userRole != "Admin" {
		handleError(c, errs.ErrPermissionDeniedOnlyForAdminOrManager) // Используем кастомную ошибку
		c.Abort()
		return
	}
	c.Next()
}
