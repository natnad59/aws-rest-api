package resolvers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/natnad59/aws-rest-api/client"
)

func GetUsers(c *gin.Context) {
	usersOut, err := client.GetUsers()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	// js, err := json.Marshal(usersOut)
	// c.JSON(http.StatusOK, gin.H{"body": string(js)})
	c.JSON(http.StatusOK, usersOut)
}
