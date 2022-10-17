package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"webMessageApp/database"
	"webMessageApp/models"
	"webMessageApp/validators"
)

func CreateChat(c *gin.Context) {
	chat := models.Chat{}
	var people []models.Person
	var input validators.CreateChatInput

	if err := c.ShouldBindJSON(&input); err != nil {
		println(err.Error(), "error2")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	chat.ChatName = input.ChatName
	database.DB.Create(&chat)
	if err := database.DB.Find(&people, input.People).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	for _, person := range people {
		err := database.DB.Model(&person).Association("Chats").Append(&chat)
		if err != nil {
			return
		}
	}
}

func CreateMessageForChat(c *gin.Context) {
	var input validators.CreateMessageInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var chat models.Chat
	if err := database.DB.Find(&chat, input.Chat).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Create message
	message := models.Message{Context: input.Context}
	database.DB.Create(&message)
	err := database.DB.Model(chat).Association("Messages").Append(&message)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": message})
}

func GetChatMessages(c *gin.Context) { // Get model if exist
	var messages []models.Message

	if err := database.DB.Preload("Person").Where("chat_id = ?", c.Param("id")).Order("id").Find(&messages).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	println(messages)
	c.JSON(http.StatusOK, gin.H{"data": messages})

}

func FindMessages(c *gin.Context) {
	var messages []models.Message
	database.DB.Find(&messages)

	c.JSON(http.StatusOK, gin.H{"data": messages})
}

func CreateMessage(c *gin.Context) {
	// Validate input
	var input validators.CreateMessageInput
	if err := c.ShouldBindJSON(&input); err != nil {
		println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create message
	message := models.Message{Context: input.Context, PersonID: input.Person, ChatID: input.Chat}
	database.DB.Create(&message)
	var chat models.Chat
	database.DB.Find(&chat, input.Chat)

	database.DB.Model(&chat).Association("Messages").Append(&message)

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
	var input validators.UpdateMessageInput
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
