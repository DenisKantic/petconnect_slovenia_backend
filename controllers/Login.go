package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"slovenia_petconnect/database"
	"slovenia_petconnect/models"
	"slovenia_petconnect/utils"
	"time"
)

func ManualLogin(c *gin.Context) {

	// check if user already has valid token (logged in)
	tokenCookie, err := c.Cookie("auth_token")
	if err == nil {
		// validate the token
		claims, err := utils.ValidateToken(tokenCookie)
		if err == nil && claims != nil {
			// token is valid prevent login again
			c.JSON(http.StatusBadRequest, gin.H{"error": "You are already logged in"})
			return
		}
	}

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

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to login"})
		return
	}

	c.SetCookie("auth_token", token, 60*60*24, "/", "", false, false)

	c.JSON(http.StatusOK, gin.H{"message": "You are logged in"})

}

func Logout(c *gin.Context) {
	c.SetCookie("auth_token", "", -1, "/", "", false, false)

	c.JSON(http.StatusOK, gin.H{"message": "Logout successfully"})
}
