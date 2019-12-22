package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jpdvi/chat-service/models"
	"log"
	"os"
)

func main() {
	db, err := gorm.Open("postgres",
		fmt.Sprint(
			" host=", os.Getenv("DBHOST"),
			" port=", os.Getenv("DBPORT"),
			" user=", os.Getenv("DBUSER"),
			" dbname=", os.Getenv("DBNAME"),
			" password=", os.Getenv("DBPASSWORD"),
			" sslmode=disable"))
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Migrating to : ", os.Getenv("DBNAME"))
	db.AutoMigrate(&models.User{}, &models.Message{}, &models.Session{})
}
