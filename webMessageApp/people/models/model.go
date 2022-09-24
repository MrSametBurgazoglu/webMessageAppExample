package models

type Person struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Nickname string `json:"nickname"`
}
