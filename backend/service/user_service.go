package service

import (
	"gin_backend/model"
	"gin_backend/repository"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func IsUserExistBySessionId(sessionId string) bool {
	return repository.IsUserExistBySessionId(sessionId)
}

func IsUserExistByName(name string) bool {
	return repository.IsUserExistByName(name)
}

func IsUserExistForLogin(user model.User) bool {
	pw := repository.GetPasswordByName(user.Name)
	if pw == "" {
		return false
	}
	// DB上のパスワードハッシュと入力パスワードを比較
	err := bcrypt.CompareHashAndPassword([]byte(pw), []byte(user.Password))
	if err != nil {
		return false
	}
	return true
}

func CreateUser(user model.User) model.User {
	// パスワードのハッシュ化
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err.Error())
	}
	user.Password = string(hash)
	repository.CreateUser(user)
	return repository.GetUserBySessionId(user.SessionId)
}

func UpdateSessionIdByName(sessionId string, name string) model.User {
	repository.UpdateSessionIdByName(sessionId, name)
	return repository.GetUserBySessionId(sessionId)
}

func UpdateUserName(loginUser model.SessionInfo, updateUser model.User) bool {
	log.Print(updateUser.Password)
	pw := repository.GetPasswordByName(loginUser.Name)
	// DB上のパスワードハッシュと入力パスワードを比較
	err := bcrypt.CompareHashAndPassword([]byte(pw), []byte(updateUser.Password))
	if err != nil {
		return false
	}
	repository.UpdateNameById(loginUser.UserId, updateUser.Name)
	return true
}
