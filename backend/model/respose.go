package model

import "time"

type MemoryWithUserName struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	IsPublic    bool      `json:"isPublic"`
	UserName    int       `json:"userName"`
}

type PublicMemoryResponse struct {
	Memories []MemoryWithUserName `json:"memories"`
}

type MyMemoryResponse struct {
	Memories []Memory `json:"memories"`
}

type LoginUserResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
