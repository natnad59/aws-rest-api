package resolvers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/natnad59/aws-rest-api/models"
)

func GetUsers(c *gin.Context) {
	users := []models.User{
		{
			Id:         1,
			Name:       "Sheldon",
			Catchprase: "Bazinga",
			Retired:    false,
		},
		{
			Id:         2,
			Name:       "Bob",
			Catchprase: "zooweemama",
			Retired:    true,
		},
	}

	c.IndentedJSON(http.StatusOK, users)
}
