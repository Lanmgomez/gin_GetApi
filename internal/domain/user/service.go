package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func parseParamIDtoInt(id string) int {
	parsedID, err := strconv.ParseInt(id, 10, 64) // 10 base, 64 bits

	if err != nil {
		fmt.Println(err)
		return 0
	}

	return int(parsedID)
}

func checkIfUserIsAdmin(c *gin.Context, input USER_STATUS_UPDATE_INPUT) bool {

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return false
	}

	if input.CurrentUserType != "ADMIN" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Access denied, you must be an admin",
		})
		return false
	}

	return true
}
