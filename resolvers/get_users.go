package resolvers

import (
	"context"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gin-gonic/gin"
	"github.com/natnad59/aws-rest-api/client"
	"github.com/natnad59/aws-rest-api/constants"
	"github.com/natnad59/aws-rest-api/models"
)

func GetUsers(c *gin.Context) {
	client, err := client.NewDynamoDbClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"error": err.Error()}})
		return
	}

	var usersOut []models.User
	paginator := dynamodb.NewScanPaginator(client, &dynamodb.ScanInput{
		TableName: aws.String(constants.TABLE_NAME),
	})

	for paginator.HasMorePages() {
		out, err := paginator.NextPage(context.TODO())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var paginatedUsers []models.User
		if err = attributevalue.UnmarshalListOfMaps(out.Items, &paginatedUsers); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		usersOut = append(usersOut, paginatedUsers...)
	}

	c.JSON(http.StatusOK, usersOut)
}
