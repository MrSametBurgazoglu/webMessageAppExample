package models

import "gorm.io/gorm"

type PersonMessage struct {
	gorm.Model
	Context string `json:"nickname"`
}

type GroupMessage struct {
	gorm.Model
	Context string `json:"nickname"`
}
