package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"webMessageApp/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	productionServer := os.Getenv("ON_PRODUCTION")
	var dsn string
	if productionServer != "" {
		databaseHost := os.Getenv("DATABASE_URL")
		databaseUser := os.Getenv("DATABASE_USER")
		databasePassword := os.Getenv("DATABASE_PASSWORD")
		databaseName := os.Getenv("DATABASE_NAME")
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
			databaseHost,
			databaseUser,
			databasePassword,
			databaseName)
		dsn = "host=" + databaseHost + " user=" + databaseUser + " password=" + databasePassword + " dbname=" + databaseName
	} else {
		dsn = "host=localhost user=message_app password=76T4376T43 dbname=web_message_app sslmode=disable"
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	migrateErr := db.AutoMigrate(&models.Message{}, &models.Chat{}, &models.Person{})
	if migrateErr != nil {
		return
	}

	DB = db
	var preCreatedPeople = []models.Person{
		{Nickname: "Person1"},
		{Nickname: "Person2"},
		{Nickname: "Person3"},
		{Nickname: "Person4"},
		{Nickname: "Person5"},
	}

	var preCreatedChats = []models.Chat{
		{ChatName: "Group1"},
		{ChatName: "Group2"},
		{},
		{},
		{},
	}

	DB.Exec("DELETE FROM people_chats")
	DB.Exec("DELETE FROM messages")
	DB.Exec("DELETE FROM people")
	DB.Exec("DELETE FROM chats")
	DB.Create(&preCreatedPeople)
	DB.Create(&preCreatedChats)

	var preCreatedMessages = []models.Message{
		{Context: "hello world", PersonID: preCreatedPeople[0].ID, ChatID: preCreatedChats[1].ID},
		{Context: "hello world to Person2", PersonID: preCreatedPeople[1].ID, ChatID: preCreatedChats[2].ID},
	}

	//DB.Create(&preCreatedMessages)
	DB.Model(&preCreatedPeople[0]).Association("Chats").Append(&preCreatedChats[2])
	DB.Model(&preCreatedPeople[0]).Association("Chats").Append(&preCreatedChats[1])
	DB.Model(&preCreatedPeople[1]).Association("Chats").Append(&preCreatedChats[2])
	DB.Model(&preCreatedPeople[2]).Association("Chats").Append(&preCreatedChats[1])
	DB.Model(&preCreatedPeople[3]).Association("Chats").Append(&preCreatedChats[1])
	DB.Model(&preCreatedChats[1]).Association("Messages").Append(&preCreatedMessages[0])
	DB.Model(&preCreatedChats[2]).Association("Messages").Append(&preCreatedMessages[1])
}
