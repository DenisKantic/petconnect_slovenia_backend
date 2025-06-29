package tests

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"slovenia_petconnect/database"
	"slovenia_petconnect/models"
)

func SetupTestDB() {

	testDB, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	database.DB = testDB

	// auto migrate models

	_ = database.DB.AutoMigrate(&models.User{})
}
