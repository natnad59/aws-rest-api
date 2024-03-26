package resolvers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/natnad59/aws-rest-api/client"
	"github.com/natnad59/aws-rest-api/models"
)

func PostUser(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, models.User{})
	}

	userOut, err := client.PostUser(user)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, models.User{})
	}

	c.IndentedJSON(http.StatusOK, userOut)
}
