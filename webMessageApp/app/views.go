package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Main website",
	})
}

func ChatPage(c *gin.Context) {
	c.HTML(http.StatusOK, "chat.html", gin.H{
		"nickname": "Samet Burgazoğlu",
	})
}
