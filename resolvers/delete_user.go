package resolvers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/natnad59/aws-rest-api/client"
	"github.com/natnad59/aws-rest-api/models"
)

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.IndentedJSON(http.StatusBadRequest, models.User{})
		return
	}

	userOut, err := client.DeleteUser(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, models.User{})
	}

	c.IndentedJSON(http.StatusOK, userOut)
}
