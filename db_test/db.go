package db

import (
	"gorm.io/driver/sqlite"
  "gorm.io/gorm"
	"log"
	"github.com/suisuss/habitum_graphQL/models"
)

func OpenDB(database string) *gorm.DB {

	db, err := gorm.Open(sqlite.Open(database), &gorm.Config{})
	if err != nil {
		log.Fatalf("%s", err)
	}

	db.AutoMigrate(&models.Habit{}, &models.HabitLog{}, &models.User{})

	return db
}