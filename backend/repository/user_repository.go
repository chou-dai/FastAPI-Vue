package repository

import (
	"gin_backend/model"

	_ "github.com/go-sql-driver/mysql"
)

func IsUserExistBySessionId(sessionId string) bool {
	dbConnect()
	defer db.Close()

	var count int
	err := db.QueryRow("SELECT Count(*) FROM users WHERE session_id = ?", sessionId).Scan(&count)
	if err != nil {
		panic(err.Error())
	}
	return count != 0
}

func IsUserExistByName(name string) bool {
	dbConnect()
	defer db.Close()

	var count int
	err := db.QueryRow("SELECT Count(*) FROM users WHERE name = ?", name).Scan(&count)
	if err != nil {
		panic(err.Error())
	}
	return count != 0
}

func GetPasswordByName(name string) string {
	dbConnect()
	defer db.Close()

	var pw string
	_ = db.QueryRow("SELECT password FROM users WHERE name = ?", name).Scan(&pw)
	return pw
}

func GetUserBySessionId(sessionId string) model.User {
	dbConnect()
	defer db.Close()

	var user model.User
	err := db.QueryRow("SELECT id, name, session_id FROM users WHERE session_id = ?", sessionId).Scan(
		&user.Id, &user.Name, &user.SessionId)
	if err != nil {
		panic(err.Error())
	}
	return user
}

func CreateUser(user model.User) {
	dbConnect()
	defer db.Close()

	insert, err := db.Query(
		"INSERT INTO users(name, password, session_id) VALUES (?, ?, ?)",
		user.Name, user.Password, user.SessionId)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}

func UpdateSessionIdByName(sessionId string, name string) {
	dbConnect()
	defer db.Close()

	update, err := db.Query(
		"UPDATE users SET session_id = ? WHERE name = ?", sessionId, name)
	if err != nil {
		panic(err.Error())
	}
	defer update.Close()
}

func UpdateNameById(id int, name string) {
	dbConnect()
	defer db.Close()

	update, err := db.Query(
		"UPDATE users SET name = ? WHERE id = ?", name, id)
	if err != nil {
		panic(err.Error())
	}
	defer update.Close()
}
