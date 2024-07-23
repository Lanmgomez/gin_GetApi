package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Lanmgomez/go-gin-api/User"
	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

var database *sql.DB
var err error

func main() {
	initDB()
	router := gin.Default()

	router.GET("/users", getUsers)
	router.GET("/users/:id", getUserByID)
	router.POST("/users", postUsers)
	router.PUT("/users/:id", updateUser)
	router.PUT("/users/updatestatus", updateUserStatus)

	router.Run(":5000")
}

func initDB() {
	dsn := "root:levelphone@tcp(127.0.0.1:3306)/levelphone"

	database, err = sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal(err)
	}

	err = database.Ping()

	if err != nil {
		log.Fatal(err)
	}
}

func getUsers(c *gin.Context) {
	rows, err := database.Query("SELECT * FROM users")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	defer rows.Close()

	var users []User.USERS

	for rows.Next() {
		var user User.USERS

		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.CreateAt,
			&user.UpdateAt,
		); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		users = append(users, user)
	}

	c.JSON(http.StatusOK, true)
}

func parseParamIDtoInt(id string) int {
	parsedID, err := strconv.ParseInt(id, 10, 64) // 10 base, 64 bits

	if err != nil {
		fmt.Println(err)
		return 0
	}

	return int(parsedID)
}

func getUserByID(context *gin.Context) {
	id := context.Param("id") // format string

	parsedIDtoInt := parseParamIDtoInt(id)

	row := database.QueryRow("SELECT * FROM users WHERE id = ?", parsedIDtoInt)
	var user User.USERS

	if err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.CreateAt,
		&user.UpdateAt,
	); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, true)
}

func postUsers(c *gin.Context) {
	var newUser User.USERS

	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	fmt.Printf("O usuário enviado é %s, o email é %s", newUser.Name, newUser.Email)

	_, err := database.Exec(
		"INSERT INTO users (name, email, createAt, updatedAt) VALUES (?, ?, ?, ?)",
		newUser.Name,
		newUser.Email,
		newUser.CreateAt,
		newUser.UpdateAt,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, true)
}

func updateUser(c *gin.Context) {
	id := c.Param("id") // format string
	parsedIDtoInt := parseParamIDtoInt(id)

	var updateUserByID User.USERS

	if err := c.BindJSON(&updateUserByID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err := database.Exec(
		"UPDATE users SET name = ?, email = ?, createAt = ?, updatedAt = ? WHERE id = ?",
		updateUserByID.Name,
		updateUserByID.Email,
		updateUserByID.CreateAt,
		updateUserByID.UpdateAt,
		parsedIDtoInt,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, true)
}

func updateUserStatus(c *gin.Context) {
	var input User.USER_STATUS_UPDATE_INPUT

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if input.CurrentUserType != "ADMIN" {
		// rever o status para um mais correto
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Usuário não permitido",
		})
		return
	}

	_, err := database.Exec("UPDATE users SET Status = ? WHERE id = ?",
		input.Status,
		input.UserID,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, true)
}

// Exemplo com dados mockados
/*
	func userData(c *gin.Context) {
		finalUsers := User.UsersData()
		c.JSON(200, finalUsers)
	}
*/
