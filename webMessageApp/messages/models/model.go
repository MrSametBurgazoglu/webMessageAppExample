package models

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	Context  string `json:"nickname"`
	PersonID uint
	ChatID   uint
}

type Chat struct {
	gorm.Model
	Messages []Message
	GroupID  uint
}
