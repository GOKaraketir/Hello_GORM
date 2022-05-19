package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

type HelloModel struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Data      string
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&HelloModel{})
	if err != nil {
		return
	}

	// Create
	db.Create(&HelloModel{
		Data: "Hello World",
	})

	// Read
	var helloModel HelloModel
	db.First(&helloModel, 1) // find helloModel with integer primary key

	// Update
	db.Model(&helloModel).Update("Data", "Hello World Updated")

	// Delete - delete helloModel
	//db.Delete(&helloModel, 1)
}
