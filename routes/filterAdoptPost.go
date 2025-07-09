package routes

import (
	"github.com/gin-gonic/gin"
	"slovenia_petconnect/controllers/adoptPost"
)

func SetupGetRoutes(r *gin.Engine) {

	postGroup := r.Group("/post")
	{
		postGroup.GET("/filter/adopt-post", adoptPost.FilterAdoptPost)
		postGroup.GET("/pagination/adopt-post", adoptPost.GetAdoptPosts)
	}
}
