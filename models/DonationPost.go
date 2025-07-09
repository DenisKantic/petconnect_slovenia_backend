// Package models containing all models for user, all type of creating post, filtering request and models for validating incoming request
package models

import (
	"github.com/lib/pq"
	"mime/multipart"
	"time"
)

type DonationPost struct {
	ID             uint           `gorm:"primary_key;auto_increment" json:"id"`
	UserID         uint           `gorm:"not null"`
	User           User           `gorm:"foreignKey:UserID"`
	PostName       string         `gorm:"size:255;not null" json:"postName"`
	AnimalCategory string         `gorm:"size:50;not null" json:"animal_category"`
	PostCategory   string         `gorm:"size:50;not null" json:"post_category"`
	Description    string         `gorm:"type:text;not null" json:"description"`
	Location       string         `gorm:"size:100;not null" json:"location"`
	ImageURLs      pq.StringArray `gorm:"type:text[]" json:"imageUrls"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
}

// DonationPostCreateRequest struct is written to test validation on all fields.
// In case if failed some field, it will notify user and the POST request will be rejected
type DonationPostCreateRequest struct {
	PostName       string                  `form:"post_name" binding:"required,max=40"`
	AnimalCategory string                  `form:"animal_category" binding:"required,max=10"`
	PostCategory   string                  `form:"post_category" binding:"required,max=10"`
	Description    string                  `form:"description" binding:"required,max=2000"`
	Location       string                  `form:"location" binding:"required"`
	Images         []*multipart.FileHeader `form:"images" binding:"required"` // <-- For receiving files
}

// DonationPostFilterResponse is created to send to the frontend the clean response, without information from
// "User" module (user infos), because it's related via foreign key through GORM so it will automatically display also user table
// With this we have more control what are we going to display in the response
type DonationPostFilterResponse struct {
	ID             uint      `json:"id"`
	PostName       string    `json:"postName"`
	AnimalCategory string    `json:"animal_category"`
	PostCategory   string    `json:"post_category"`
	Description    string    `json:"description"`
	Location       string    `json:"location"`
	ImageURLs      []string  `json:"imageUrls"`
	CreatedAt      time.Time `json:"createdAt"`
}
