package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"slovenia_petconnect/database"
	"slovenia_petconnect/models"
	"slovenia_petconnect/utils"
)

func RegisterWithEmailUser(c *gin.Context) {

	var request models.RegisterUserRequest

	// validating JSON body
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You didn't fill all the fields"})
		return
	}

	// checking if user already exists
	var existingUser models.User
	if err := database.DB.Where("email = ?", request.Email).First(&existingUser).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already in use"})
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
		Username:     request.Email,
		Email:        request.Email,
		PasswordHash: &hashedPassword,
		Provider:     "manual",
		Location:     request.Location,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	// success
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}
