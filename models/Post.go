package models

import (
	"mime/multipart"
	"time"
)

// AdoptPost is the GORM model that represents a post in the database.
type AdoptPost struct {
	ID          uint      `gorm:"primary_key;auto_increment" json:"id"`
	UserID      uint      `gorm:"not null"`
	User        User      `gorm:"foreignKey:UserID"`
	PostName    string    `gorm:"size:255;not null" json:"postName"`
	Category    string    `gorm:"size:50;not null" json:"category"`
	Description string    `gorm:"type:text;not null" json:"description"`
	Sex         string    `gorm:"size:10;not null" json:"sex"`
	Vaccinated  bool      `json:"vaccinated"`
	Chipped     bool      `json:"chipped"`
	Location    string    `gorm:"size:100;not null" json:"location"`
	ImageURLs   []string  `gorm:"type:text" json:"imageUrls"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type AdoptPostCreateRequest struct {
	PostName    string                  `form:"postName" binding:"required,max=40"`
	Category    string                  `form:"category" binding:"required,max=10"`
	Description string                  `form:"description" binding:"required,max=2000"`
	Sex         string                  `form:"sex" binding:"required"`
	Vaccinated  string                  `form:"vaccinated" binding:"required"`
	Chipped     string                  `form:"chipped" binding:"required"`
	Location    string                  `form:"location" binding:"required"`
	Images      []*multipart.FileHeader `form:"images" binding:"required"` // <-- For receiving files
}
