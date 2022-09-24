package main

import (
	"github.com/gin-gonic/gin"
	"webMessageApp/app"
	"webMessageApp/database"
)

func main() {
	r := gin.Default()

	database.ConnectDatabase()

	app.SetRouter(r)

	r.LoadHTMLGlob("templates/*")

	r.Run()
}
