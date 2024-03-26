package resolvers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/natnad59/aws-rest-api/client"
	"github.com/natnad59/aws-rest-api/models"
)

func GetUsers(c *gin.Context) {
	usersOut, err := client.GetUsers()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, []models.User{})
	}

	c.IndentedJSON(http.StatusOK, usersOut)
}
