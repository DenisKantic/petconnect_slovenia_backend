package adoptPost

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"slovenia_petconnect/database"
	"slovenia_petconnect/models"
	"strconv"
)

func GetAdoptPosts(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit := 20
	offset := (page - 1) * limit

	var posts []models.AdoptPost
	var total int64

	// Get total count
	if err := database.DB.Model(&models.AdoptPost{}).Count(&total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not count posts"})
		return
	}

	// Get paginated data
	if err := database.DB.
		Preload("User").
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch posts"})
		return
	}

	//convert response format
	var response []models.AdoptPostFilterResponse
	for _, p := range posts {
		response = append(response, models.AdoptPostFilterResponse{
			ID:          p.ID,
			PostName:    p.PostName,
			Category:    p.Category,
			Description: p.Description,
			Sex:         p.Sex,
			Vaccinated:  p.Vaccinated,
			Chipped:     p.Chipped,
			Location:    p.Location,
			ImageURLs:   p.ImageURLs,
			CreatedAt:   p.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data":       response,
		"total":      total,
		"page":       page,
		"totalPages": int((total + int64(limit) - 1) / int64(limit)),
	})
}
