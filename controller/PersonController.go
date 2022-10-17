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
					println(p.ID, person.ID)
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

func GetPersonAddFriendList(c *gin.Context) {
	var personIDList []int
	var personAddFriendList []models.Person
	result := database.DB.Raw("SELECT id FROM people WHERE id NOT IN (SELECT DISTINCT person_id FROM people_chats WHERE chat_id IN (SELECT chat_id from people_chats WHERE person_id = ?)) AND id <> ?;", c.Param("id"), c.Param("id")).Scan(&personIDList)
	if result.RowsAffected != 0 {
		database.DB.Find(&personAddFriendList, personIDList)
	}
	c.JSON(http.StatusOK, gin.H{"data": personAddFriendList})
}

func PersonAddFriendList(c *gin.Context) {
	// Get model if exist

	// Validate input
	var input validators.CreateChatInput
	if err := c.ShouldBindJSON(&input); err != nil {
		println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var chat = models.Chat{ChatName: ""}
	database.DB.Create(&chat)
	for _, personId := range input.People {
		var person models.Person
		database.DB.Find(&person, personId)
		println(personId, "added to association")
		database.DB.Model(&person).Association("Chats").Append(&chat)
	}
	c.JSON(http.StatusOK, gin.H{"data": chat})
}

func GetChatCreateFriendList(c *gin.Context) {
	var personIDList []int
	var personAddFriendList []models.Person
	result := database.DB.Raw("SELECT DISTINCT person_id FROM people_chats WHERE chat_id IN (SELECT chat_id from people_chats WHERE person_id = ?) AND person_id <> ?;", c.Param("id"), c.Param("id")).Scan(&personIDList)
	if result.RowsAffected != 0 {
		database.DB.Find(&personAddFriendList, personIDList)
	}
	c.JSON(http.StatusOK, gin.H{"data": personAddFriendList})
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
