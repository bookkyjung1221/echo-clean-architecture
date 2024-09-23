package main

import (
	"fmt"

	"github.com/bookkyjung1221/echo-clean-architecture/db"
	"github.com/bookkyjung1221/echo-clean-architecture/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}
