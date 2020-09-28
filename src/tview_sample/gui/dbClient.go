package gui

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var Client *DbClient

type DbClient struct {
	*dynamodb.DynamoDB
}

func NewDbClient() *DbClient {
    sess := session.Must(session.NewSessionWithOptions(session.Options{
		Profile: "default",
		SharedConfigState: session.SharedConfigEnable,
	}))

    svc := dynamodb.New(
		sess,
		aws.NewConfig().WithEndpoint("http://localhost:8000"),
	)

	Client = &DbClient{svc}

	return Client
}
