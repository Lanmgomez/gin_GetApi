package main

import (
	"fmt"
	"net/http"

	"github.com/Lanmgomez/go-gin-api/User"
	"github.com/gin-gonic/gin"
)

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

	finalUsers = append(finalUsers, newUser)
	c.IndentedJSON(http.StatusCreated, finalUsers)
}

func main() {
	router := gin.Default()

	router.GET("/users", userData)
	router.POST("/users", postUsers)

	router.Run(":5000")
}
