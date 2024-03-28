package resolvers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/natnad59/aws-rest-api/client"
)

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.String(http.StatusBadRequest, errors.New("please provide id path parameter").Error())
		return
	}

	userOut, err := client.DeleteUser(id)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, userOut)
}
