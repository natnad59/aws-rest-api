package client

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/natnad59/aws-rest-api/constants"
	"github.com/natnad59/aws-rest-api/models"
)

func DeleteUser(id string) (*models.User, error) {
	client, err := NewDynamoDbClient()
	if err != nil {
		return nil, err
	}

	out, err := client.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: aws.String(constants.TABLE_NAME),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
		ReturnValues: "ALL_OLD",
	})
	if err != nil {
		return nil, err
	}

	var user models.User
	if err = attributevalue.UnmarshalMap(out.Attributes, &user); err != nil {
		return nil, err
	}

	return &user, nil
}
