package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Mariadb() *sql.DB {

	// load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("can not load .env file", err)
	}

	db, err := sql.Open("mysql", os.Getenv("MARIADB"))

	if err != nil {
		panic(err.Error())
	} else {
		err = db.Ping()
		if err != nil {
			panic(err.Error())
		} else {
			fmt.Println("Mariadb Connected")
		}
	}
	return db
}
