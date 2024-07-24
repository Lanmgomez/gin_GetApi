package router

import (
	"github.com/gin-gonic/gin"
)

func Routers(c *gin.Context) {
	router := gin.Default()

	router.GET("/users" /* função para rota */)
	router.GET("/users/:id" /* função para rota */)
	router.POST("/users" /* função para rota */)
	router.PUT("/users/:id" /* função para rota */)
	router.PUT("/users/updatestatus" /* função para rota */)

	router.Run(":5000")
}
