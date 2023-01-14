package repository

// ユーザーに関するDB操作

import (
	"fmt"
	"gin_backend/model"

	_ "github.com/go-sql-driver/mysql"
)

func GetAllUsers() []model.User {
	dbConnect()
	defer db.Close()

	results, err := db.Query("SELECT id, name, password FROM users")

	if err != nil {
		fmt.Println("Err", err.Error())
	}
	users := []model.User{}
	for results.Next() {
		var user model.User
		results.Scan(&user.Id, &user.Name, &user.Password)
		users = append(users, user)
	}
	return users
}

func CreateUser(user model.User) {
	dbConnect()
	defer db.Close()

	insert, err := db.Query(
		"INSERT INTO users(name, password) VALUES (?, ?)",
		user.Name, user.Password)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}
