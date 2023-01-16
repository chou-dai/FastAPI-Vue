package controller

import (
	"gin_backend/model"
	"gin_backend/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllMemories(c *gin.Context) {
	memories := service.GetAllMemories()
	c.JSON(http.StatusOK, memories)
}

func CreateMemory(c *gin.Context) {
	var memory model.Memory
	c.BindJSON(&memory)

	service.CreateMemory(memory)
	c.Status(http.StatusCreated)
}

func UpdateMemory(c *gin.Context) {
	memoryId, _ := strconv.Atoi(c.Param("id"))
	var memory model.Memory
	c.BindJSON(&memory)
	memory.Id = memoryId

	service.UpdateMemory(memory)
	c.Status(http.StatusNoContent)

}

func DeleteMemory(c *gin.Context) {
	memoryId, _ := strconv.Atoi(c.Param("id"))

	service.DeleteMemory(memoryId)
	c.Status(http.StatusNoContent)
}
