// Package controllers  package used for handling API methods
package controllers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"slovenia_petconnect/database"
	"slovenia_petconnect/models"
	"slovenia_petconnect/utils"
	"time"
)

// RegisterWithEmailUser handles user registration with email and password*
//
// It performs the following steps:
//
//   - parses and validates the incoming JSON body with user model
//
//   - checks if the user already exists by email
//     -hashes the user's password securely
//     -stores the new user in the database
//
//   - returns appropriate HTTP status codes and messages
//
// Expected input JSON:
//
//	{
//		"email": "user@test.com",
//		"username": "John Doe",
//		"password: "yourPassword",
//		"location": "New York",
//	}
//
// Possible responses:
//
//   - 201 Created: user successfully registered
//
//   - 400 Bad request: validation failed or email in use
//
//   - 500 Internal Server Error: on DB or hashing failure
func RegisterWithEmailUser(c *gin.Context) {

	userIP := c.ClientIP()

	fmt.Println("USER IP ", []string{userIP})
	var request models.RegisterUserRequest

	// validating JSON body
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You didn't fill all the fields"})
		return
	}

	// checking if user already exists
	var existingUser models.User
	if err := database.DB.Where("email = ?", request.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already in use"})
		return
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not find the user"})
		return
	}

	// hash password
	hashedPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Problem with registration"})
		return
	}

	// create new user
	user := models.User{
		Username:     request.Username,
		Email:        request.Email,
		PasswordHash: &hashedPassword,
		Provider:     "manual",
		ProviderID:   nil,
		Location:     request.Location,
		//ClientIP:     userIP, // wrap in slice
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		LastLogin: nil,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	// success
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}
