package main

import (
	"fmt"

	"github.com/shinsx/golang-blog/db"
	"github.com/shinsx/golang-blog/model"
)

func main() {
	dbConn := db.NewDB()
	defer db.CloseDB(dbConn)
	if err := dbConn.AutoMigrate(&model.User{}); err != nil {
		fmt.Println("Failed to migrate database:", err)
		return
	}
	fmt.Println("Successfully Migrated")
}
