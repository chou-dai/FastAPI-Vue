package controller

import (
	"gin_backend/model"
	"gin_backend/service"
	"gin_backend/session"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPublicMemories(c *gin.Context) {
	memories := service.GetPublicMemories()
	c.JSON(http.StatusOK, memories)
}

func GetMyMemories(c *gin.Context) {
	loginUser := session.GetLoginUserFromSession(c)

	memories := service.GetMyMemories(loginUser.UserId)
	c.JSON(http.StatusOK, memories)
}

func CreateMemory(c *gin.Context) {
	var memory model.Memory
	c.BindJSON(&memory)

	loginUser := session.GetLoginUserFromSession(c)
	memory.UserId = loginUser.UserId

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
