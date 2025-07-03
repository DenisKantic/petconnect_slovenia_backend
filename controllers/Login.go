package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"slovenia_petconnect/database"
	"slovenia_petconnect/models"
	"time"
)

func ManualLogin(c *gin.Context) {

	var login models.ManualLoginUserRequest
	var user models.User
	// userIP := c.ClientIP()

	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	fmt.Println("USERNAME", login.Username)

	if err := database.DB.Where("username = ? AND PROVIDER = ?", login.Username, "manual").First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(*user.PasswordHash), []byte(login.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password"})
		return
	}

	// save the last date of user logged in
	now := time.Now()
	user.LastLogin = &now
	if err := database.DB.Model(&user).Update("last_login", user.LastLogin).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update last_login"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "You are logged in"})

}
