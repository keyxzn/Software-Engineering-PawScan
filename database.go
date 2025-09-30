package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
							os.Getenv("DB_HOST"), 
							os.Getenv("DB_PORT"), 
							os.Getenv("DB_USER"),
							os.Getenv("DB_PASSWORD"),
							os.Getenv("DB_NAME"))
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	DB = db
	log.Println("Database connected successfully")
}