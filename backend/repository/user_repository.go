package repository

// ユーザーに関するDB操作

import (
	"database/sql"
	"fmt"
	"gin_backend/model"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var (
	db  *sql.DB
	err error
)

func Init() {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("読み込みに失敗しました: %v", err)
	}

	USER := os.Getenv("MYSQL_USER")
	PASS := os.Getenv("MYSQL_PASSWORD")
	PROTOCOL := "tcp(database:3306)"
	DB_NAME := os.Getenv("MYSQL_DATABASE")
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DB_NAME

	db, err = sql.Open("mysql", CONNECT)
	if err != nil {
		panic(err)
	}
}

func GetAllUsers() []model.UserRequestModel {
	Init()
	defer db.Close()
	results, err := db.Query("SELECT name, password FROM users")

	if err != nil {
		// simply print the error to the console
		fmt.Println("Err", err.Error())
	}
	users := []model.UserRequestModel{}
	for results.Next() {
		var user model.UserRequestModel
		results.Scan(&user.Name, &user.Password)
		users = append(users, user)
	}
	return users
}

func CreateUser(user model.UserRequestModel) {
	Init()
	defer db.Close()

	insert, err := db.Query(
		"INSERT INTO users(name,password) VALUES (?,?)",
		user.Name, user.Password)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}
