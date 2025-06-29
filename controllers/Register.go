package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"slovenia_petconnect/database"
	"slovenia_petconnect/models"
)

func RegisterNewUser(c *gin.Context) {

	var request models.RegisterUserRequest

	// validating JSON body
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You didn't fill all the fields"})
		return
	}

	// checking if user already exists
	var existingUser models.User
	if err := database.DB.Where("email = ?", request.Email).First(&existingUser).Error; err != nil {

	}
}
