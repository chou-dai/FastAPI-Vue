package service

import (
	"gin_backend/model"
	"gin_backend/repository"
)

func GetPublicMemories() model.PublicMemoryResponse {
	memories := repository.GetPublicMemories()
	return model.PublicMemoryResponse{Memories: memories}
}

func GetMyMemories(userId int) model.MyMemoryResponse {
	memories := repository.GetMyMemories(userId)
	return model.MyMemoryResponse{Memories: memories}
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
