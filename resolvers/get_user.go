package resolvers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/natnad59/aws-rest-api/models"
)

func GetUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.IndentedJSON(http.StatusBadRequest, models.User{})
		return
	}

	user := models.User{
		Id:         1,
		Name:       "Sheldon",
		Catchprase: "Bazinga",
		Retired:    false,
	}

	c.IndentedJSON(http.StatusOK, user)
}
