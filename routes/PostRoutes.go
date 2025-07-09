package routes

import (
	"github.com/gin-gonic/gin"
	"slovenia_petconnect/controllers/adoptPost"
	"slovenia_petconnect/middleware"
)

func SetupPostRoutes(r *gin.Engine) {

	postGroup := r.Group("/post")
	postGroup.Use(middleware.AuthMiddleware())
	{
		postGroup.POST("/create-post/adopt", adoptPost.UploadPost)
	}
}
