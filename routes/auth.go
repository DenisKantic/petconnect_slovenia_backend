package rkoutes

import (
	"github.com/gin-gonic/gin"
	"slovenia_petconnect/controllers"
)

func SetupAuthRoutes(r *gin.Engine) {

	registerGroup := r.Group("/auth")
	{
		registerGroup.POST("/signup/email", controllers.RegisterWithEmailUser)
	}
}
