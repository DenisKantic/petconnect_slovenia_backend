package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"slovenia_petconnect/utils"
)

func AdoptPostUpload(c *gin.Context) {

	//var request models.AdoptPostCreateRequest
	////parse multipart form and validate
	//if err := c.ShouldBind(&request); err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": "Some fields are missing"})
	//	return
	//}

	claimsData, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	claims, ok := claimsData.(*utils.Claims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	userID := claims.UserID // getting the User ID from the jwt token which is going to be stored inside db for foreign key (relation with the post)

	fmt.Println("USER ID", userID)
}
