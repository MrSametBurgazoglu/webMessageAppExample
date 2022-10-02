package models

import "gorm.io/gorm"

type Chat struct {
	gorm.Model
	Messages        []Message
	ChatName        string
	ChatLastMessage string
	Person          []*Person `gorm:"many2many:people_chats;"`
}
