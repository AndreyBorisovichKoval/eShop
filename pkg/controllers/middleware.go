// C:\GoProject\src\eShop\pkg\controllers\middleware.go

package controllers

import (
	"eShop/errs"
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
	// Проверяем наличие роли, если её нет, возвращаем ошибку валидации...
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
