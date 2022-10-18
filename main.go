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
	r.SetTrustedProxies([]string{"192.168.1.2"})
	println("\n\n")
	println(os.Getenv("RENDER_EXTERNAL_HOSTNAME"))
	println("\n\n")

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
