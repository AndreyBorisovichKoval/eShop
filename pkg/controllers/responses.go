package controllers

// TokenResponse представляет ответ с токеном доступа и идентификатором пользователя...
type accessTokenResponse struct {
	AccessToken string `json:"access_token"` // JWT токен для аутентификации пользователя...
}

type defaultResponse struct {
	Message string `json:"message"`
}

func newDefaultResponse(message string) defaultResponse {
	return defaultResponse{
		Message: message,
	}
}
