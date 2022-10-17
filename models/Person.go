package models

import "gorm.io/gorm"

type Person struct {
	gorm.Model
	Nickname string  `json:"nickname"`
	Chats    []*Chat `gorm:"many2many:people_chats;"`
}
