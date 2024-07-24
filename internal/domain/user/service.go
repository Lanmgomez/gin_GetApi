// escrever as regras de negócio
package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Lanmgomez/go-gin-api/User"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	// rows, err := db.Query("SELECT * FROM users")
	rows, err := db.Query("SELECT ID, Name, Email, CreateAt, UpdatedAt FROM users")

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
			&user.UpdatedAt,
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

func GetUserByID(context *gin.Context) {
	id := context.Param("id") // format string

	parsedIDtoInt := parseParamIDtoInt(id)

	// row := db.QueryRow("SELECT * FROM users WHERE id = ?", parsedIDtoInt)
	row := db.QueryRow("SELECT ID, Name, Email, CreateAt, UpdatedAt FROM users WHERE id = ?", parsedIDtoInt)
	var user User.USERS

	if err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.CreateAt,
		&user.UpdatedAt,
	); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, true)
}

func PostUsers(c *gin.Context) {
	var newUser User.USERS

	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err := db.Exec(
		"INSERT INTO users (name, email, createAt, updatedAt) VALUES (?, ?, ?, ?)",
		newUser.Name,
		newUser.Email,
		newUser.CreateAt,
		newUser.UpdatedAt,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, true)
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id") // format string
	parsedIDtoInt := parseParamIDtoInt(id)

	var updateUserByID User.USERS

	if err := c.BindJSON(&updateUserByID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err := db.Exec(
		"UPDATE users SET name = ?, email = ?, createAt = ?, updatedAt = ? WHERE id = ?",
		updateUserByID.Name,
		updateUserByID.Email,
		updateUserByID.CreateAt,
		updateUserByID.UpdatedAt,
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

func UpdateUserStatus(c *gin.Context) {
	var input User.USER_STATUS_UPDATE_INPUT

	if (input.CurrentUserType) != "ADMIN" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Access denied, you must be an admin",
		})
		return
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err := db.Exec("UPDATE users SET Status = ? WHERE id = ?",
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

func parseParamIDtoInt(id string) int {
	parsedID, err := strconv.ParseInt(id, 10, 64) // 10 base, 64 bits

	if err != nil {
		fmt.Println(err)
		return 0
	}

	return int(parsedID)
}
