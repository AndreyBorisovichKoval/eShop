// C:\GoProject\src\eShop\pkg\service\jwt.go

package service

import (
	"eShop/configs"
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
	// IssuerTMP := configs.AppSettings.AppParams.ServerName
	// fmt.Printf("issuer: %s\n", IssuerTMP)
	// ttlMinutesStr := configs.AppSettings.AuthParams.JwtTtlMinutes
	// fmt.Printf("ttl_minutes: %s\n", ttlMinutesStr)
	// ttlMinutes, _ := strconv.Atoi(ttlMinutesStr)
	// duration := time.Duration(ttlMinutes) * time.Minute
	// fmt.Printf("duration: %d\n", duration)
	// expiresAt := time.Now().Add(duration).Unix()

	// // Debugging: Print time now and expiration time
	// fmt.Printf("Current time         : %s\n", time.Now().Format(time.RFC3339))
	// fmt.Printf("Token expiration time: %s\n", time.Unix(expiresAt, 0).Format(time.RFC3339))

	claims := CustomClaims{
		UserID:   userID,
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			// ExpiresAt: expiresAt,
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(configs.AppSettings.AuthParams.JwtTtlMinutes)).Unix(),
			// ExpiresAt: time.Now().Add(time.Minute * 60).Unix(), // токен истекает через 1 час
			Issuer: configs.AppSettings.AppParams.ServerName,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
}

// ParseToken парсит JWT токен и возвращает кастомные поля
func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Проверяем метод подписи токена
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
