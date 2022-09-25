package models

import (
	"gorm.io/gorm"
	models2 "webMessageApp/messages/models"
	"webMessageApp/people/models"
)

type Group struct {
	gorm.Model
	GroupName string          `json:"group_name"`
	People    []models.Person `gorm:"many2many:group_people;"`
	Chat      models2.Chat
}
