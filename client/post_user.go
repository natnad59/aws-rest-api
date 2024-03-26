package client

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/google/uuid"
	"github.com/natnad59/aws-rest-api/constants"
	"github.com/natnad59/aws-rest-api/models"
)

func PostUser(user models.User) (*models.User, error) {
	client, err := NewDynamoDbClient()
	if err != nil {
		return nil, err
	}

	// Generate a new uuid for the user
	user.Id = uuid.New().String()

	userIn, err := attributevalue.MarshalMap(user)

	_, err = client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName:           aws.String(constants.TABLE_NAME),
		Item:                userIn,
		ConditionExpression: aws.String("attribute_not_exists(id)"),
	})
	if err != nil {
		return nil, err
	}

	return &user, nil
}
