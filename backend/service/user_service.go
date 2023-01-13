package service

// ユーザーに関するビジネスロジック

import (
	"gin_backend/model"
	"gin_backend/repository"
)

func GetAllUsers() []model.UserRequestModel {
	users := repository.GetAllUsers()
	return users
}

func CreateUser(user model.UserRequestModel) {
	repository.CreateUser(user)
}
