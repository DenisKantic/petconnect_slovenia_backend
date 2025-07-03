package routes

import (
	"github.com/gin-gonic/gin"
	"slovenia_petconnect/controllers"
)

func SetupAuthRoutes(r *gin.Engine) {

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/signup/email", controllers.RegisterWithEmailUser)
		authGroup.POST("/login/manual", controllers.ManualLogin)
		authGroup.POST("/logout", controllers.Logout)
	}
}
