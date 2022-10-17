package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Context  string `json:"context"`
	PersonID uint
	Person   Person
	ChatID   uint
	Chat     Chat
}
