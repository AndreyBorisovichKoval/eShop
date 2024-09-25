// C:\GoProject\src\eShop\models\request_history.go

package models

import (
	"time"
)

// RequestHistory defines the structure for capturing all user requests
type RequestHistory struct {
	ID uint `gorm:"primaryKey" json:"id"` // Unique ID for each request
	// UserID    uint      `json:"user_id"`                          // User ID making the request
	UserID          uint      `gorm:"column:user_id;not null;foreignKey:" json:"user_id"` // User ID (no foreign key constraint)
	Username        string    `json:"username"`                                           // Username of the user
	FullName        string    `json:"full_name"`                                          // Full name of the user
	Email           string    `json:"email"`                                              // Email of the user
	Phone           string    `json:"phone"`                                              // Phone number of the user
	Role            string    `json:"role"`                                               // Role of the user (admin, seller, etc.)
	Path            string    `json:"path"`                                               // The path or route requested
	Method          string    `json:"method"`                                             // HTTP method (GET, POST, etc.)
	ClientIPAddress string    `json:"client_ip_address"`                                  // IP address of the requester
	CreatedAt       time.Time `gorm:"autoCreateTime" json:"created_at"`                   // Timestamp of when the request was made
}

// TableName sets the table name for GORM
func (RequestHistory) TableName() string {
	return "request_history"
}
