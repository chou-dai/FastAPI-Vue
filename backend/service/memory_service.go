package service

import (
	"gin_backend/model"
	"gin_backend/repository"
	"time"
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
	now := time.Now()
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	nowJST := now.In(jst)
	memory.CreatedAt = nowJST
	repository.CreateMemory(memory)
}

func UpdateMemory(memory model.Memory) {
	repository.UpdateMemory(memory)
}

func DeleteMemory(memoryId int) {
	repository.DeleteMemory(memoryId)
}
