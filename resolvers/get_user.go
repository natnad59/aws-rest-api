package resolvers

import (
	"context"
	"errors"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/gin-gonic/gin"
	"github.com/natnad59/aws-rest-api/client"
	"github.com/natnad59/aws-rest-api/constants"
	"github.com/natnad59/aws-rest-api/models"
)

func GetUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, errors.New("please provide id path parameter").Error())
		return
	}

	client, err := client.NewDynamoDbClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	out, err := client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(constants.TABLE_NAME),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	var user models.User
	if err = attributevalue.UnmarshalMap(out.Item, &user); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}
