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
	err := r.SetTrustedProxies([]string{"192.168.1.2", os.Getenv("RENDER_EXTERNAL_HOSTNAME"), "0.0.0.0"})
	if err != nil {
		println("error:%s", err)
	}

	database.ConnectDatabase()

	webApp.SetRouter(r)

	r.LoadHTMLGlob("templates/*")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	println("port:", port)

	if err := r.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}
