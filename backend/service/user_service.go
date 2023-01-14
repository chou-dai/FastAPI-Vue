package service

// ユーザーに関するビジネスロジック

import (
	"gin_backend/model"
	"gin_backend/repository"
)

func GetAllUsers() []model.User {
	users := repository.GetAllUsers()
	return users
}

func CreateUser(user model.User) {
	repository.CreateUser(user)
}
