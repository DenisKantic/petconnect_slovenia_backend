package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"slovenia_petconnect/database"
	"slovenia_petconnect/models"
	"slovenia_petconnect/utils"
)

func AdoptPostUpload(c *gin.Context) {

	var request models.AdoptPostCreateRequest

	claims, ok := utils.GetClaims(c)
	if !ok {
		return // error already returned in the utils
	}

	//parse multipart form and validate
	if err := c.ShouldBind(&request); err != nil {
		fmt.Println("ERROR BIND", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Some fields are missing"})
		return
	}

	// ✅ Manually validate that images are uploaded
	// Check for real uploaded files (not just empty placeholders)
	var realFiles []*multipart.FileHeader
	for _, file := range request.Images {
		if file != nil && file.Filename != "" && file.Size > 0 {
			realFiles = append(realFiles, file)
		}
	}

	if len(realFiles) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please upload at least one image"})
		return
	}

	fmt.Println("IMAGES LENGTH", len(request.Images))

	userID := claims.UserID // getting the User ID from the jwt token which is going to be stored inside db for foreign key (relation with the post)

	// creating unique folder for storing uploaded images
	uniqueID := uuid.New().String()[:8]
	imagesFolderPath := fmt.Sprintf("./static/adoptImages/%s_%s", request.PostName, uniqueID)

	if err := os.MkdirAll(imagesFolderPath, os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store images"})
		return
	}

	// process and save uploaded images
	var savedImagePaths []string
	for _, fileHeader := range request.Images {
		imageFilePath := filepath.Join(imagesFolderPath, fileHeader.Filename)

		srcFile, err := fileHeader.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error opening image files"})
			return
		}

		defer srcFile.Close()

		dstFile, err := os.Create(imageFilePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving image files"})
			return
		}
		defer dstFile.Close()

		if _, err := io.Copy(dstFile, srcFile); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error copying image files to server"})
			return
		}

		savedImagePaths = append(savedImagePaths, imageFilePath)
	}

	// Create AdoptPost model and save to database
	post := models.AdoptPost{
		UserID:      userID,
		PostName:    request.PostName,
		Category:    request.Category,
		Description: request.Description,
		Sex:         request.Sex,
		Vaccinated:  *request.Vaccinated,
		Chipped:     *request.Chipped,
		Location:    request.Location,
		ImageURLs:   savedImagePaths,
	}

	if err := database.DB.Create(&post).Error; err != nil {
		os.RemoveAll(imagesFolderPath) // cleanup on DB failure
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save post"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Post successfully created",
		"post":    post,
	})
}
