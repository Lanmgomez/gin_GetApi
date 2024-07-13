package main

import (
	"github.com/Lanmgomez/go-gin-api/User"
	"github.com/gin-gonic/gin"
)

func userData(c *gin.Context) {

	finalUsers := User.UsersData()

	c.JSON(200, finalUsers)
}

func main() {
	r := gin.Default()

	r.GET("/users", userData)

	r.Run(":5000")
}
