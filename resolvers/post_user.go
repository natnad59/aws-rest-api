package resolvers

import (
	"context"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/natnad59/aws-rest-api/client"
	"github.com/natnad59/aws-rest-api/constants"
	"github.com/natnad59/aws-rest-api/models"
)

func PostUser(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	client, err := client.NewDynamoDbClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	// Generate a new uuid for the user
	user.Id = uuid.New().String()

	userIn, err := attributevalue.MarshalMap(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	_, err = client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName:           aws.String(constants.TABLE_NAME),
		Item:                userIn,
		ConditionExpression: aws.String("attribute_not_exists(id)"),
		ReturnValues:        "ALL_NEW",
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}
