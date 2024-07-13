package main

import (
	"github.com/gin-gonic/gin"
)

type USERS struct {
	ID       float64
	Name     string
	Email    string
	CreatAt  string
	UpdateAt string
}

func userData(c *gin.Context) {

	users := []USERS{
		{
			Name:     "Islan",
			Email:    "islan_gomes@hotmail.com",
			CreatAt:  "10/08/2024",
			UpdateAt: "10/08/2024",
		},
		{
			Name:     "Vinicius",
			Email:    "vinicius_mocci@hotmail.com",
			CreatAt:  "10/08/2024",
			UpdateAt: "10/08/2024",
		},
	}

	finalUsers := []USERS{}

	for index, user := range users {
		var currentUser USERS

		currentUser.ID = float64(index) + 1
		currentUser.Name = user.Name
		currentUser.Email = user.Email
		currentUser.CreatAt = user.CreatAt
		currentUser.UpdateAt = user.UpdateAt

		finalUsers = append(finalUsers, currentUser)
	}

	c.JSON(200, finalUsers)
}

func main() {
	r := gin.Default()

	r.GET("/users", userData)

	r.Run(":5000")
}
