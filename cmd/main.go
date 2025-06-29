package main

import (
	"github.com/gin-gonic/gin"
	"slovenia_petconnect/config"
	"slovenia_petconnect/database"
	"slovenia_petconnect/middleware"
	"slovenia_petconnect/utils"
)

func main() {

	utils.InitLogger() // logging

	r := gin.New()
	r.Use(middleware.ZapLogger())

	// load .env
	config.LoadEnv()
	//connect Postgres database
	database.ConnectDB()

	// ----------------auth routes (login, singup)----------------
	routes.SetupAuthRoutes(r)
	//------------------------------------------------

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Pong tests"})
	})

	r.Run(":8080")
}
