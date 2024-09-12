// C:\GoProject\src\eShop\pkg\service\jwt.go

package service

import (
	"eShop/configs"
	"eShop/logger"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// CustomClaims определяет кастомные поля токена
type CustomClaims struct {
	UserID   uint   `json:"user_id"`
	Role     string `json:"role"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenerateToken генерирует JWT токен с кастомными полями
func GenerateToken(userID uint, username string, role string) (string, error) {
	// Логируем попытку генерации токена...
	logger.Info.Printf("Generating token for user ID: %d, Username: %s, Role: %s", userID, username, role)

	// Устанавливаем данные для токена...
	claims := CustomClaims{
		UserID:   userID,
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(configs.AppSettings.AuthParams.JwtTtlMinutes)).Unix(),
			Issuer:    configs.AppSettings.AppParams.ServerName,
		},
	}

	// Создаём новый токен...
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Подписываем токен и логируем успешную генерацию...
	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		logger.Error.Printf("Error generating token for user ID: %d, error: %v", userID, err)
		return "", err
	}

	logger.Info.Printf("Token generated successfully for user ID: %d", userID)
	return signedToken, nil
}

// ParseToken парсит JWT токен и возвращает кастомные поля...
func ParseToken(tokenString string) (*CustomClaims, error) {
	logger.Info.Printf("Attempting to parse token...")

	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Проверяем метод подписи токена
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			logger.Error.Printf("Unexpected signing method: %v", token.Header["alg"])
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		logger.Error.Printf("Error parsing token: %v", err)
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		logger.Info.Printf("Token parsed successfully!")
		return claims, nil
	}

	logger.Error.Println("Invalid token")
	return nil, fmt.Errorf("invalid token")
}
