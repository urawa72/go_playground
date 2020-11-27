package tui

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var Client *dynamodb.DynamoDB

// type DynamodbClient struct {
// 	*dynamodb.DynamoDB
// }

func NewClient() *dynamodb.DynamoDB {
    sess := session.Must(session.NewSessionWithOptions(session.Options{
		// Profile: "kobayashi.yuki",
		SharedConfigState: session.SharedConfigEnable,
	}))

    svc := dynamodb.New(
		sess,
		aws.NewConfig(),
	)

	Client = svc
	// Client = &DynamodbClient{svc}

	return Client
}
