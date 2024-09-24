// C:\GoProject\src\eShop\models\user_transform.go

package models

// ConvertToUserResponse преобразует полную модель User в UserResponse
func ConvertToUserResponse(user User) UserResponse {
	return UserResponse{
		ID:                    user.ID,
		FullName:              user.FullName,
		Username:              user.Username,
		Email:                 user.Email,
		Phone:                 user.Phone,
		PasswordResetRequired: user.PasswordResetRequired,
		Role:                  user.Role,
		IsBlocked:             user.IsBlocked,
		BlockedAt:             user.BlockedAt,
		CreatedAt:             user.CreatedAt,
	}
}

// ConvertToUserResponses преобразует массив моделей User в массив UserResponse
func ConvertToUserResponses(users []User) []UserResponse {
	var userResponses []UserResponse
	for _, user := range users {
		userResponses = append(userResponses, ConvertToUserResponse(user))
	}
	return userResponses
}
