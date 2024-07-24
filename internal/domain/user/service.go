// escrever as regras de negócio
package user

import (
	"fmt"

	"github.com/Lanmgomez/go-gin-api/internal/domain/user"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	db := user.InitDB()
	fmt.Print("teste")
}
