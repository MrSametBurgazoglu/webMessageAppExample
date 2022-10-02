package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"webMessageApp/database"
	"webMessageApp/models"
	"webMessageApp/validators"
)

func FindPeople(c *gin.Context) {
	var people []models.Person
	database.DB.Preload("Chats").Find(&people)

	c.JSON(http.StatusOK, gin.H{"data": people})
}

func CreatePerson(c *gin.Context) {
	// Validate input
	var nickname = c.PostForm("Nickname")
	if nickname == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nickname cannot be empty"})
		return
	}

	// Create person
	person := models.Person{Nickname: nickname}
	database.DB.Create(&person)

	c.JSON(http.StatusOK, gin.H{"data": person})
	return
}

func FindPerson(c *gin.Context) { // Get model if exist
	var person models.Person

	if err := database.DB.Preload("Chats").Where("id = ?", c.Param("id")).First(&person).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": person})

}

func GetPersonChats(c *gin.Context) { // Get model if exist
	var person models.Person

	if err := database.DB.Preload("Chats").Where("id = ?", c.Param("id")).First(&person).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	for _, chat := range person.Chats {
		var message models.Message
		if chat.ChatName == "" {
			var chatDb models.Chat
			database.DB.Preload("Person").Find(&chat).First(&chatDb)
			for _, p := range chatDb.Person {
				if p.ID != person.ID {
					chat.ChatName = p.Nickname
				}
			}
		}
		if err := database.DB.Model(&message).Where("chat_id = ?", chat.ID).Last(&message).Error; err != nil {
			chat.ChatLastMessage = "No message yet"
		} else {
			chat.ChatLastMessage = message.Context
		}
	}
	for _, chat := range person.Chats {
		println(chat.ChatLastMessage, "last message")
	}
	c.JSON(http.StatusOK, gin.H{"data": person.Chats})

}

func UpdatePerson(c *gin.Context) {
	// Get model if exist
	var person models.Person
	if err := database.DB.Where("id = ?", c.Param("id")).First(&person).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input validators.UpdatePersonInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Model(&person).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": person})
}

func DeletePerson(c *gin.Context) {
	// Get model if exist
	var person models.Person
	if err := database.DB.Where("id = ?", c.Param("id")).First(&person).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	database.DB.Delete(&person)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
