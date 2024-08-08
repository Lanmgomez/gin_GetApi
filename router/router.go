package router

import (
	"github.com/Lanmgomez/go-gin-api/internal/domain/user"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	router := gin.Default()

	router.GET("/users", user.GetUsers)
	router.GET("/users/:id", user.GetUserByID)

	router.POST("/users", user.PostUsers)

	router.PUT("/users/:id", user.UpdateUser)
	router.PUT("/users/updatestatus", user.UpdateUserStatus)
	router.PUT("/users/deletelogic/:id", user.DeleteLogicalUserByID)

	router.DELETE("/users/:id", user.DeleteUserByID)

	return router
}
