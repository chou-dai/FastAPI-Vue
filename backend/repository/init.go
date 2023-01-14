package repository

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	db  *sql.DB
	err error
)

func dbConnect() {

	godotenv.Load(".env")

	USER := os.Getenv("MYSQL_USER")
	PASS := os.Getenv("MYSQL_PASSWORD")
	HOST := os.Getenv("DB_HOST")
	PORT := os.Getenv("DB_PORT")
	DB_NAME := os.Getenv("MYSQL_DATABASE")
	DSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", USER, PASS, HOST, PORT, DB_NAME)

	db, err = sql.Open("mysql", DSN)
	if err != nil {
		panic(err)
	}
}
