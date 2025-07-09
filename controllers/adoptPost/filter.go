package adoptPost

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"slovenia_petconnect/database"
	"slovenia_petconnect/models"
	"time"
)

func FilterAdoptPost(c *gin.Context) {
	location := c.Query("location")
	category := c.Query("category")
	createdAfter := c.Query("created_after")

	var posts []models.AdoptPost
	query := database.DB.Model(&models.AdoptPost{})

	if location != "" {
		query = query.Where("location = ?", location)
	}

	if category != "" {
		query = query.Where("category = ?", category)
	}

	if createdAfter != "" {
		parsedDate, err := time.Parse("2006-01-02", createdAfter)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
			return
		}

		query = query.Where("created_at > ?", parsedDate)
	}

	if err := query.Order("created_at DESC").Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch posts"})
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

	c.JSON(http.StatusOK, gin.H{"result": response})
}
