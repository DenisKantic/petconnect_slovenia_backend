package models

import "time"

type User struct {
	ID           uint    `gorm:"primary_key"`
	Username     string  `gorm:"uniqueIndex;size:50" binding:"required"`
	Email        string  `gorm:"uniqueIndex;not null" binding:"required"`
	PasswordHash *string // pointer: nil if using OAuth
	Provider     string  // "manual" or "google"
	ProviderID   *string `gorm:"default:null"` // e.g Google-s user ID, nil for manual
	Location     string
	// ClientIP     []string `gorm:"type:text[]"`
	CreatedAt time.Time
	UpdatedAt time.Time
	LastLogin *time.Time `gorm:"default:null"`
}

// RegisterUserRequest Validation for manual registration
type RegisterUserRequest struct {
	Username string `json:"username" binding:"required,min=3,max=25"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Location string `json:"location" binding:"required"`
}

type ManualLoginUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
