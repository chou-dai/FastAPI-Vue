package model

import "time"

type User struct {
	Id        int    `json:"id"`
	Name      string `json:"name" binding:"required"`
	Password  string `json:"password" binding:"required"`
	SessionId string `json:"sessionId"`
}

type Memory struct {
	Id          int       `json:"id"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	IsPublic    bool      `json:"isPublic"`
}
