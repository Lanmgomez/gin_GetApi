package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/Lanmgomez/go-gin-api/User"
	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func initDB() {
	dsn := "root:levelphone@tcp(127.0.0.1:3306)/levelphone"

	db, err = sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}
}

func userData(c *gin.Context) {

	finalUsers := User.UsersData()

	c.JSON(200, finalUsers)
}

func postUsers(c *gin.Context) {
	var newUser User.USERS
	finalUsers := User.UsersData()

	if err := c.BindJSON(&newUser); err != nil {
		fmt.Println("Error")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	fmt.Printf("O usuário enviado é %s, o email é %s", newUser.Name, newUser.Email)

	_, err := db.Exec(
		"INSERT INTO users (id, name, email, createAt, updatedAt) VALUES (?, ?, ?, ?, ?)",
		newUser.ID,
		newUser.Name,
		newUser.Email,
		newUser.CreatAt,
		newUser.UpdateAt,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	finalUsers = append(finalUsers, newUser)
	c.IndentedJSON(http.StatusCreated, finalUsers)
}

func main() {
	initDB()
	router := gin.Default()

	router.GET("/users", userData)
	router.POST("/users", postUsers)

	router.Run(":5000")
}
