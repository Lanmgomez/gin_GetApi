package main

import (
	"github.com/Lanmgomez/go-gin-api/internal/domain/user"
	"github.com/Lanmgomez/go-gin-api/router"
	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	c := &gin.Context{}
	user.InitDB(c)

	r := router.Routers()
	r.Run(":5000")
}
