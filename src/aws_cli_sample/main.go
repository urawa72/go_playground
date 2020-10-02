package main

import (
	"fmt"
    "encoding/json"
	// "reflect"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/deckarep/golang-set"
)

type KeyDetail struct {
	KeyType string
	AttributeType *string
}

var header = []string{
	"HASH",
	"SORT",
}

func main() {
    sess := session.Must(session.NewSessionWithOptions(session.Options{
		Profile: "default",
		SharedConfigState: session.SharedConfigEnable,
	}))

    svc := dynamodb.New(
		sess,
		aws.NewConfig().WithEndpoint("http://localhost:8000"),
	)

	tableName := "hoge"
	input := &dynamodb.DescribeTableInput{ TableName: &tableName }
    tableDetail, _ := svc.DescribeTable(input)
	schema := tableDetail.Table.KeySchema
	var hashKey, sortKey string
	for _, s := range schema {
		if *s.KeyType == "HASH" {
			hashKey = *s.AttributeName
		}
		if *s.KeyType == "SORT" {
			sortKey = *s.AttributeName
		}
	}

	params := &dynamodb.ScanInput{
		TableName: aws.String("hoge"),
	}

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

	keys := mapset.NewSet()
	for _, item := range items {
		for k, _ := range item.Data {
			keys.Add(k)
		}
	}

	keyArray := keys.ToSlice()
	fmt.Println(keyArray)
	for _, item := range items {
		fmt.Println(hashKey, ":", item.Data[hashKey])
		if sortKey != "" {
			fmt.Println(item.Data[sortKey])
		}
		for _, key := range keyArray {
			if key == hashKey || key == sortKey {
				continue
			}
			j, err := json.Marshal(item.Data[key.(string)])
			if err != nil {
				panic(err)
			}
			fmt.Println(key, ":", string(j))
		}
	}

	// fmt.Println(keys)
}

type Item struct {
	Data map[string]interface{}
}
