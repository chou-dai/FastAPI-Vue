package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "SUCCESS!!",
			})
		})
	}

	router.Run(":8080")
}