package controller

// メモリーのバックエンド処理のエントリーポイント
// バリデーション + service層へのパイプライン

import (
	"gin_backend/model"
	"gin_backend/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
全てのメモリーを取得します。

	Returns
	-------
	memories: model.MemoryResponse
		全メモリー
	status: 200
*/
func GetAllMemories(c *gin.Context) {
	memories := service.GetAllMemories()
	c.JSON(http.StatusOK, memories)
}

/*
メモリーを作成します。

	Parameters
	----------
	memory: model.Memory
		メモリーレコード

	Returns
	-------
	status: 201

	Exceptions
	----------
	status: 400
		リクエストエラー
*/
func CreateMemory(c *gin.Context) {
	var memory model.Memory
	err := c.BindJSON(&memory)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}

	service.CreateMemory(memory)
	c.Status(http.StatusCreated)
}

/*
メモリーレコードをUPDATEします。

	Parameters
	----------
	memoryId: int
		更新するメモリーのID
	memory: model.Memory
		メモリーレコード

	Returns
	-------
	status: 204

	Exceptions
	----------
	status: 400
		リクエストエラー
*/
func UpdateMemory(c *gin.Context) {
	memoryId, _ := strconv.Atoi(c.Param("id"))
	var memory model.Memory
	err := c.BindJSON(&memory)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	memory.Id = memoryId

	service.UpdateMemory(memory)
	c.Status(http.StatusNoContent)

}

/*
メモリーを削除します。

	Parameters
	----------
	memoryId: int
		削除するメモリーのID

	Returns
	-------
	status: 204
*/
func DeleteMemory(c *gin.Context) {
	memoryId, _ := strconv.Atoi(c.Param("id"))

	service.DeleteMemory(memoryId)
	c.Status(http.StatusNoContent)
}
