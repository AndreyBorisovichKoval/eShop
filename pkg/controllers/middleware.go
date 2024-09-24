// C:\GoProject\src\eShop\pkg\controllers\middleware.go

package controllers

import (
	"eShop/errs"
	"eShop/logger"
	"eShop/pkg/service"
	"net/http"
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
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "empty auth header",
		})
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "invalid auth header",
		})
		return
	}

	accessToken := headerParts[1]

	claims, err := service.ParseToken(accessToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Проверяем наличие роли, если её нет, возвращаем ошибку валидации...
	if claims.Role == "" {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": errs.ErrPermissionDenied.Error()})
		return
	}

	// Устанавливаем идентификатор пользователя и роль в контекст...
	c.Set(userIDCtx, claims.UserID)
	c.Set(userRoleCtx, claims.Role)

	c.Next()
}

// checkUserBlocked проверяет, заблокирован ли пользователь
func checkUserBlocked(c *gin.Context) {
	userID, exists := c.Get(userIDCtx)
	if !exists {
		logger.Warning.Println("User ID not found in context")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
		c.Abort()
		return
	}

	user, err := service.GetUserByID(userID.(uint))
	if err != nil {
		logger.Error.Printf("Error getting user by ID: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving user information"})
		c.Abort()
		return
	}

	if user.IsBlocked {
		logger.Warning.Printf("Blocked user attempting to access: User ID: %d", userID)
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied. User is blocked."})
		c.Abort()
		return
	}

	c.Next()
}

// CheckPasswordResetRequired проверяет, нужно ли пользователю сменить пароль
func CheckPasswordResetRequired(c *gin.Context) {
	userID, exists := c.Get(userIDCtx)
	if !exists {
		logger.Warning.Println("User ID not found in context")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
		c.Abort()
		return
	}

	user, err := service.GetUserByID(userID.(uint))
	if err != nil {
		logger.Error.Printf("Error getting user by ID: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving user information"})
		c.Abort()
		return
	}

	if user.PasswordResetRequired {
		logger.Warning.Printf("User with ID %d is required to reset their password.", userID)
		c.JSON(http.StatusForbidden, gin.H{"error": "Password reset required. Please change your password to continue."})
		c.Abort()
		return
	}

	c.Next()
}
