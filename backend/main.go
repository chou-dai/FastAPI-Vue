package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var (
	db  *gorm.DB
	err error
)

func Init() {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("読み込みに失敗しました: %v", err)
	}

	USER := os.Getenv("MYSQL_USER")
	PASS := os.Getenv("MYSQL_PASSWORD")
	PROTOCOL := "tcp(database:3306)"
	DB_NAME := os.Getenv("MYSQL_DATABASE")
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DB_NAME

	db, err = gorm.Open("mysql", CONNECT)
	if err != nil {
		panic(err)
	}
}

type User struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func main() {
	router := gin.Default()
	Init()

	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			var u []User

			db.Find(&u)

			c.JSON(200, u)
		})

		api.POST("/", func(c *gin.Context) {
			user := User{}
			err := c.BindJSON(&user)
			if err != nil {
				c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
			}
			db.Create(&user)

			c.JSON(201, user)
		})
	}

	router.Run(":8080")
}
