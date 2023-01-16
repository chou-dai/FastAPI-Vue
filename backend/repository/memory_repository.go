package repository

import (
	"gin_backend/model"

	_ "github.com/go-sql-driver/mysql"
)

func GetAllMemories() []model.Memory {
	dbConnect()
	defer db.Close()

	rows, err := db.Query("SELECT id, title, description, created_at, is_public FROM memories")
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

	insert, err := db.Query(
		"INSERT INTO memories(title, description, is_public) VALUES (?, ?, ?)",
		memory.Title, memory.Description, memory.IsPublic)
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
