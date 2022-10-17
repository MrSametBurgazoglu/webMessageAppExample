package webApp

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"webMessageApp/database"
	"webMessageApp/models"
)

func HomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Web Message App",
	})
}

func ChatPage(c *gin.Context) {
	var person models.Person
	database.DB.First(&person, c.Param("id"))
	c.HTML(http.StatusOK, "chat.html", gin.H{
		"nickname": person.Nickname,
		"id":       c.Param("id"),
	})
}

func SignupPage(c *gin.Context) {
	c.HTML(http.StatusOK, "chat.html", gin.H{
		"nickname": "Samet Burgazoğlu",
	})
}

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "chat.html", gin.H{
		"nickname": "Samet Burgazoğlu",
	})
}

func CreatePersonPage(c *gin.Context) {
	c.HTML(http.StatusOK, "create.html", gin.H{
		"nickname": "Samet Burgazoğlu",
	})
}
