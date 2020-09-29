package main

import (
	"fmt"
    "encoding/json"
	"reflect"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func main() {
    sess := session.Must(session.NewSessionWithOptions(session.Options{
		Profile: "default",
		SharedConfigState: session.SharedConfigEnable,
	}))

	params := &dynamodb.ScanInput{
		TableName: aws.String("hoge"),
	}

    svc := dynamodb.New(
		sess,
		aws.NewConfig().WithEndpoint("http://localhost:8000"),
	)

	tableName := "hoge"
	input := &dynamodb.DescribeTableInput{ TableName: &tableName }
    result2, err := svc.DescribeTable(input)
    if err != nil {
        panic(err)
    }
	fmt.Println(result2)

	result, err := svc.Scan(params)
	if err != nil {
		panic(err)
	}

    items := []Item{}

	for _, item := range result.Items {
		var i Item
		err = dynamodbattribute.UnmarshalMap(item, &i.Data)
		if err != nil {
			panic(err)
		}
		items = append(items, i)
	}

	for _, item := range items {
		for _, v := range item.Data {
			j, err := json.Marshal(v)
			if err != nil {
				panic(err)
			}
			fmt.Println(reflect.TypeOf(v))
			fmt.Println(string(j))
			// fmt.Println(v)
		}
	}
}

type Item struct {
	Data map[string]interface{}
}
