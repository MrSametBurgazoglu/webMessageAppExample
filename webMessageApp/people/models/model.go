package models

import "webMessageApp/messages/models"

type Person struct {
	ID       uint           `json:"id" gorm:"primary_key"`
	Nickname string         `json:"nickname"`
	Chats    []*models.Chat `gorm:"many2many:people_chats;"`
}
