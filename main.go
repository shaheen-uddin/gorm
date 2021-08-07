package main

import (
	"github.com/shaheen-uddin/gorm/database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initDatabase() {
	var err error
	database.DbCon, err = gorm.Open(sqlite.Open("book.db"), &gorm.Config{})
}
