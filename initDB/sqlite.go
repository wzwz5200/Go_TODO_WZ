package initdb

import (
	"GO_CLI/model"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var ReturnDB *gorm.DB

type User struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

func InitDB() *gorm.DB {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&User{}, &model.Task{})

	ReturnDB = db
	return ReturnDB

}
