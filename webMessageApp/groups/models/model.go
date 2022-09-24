package models

import (
	"gorm.io/gorm"
	"webMessageApp/people/models"
)

type Group struct {
	gorm.Model
	GroupName string          `json:"group_name"`
	People    []models.Person `gorm:"many2many:group_people;"`
}
