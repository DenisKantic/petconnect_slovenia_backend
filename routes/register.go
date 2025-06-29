package routes

import (
	"github.com/gin-gonic/gin"
	"slovenia_petconnect/controllers"
)

func RegisterWithEmailUser(r *gin.Engine) {

	r.POST("/register", controllers.RegisterNewUser)
}
