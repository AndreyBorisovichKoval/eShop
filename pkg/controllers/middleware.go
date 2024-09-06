package controllers

import (
	"eShop/logger"
	"eShop/pkg/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	sellerIDCtx         = "sellerID"
	sellerRoleCtx       = "sellerRole"
)

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

	if accessToken == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "token is empty",
		})
		return
	}

	claims, err := service.ParseToken(accessToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Логирование claims для отладки
	logger.Info.Printf("Claims: %+v", claims)

	c.Set(sellerIDCtx, claims.sellerID)
	c.Set(sellerRoleCtx, claims.Role)
	c.Next()
}
