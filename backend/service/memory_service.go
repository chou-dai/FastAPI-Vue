package service

import (
	"gin_backend/model"
	"gin_backend/repository"
)

func GetAllMemories() model.MemoryResponse {
	memories := repository.GetAllMemories()
	return model.MemoryResponse{Memories: memories}
}

func CreateMemory(memory model.Memory) {
	repository.CreateMemory(memory)
}

func UpdateMemory(memory model.Memory) {
	repository.UpdateMemory(memory)
}

func DeleteMemory(memoryId int) {
	repository.DeleteMemory(memoryId)
}
