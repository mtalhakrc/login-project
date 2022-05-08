package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

const DNS = "postgres://talha:aaaa@localhost:3030/logindb?sslmode=disable"

var dbConn *gorm.DB
var err error

func DB() *gorm.DB {
	return dbConn
}

func InitializeDB() {
	dbConn, err = gorm.Open(postgres.Open(DNS))
	if err != nil {
		log.Fatalln("cant connect to database", err)
	}
	fmt.Println("connected to the database")
}
