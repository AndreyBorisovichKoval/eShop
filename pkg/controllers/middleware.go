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

// CheckUserBlocked проверяет, заблокирован ли пользователь
func checkUserBlocked() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get(userIDCtx) // Получаем ID пользователя из контекста
		if !exists {
			logger.Warning.Println("User ID not found in context")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
			c.Abort()
			return
		}

		// Получаем пользователя по ID
		user, err := service.GetUserByID(userID.(uint))
		if err != nil {
			logger.Error.Printf("Error getting user by ID: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving user information"})
			c.Abort()
			return
		}

		// Проверяем, заблокирован ли пользователь
		if user.IsBlocked {
			logger.Warning.Printf("Blocked user attempting to access: User ID: %d", userID)
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied. User is blocked."})
			c.Abort()
			return
		}

		c.Next() // Если пользователь не заблокирован, продолжаем выполнение запроса
	}
}
