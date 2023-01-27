package repository

import (
	"fmt"
	"gin_backend/model"

	_ "github.com/go-sql-driver/mysql"
)

func GetPublicMemories() []model.MemoryWithUserName {
	dbConnect()
	defer db.Close()

	rows, err := db.Query(
		`SELECT memories.id, title, description, created_at, is_public, name
			FROM memories
			LEFT JOIN users
				ON memories.user_id = users.id
			WHERE is_public = true
			ORDER BY memories.created_at DESC
			LIMIT 100`)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	memories := []model.MemoryWithUserName{}
	for rows.Next() {
		var memory model.MemoryWithUserName
		rows.Scan(&memory.Id, &memory.Title, &memory.Description, &memory.CreatedAt, &memory.IsPublic, &memory.UserName)
		memories = append(memories, memory)
	}

	return memories
}

func GetMyMemories(userId int) []model.Memory {
	dbConnect()
	defer db.Close()

	query := fmt.Sprintf(
		`SELECT id, title, description, created_at, is_public
			FROM memories
			WHERE user_id = %d
			ORDER BY memories.created_at DESC`, userId)
	rows, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	memories := []model.Memory{}
	for rows.Next() {
		var memory model.Memory
		rows.Scan(&memory.Id, &memory.Title, &memory.Description, &memory.CreatedAt, &memory.IsPublic)
		memories = append(memories, memory)
	}

	return memories
}

func CreateMemory(memory model.Memory) {
	dbConnect()
	defer db.Close()
	var layout = "2006-01-02 15:04:05"
	dateTime := memory.CreatedAt.Format(layout)

	insert, err := db.Query(
		"INSERT INTO memories(title, description, is_public, user_id, created_at) VALUES (?, ?, ?, ?, ?)",
		memory.Title, memory.Description, memory.IsPublic, memory.UserId, dateTime)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}

func UpdateMemory(memory model.Memory) {
	dbConnect()
	defer db.Close()

	update, err := db.Query(
		"UPDATE memories SET title = ?, description = ?, is_public = ? WHERE id = ?",
		memory.Title, memory.Description, memory.IsPublic, memory.Id)
	if err != nil {
		panic(err.Error())
	}
	defer update.Close()
}

func DeleteMemory(memoryId int) {
	dbConnect()
	defer db.Close()

	delete, err := db.Query("DELETE FROM memories WHERE id = ?", memoryId)
	if err != nil {
		panic(err.Error())
	}
	defer delete.Close()
}
