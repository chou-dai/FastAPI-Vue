package service

// メモリーに関するビジネスロジック

import (
	"gin_backend/model"
	"gin_backend/repository"
)

/*
全てのメモリーを取得します。

	Returns
	-------
	memories: []model.Memory
		全メモリー
*/
func GetAllMemories() model.MemoryResponse {
	memories := repository.GetAllMemories()
	return model.MemoryResponse{Memories: memories}
}

/*
メモリーを作成します。

	Parameters
	----------
	memory: model.Memory
		メモリーレコード
*/
func CreateMemory(memory model.Memory) {
	// memory.CreatedAt = time.Now()
	repository.CreateMemory(memory)
}

/*
メモリーを更新します。

	Parameters
	----------
	memory: model.Memory
		メモリーレコード
*/
func UpdateMemory(memory model.Memory) {
	repository.UpdateMemory(memory)
}

/*
メモリーを削除します。

	Parameters
	----------
	memoryId: int
		削除するメモリーのID
*/
func DeleteMemory(memoryId int) {
	repository.DeleteMemory(memoryId)
}
