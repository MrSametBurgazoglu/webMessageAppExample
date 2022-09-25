package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"webMessageApp/database"
	"webMessageApp/messages/models"
	models2 "webMessageApp/people/models"
)

func CreateChat(c *gin.Context) {
	chat := models.Chat{}
	database.DB.Create(&chat)
	var people []models2.Person
	peopleIdStrings := c.Query("people_id")
	var peopleIdInt = make([]uint, len(peopleIdStrings))

	if err := database.DB.Find(people, peopleIdInt).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	for _, person := range people {
		err := database.DB.Model(person).Association("Chats").Append(&chat)
		if err != nil {
			return
		}
	}
}

func CreateMessageForChat(c *gin.Context) {
	var input CreateMessageInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var chat models.Chat
	if err := database.DB.Find(&chat, input.ChatID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Create book
	message := models.Message{Context: input.Context}
	database.DB.Create(&message)
	err := database.DB.Model(chat).Association("Messages").Append(&message)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": message})
}

func FindMessages(c *gin.Context) {
	var messages []models.Message
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
	message := models.Message{Context: input.Context}
	database.DB.Create(&message)

	c.JSON(http.StatusOK, gin.H{"data": message})
}

func FindMessage(c *gin.Context) { // Get model if exist
	var message models.Message

	if err := database.DB.Where("id = ?", c.Param("id")).First(&message).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": message})
}

func UpdateMessage(c *gin.Context) {
	// Get model if exist
	var group models.Message
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
	var group models.Message
	if err := database.DB.Where("id = ?", c.Param("id")).First(&group).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	database.DB.Delete(&group)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
