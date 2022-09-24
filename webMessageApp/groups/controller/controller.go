package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"webMessageApp/database"
	"webMessageApp/groups/models"
)

func FindGroups(c *gin.Context) {
	var groups []models.Group
	database.DB.Find(&groups)

	c.JSON(http.StatusOK, gin.H{"data": groups})
}

func CreateGroup(c *gin.Context) {
	// Validate input
	var input CreateGroupInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	group := models.Group{GroupName: input.Title}
	database.DB.Create(&group)

	c.JSON(http.StatusOK, gin.H{"data": group})
}

func FindGroup(c *gin.Context) { // Get model if exist
	var group models.Group

	if err := database.DB.Where("id = ?", c.Param("id")).First(&group).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": group})
}

func UpdateGroup(c *gin.Context) {
	// Get model if exist
	var group models.Group
	if err := database.DB.Where("id = ?", c.Param("id")).First(&group).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateGroupInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Model(&group).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": group})
}

func DeleteGroup(c *gin.Context) {
	// Get model if exist
	var group models.Group
	if err := database.DB.Where("id = ?", c.Param("id")).First(&group).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	database.DB.Delete(&group)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
