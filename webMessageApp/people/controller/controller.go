package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"webMessageApp/database"
	"webMessageApp/people/models"
)

func FindPeople(c *gin.Context) {
	var people []models.Person
	database.DB.Find(&people)

	c.JSON(http.StatusOK, gin.H{"data": people})
}

func CreatePerson(c *gin.Context) {
	// Validate input
	var input CreatePersonInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	person := models.Person{Nickname: input.Title}
	database.DB.Create(&person)

	c.JSON(http.StatusOK, gin.H{"data": person})
}

func FindPerson(c *gin.Context) { // Get model if exist
	var person models.Person

	if err := database.DB.Where("id = ?", c.Param("id")).First(&person).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": person})
}

func UpdatePerson(c *gin.Context) {
	// Get model if exist
	var person models.Person
	if err := database.DB.Where("id = ?", c.Param("id")).First(&person).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdatePersonInput
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
