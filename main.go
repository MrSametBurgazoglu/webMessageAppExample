package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"webMessageApp/database"
	"webMessageApp/webApp"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	database.ConnectDatabase()

	webApp.SetRouter(r)

	r.LoadHTMLGlob("templates/*")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}
