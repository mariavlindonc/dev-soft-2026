package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB // It's defined outside any function to be able to use it everywhere

func StartDB() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	var err error

	// dsn := "user:#MySQLpassword512@tcp(localhost:3306)/db_vinilos?charset=utf8mb4&parseTime=True&loc=Local"
	// This last command can be parametrized

	user := "root"
	pass := "#MySQLpassword512"
	host := "localhost"
	port := "3306"
	name := "db_vinilos"

	dsn := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + name + "?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database connected")
}
