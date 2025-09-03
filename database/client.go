package database

import (
	"log"
	"strings"
	"test-jwt-go/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB

var dbError error

func Connect(connectionString string, configs ...string) {
	dsn := strings.NewReplacer(configs...).Replace(connectionString)
	Instance, dbError = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to DB")
	}

	log.Println("Connected to DB")
}

func Migrate() {
	Instance.AutoMigrate(&model.User{})
	log.Println("Database Migration Completed")
}
