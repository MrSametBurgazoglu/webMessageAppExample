package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	groupModels "webMessageApp/groups/models"
	messageModels "webMessageApp/messages/models"
	peopleModels "webMessageApp/people/models"
)

var DB *gorm.DB

func ConnectDatabase() {

	dsn := "host=localhost user=message_app password=76T4376T43 dbname=web_message_app sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	migrateErr := db.AutoMigrate(&groupModels.Group{}, &messageModels.Message{}, &messageModels.Chat{}, &peopleModels.Person{})
	if migrateErr != nil {
		return
	}

	DB = db
}
