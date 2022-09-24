package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HomePage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "hey"})
}

func ChatPage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "hey"})
}
