package configs

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

var err error

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	dbConfig := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	db, err = sql.Open("postgres", dbConfig)
	if err != nil {
		panic("error db connection !")
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		log.Println("database connected . .")
	}
}

func CreateCon() *sql.DB {
	return db
}
