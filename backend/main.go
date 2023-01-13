package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var (
	db  *sql.DB
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

	db, err = sql.Open("mysql", CONNECT)
	if err != nil {
		panic(err)
	}
}

type User struct {
	Id       int
	Name     string
	Password string
}

func main() {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			Init()
			defer db.Close()
			results, err := db.Query("SELECT id, name, password FROM users")

			if err != nil {
				// simply print the error to the console
				fmt.Println("Err", err.Error())
			}
			users := []User{}
			for results.Next() {
				var user User
				results.Scan(&user.Id, &user.Name, &user.Password)
				users = append(users, user)
			}

			c.JSON(200, users)
		})

		api.POST("/", func(c *gin.Context) {
			Init()
			defer db.Close()

			var user User
			err := c.BindJSON(&user)
			if err != nil {
				c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
			}
			insert, err := db.Query(
				"INSERT INTO users(name,password) VALUES (?,?)",
				user.Name, user.Password)

			// if there is an error inserting, handle it
			if err != nil {
				panic(err.Error())
			}

			defer insert.Close()

			c.JSON(201, user)
		})
	}

	router.Run(":8080")
}
