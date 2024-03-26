package client

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/natnad59/aws-rest-api/constants"
)

func NewDynamoDbClient() (*dynamodb.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), func(opts *config.LoadOptions) error {
		opts.Region = constants.AWS_REGION
		return nil
	})
	if err != nil {
		return nil, err
	}
	session := dynamodb.NewFromConfig(cfg)
	return session, nil
}
