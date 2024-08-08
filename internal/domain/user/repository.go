package user

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func InitDB(c *gin.Context) {
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

func GetUsers(c *gin.Context) {
	var showActiveUsers = "ATIVO"
	rows, err := db.Query("SELECT * FROM users WHERE Status = ?", showActiveUsers)
	// rows, err := db.Query("SELECT ID, Name, Email, CreateAt, UpdatedAt FROM users")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	defer rows.Close()

	var users []USERS

	for rows.Next() {
		var user USERS

		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Status,
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

	c.JSON(http.StatusOK, users)
}

func GetUserByID(context *gin.Context) {
	id := context.Param("id") // format string

	parsedIDtoInt := parseParamIDtoInt(id)

	row := db.QueryRow("SELECT * FROM users WHERE id = ?", parsedIDtoInt)
	// row := db.QueryRow("SELECT ID, Name, Email, CreateAt, UpdatedAt FROM users WHERE id = ?", parsedIDtoInt)
	var user USERS

	if err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Status,
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
	var newUser USERS

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

	var updateUserByID USERS

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
	var input USER_STATUS_UPDATE_INPUT

	permissionToUpdate := checkIfUserIsAdmin(c, input)

	if !permissionToUpdate {
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

// teste - exclusão física
func DeleteUserByID(c *gin.Context) {
	id := c.Param("id")
	parsedIDtoInt := parseParamIDtoInt(id)

	var deleteUserID USERS

	if err := c.BindJSON(&deleteUserID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err := db.Exec("DELETE FROM users WHERE id = ?", parsedIDtoInt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, true)
}

// teste Delete lóogico
func DeleteLogicalUserByID(c *gin.Context) {
	id := c.Param("id")
	parsedIDtoInt := parseParamIDtoInt(id)

	var logicDelete USERS
	var inactiveUser = "INATIVO"

	if err := c.BindJSON(&logicDelete); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err := db.Exec("UPDATE users SET Status = ? WHERE id = ?",
		inactiveUser,
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
