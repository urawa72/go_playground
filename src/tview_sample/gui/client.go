package gui

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var Client *DynamodbClient

type DynamodbClient struct {
	*dynamodb.DynamoDB
}

func NewClient() *DynamodbClient {
    sess := session.Must(session.NewSessionWithOptions(session.Options{
		Profile: "default",
		SharedConfigState: session.SharedConfigEnable,
	}))

    svc := dynamodb.New(
		sess,
		aws.NewConfig().WithEndpoint("http://localhost:8000"),
	)

	Client = &DynamodbClient{svc}

	return Client
}
