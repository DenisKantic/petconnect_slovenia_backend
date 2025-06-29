package models

import "time"

type User struct {
	ID           uint    `gorm:"primary_key"`
	Username     string  `gorm:"uniqueIndex;size:50" binding:"required"`
	Email        string  `gorm:"uniqueIndex;not null"`
	PasswordHash *string // pointer: nil if using OAuth
	Provider     string  // "manual" or "google"
	ProviderID   *string // e.g Google-s user ID, nil for manual
	Location     string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	LastLogin    *time.Time
}

// RegisterUserRequest Validation for manual registration
type RegisterUserRequest struct {
	Username string `json:"username" binding:"required,min=3,max=25"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Location string `json:"location" binding:"required"`
}
