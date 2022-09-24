package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"webMessageApp/database"
	"webMessageApp/messages/models"
)

func FindMessages(c *gin.Context) {
	var messages []models.PersonMessage
	database.DB.Find(&messages)

	c.JSON(http.StatusOK, gin.H{"data": messages})
}

func CreateMessage(c *gin.Context) {
	// Validate input
	var input CreateMessageInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	message := models.PersonMessage{Context: input.Title}
	database.DB.Create(&message)

	c.JSON(http.StatusOK, gin.H{"data": message})
}

func FindMessage(c *gin.Context) { // Get model if exist
	var message models.PersonMessage

	if err := database.DB.Where("id = ?", c.Param("id")).First(&message).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": message})
}

func UpdateMessage(c *gin.Context) {
	// Get model if exist
	var group models.PersonMessage
	if err := database.DB.Where("id = ?", c.Param("id")).First(&group).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateMessageInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Model(&group).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": group})
}

func DeleteMessage(c *gin.Context) {
	// Get model if exist
	var group models.PersonMessage
	if err := database.DB.Where("id = ?", c.Param("id")).First(&group).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	database.DB.Delete(&group)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
