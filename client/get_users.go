package client

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/natnad59/aws-rest-api/constants"
	"github.com/natnad59/aws-rest-api/models"
)

func GetUsers() (*[]models.User, error) {
	client, err := NewDynamoDbClient()
	if err != nil {
		return nil, err
	}

	var usersOut []models.User
	paginator := dynamodb.NewScanPaginator(client, &dynamodb.ScanInput{
		TableName: aws.String(constants.TABLE_NAME),
	})

	for paginator.HasMorePages() {
		out, err := paginator.NextPage(context.TODO())
		if err != nil {
			return nil, err
		}

		var paginatedUsers []models.User
		if err = attributevalue.UnmarshalListOfMaps(out.Items, &paginatedUsers); err != nil {
			return nil, err
		}

		usersOut = append(usersOut, paginatedUsers...)
	}

	return &usersOut, nil
}
