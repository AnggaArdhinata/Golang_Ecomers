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

func Init()  {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	dbConfig := "user="+os.Getenv("DB_USER")+" password="+os.Getenv("DB_PASSWORD")+" dbname="+os.Getenv("DB_NAME")+" sslmode=disable"

	db, err = sql.Open("postgres", dbConfig)
	if err != nil {
		panic("error db connection !")
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("database connected . .")
	}
}

func CreateCon()  *sql.DB {
	return db
}