package main

import (
	"github.com/gin-gonic/gin"
	"webMessageApp/database"
	"webMessageApp/webApp"
)

func main() {
	r := gin.Default()

	database.ConnectDatabase()

	webApp.SetRouter(r)

	r.LoadHTMLGlob("templates/*")

	r.Run()
}
