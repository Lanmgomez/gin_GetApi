package main

import (
	"github.com/gin-gonic/gin"
)

func userData(c *gin.Context) {
	users := []map[string]interface{}{
		{
			"name":     "Islan",
			"email":    "islan_gomes@hotmail.com",
			"creatAt":  "10/08/2024",
			"updateAt": "10/08/2024",
		},
		{
			"name":     "Vinicius",
			"email":    "vinicius_mocci@hotmail.com",
			"creatAt":  "11/08/2024",
			"updateAt": "11/08/2024",
		},
	}

	for index, user := range users {
		user["id"] = index + 1
	}

	c.JSON(200, users)
}

func main() {
	r := gin.Default()

	r.GET("/users", userData)

	r.Run(":5000")
}
