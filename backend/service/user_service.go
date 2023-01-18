package service

import (
	"gin_backend/model"
	"gin_backend/repository"
)

func IsUserExistBySessionId(sessionId string) bool {
	return repository.IsUserExistBySessionId(sessionId)
}

func IsUserExistByName(name string) bool {
	return repository.IsUserExistByName(name)
}

func IsUserExist(user model.User) bool {
	return repository.IsUserExistByNameAndPwd(user.Name, user.Password)
}

func CreateUser(user model.User) model.User {
	repository.CreateUser(user)
	return repository.GetUserBySessionId(user.SessionId)
}

func UpdateSessionIdByName(sessionId string, name string) model.User {
	repository.UpdateSessionIdByName(sessionId, name)
	return repository.GetUserBySessionId(sessionId)
}
