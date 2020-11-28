package main

import (
	"os"
	// "bytes"
	// "encoding/json"
	// "errors"
	// "fmt"
	// "os/exec"

	// "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	// "github.com/aws/aws-sdk-go/service/dynamodb"
	// "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/urawa72/partiql_tui/tui"
)

func start() int {
	if err := tui.New().Run(); err != nil {
		return 1
	}

	return 0
}

func main() {
	os.Exit(start())
	// res, _ := RunCmd()
}

// type Item struct {
// 	Item map[string]interface{}
// }
//
// type Items struct {
// 	Items []map[string]interface{} `json:"Items"`
// }
//
// func RunCmd() (string, error) {
// 	sql := "select * from books"
// 	buf := bytes.Buffer{}
// 	cmd := exec.Command("aws", "dynamodb", "execute-statement", "--statement", sql)
// 	cmd.Stdout = &buf
// 	cmd.Stderr = &buf
// 	if err := cmd.Run(); err != nil {
// 		return "", errors.New(buf.String())
// 	}
//
// 	jsonStr := []byte(buf.String())
// 	var items Items
// 	if err := json.Unmarshal(jsonStr, &items); err != nil {
// 		fmt.Println("Errrr!")
// 	}
//
// 	var list []Item
// 	for _, m := range items.Items {
// 		values := make(map[string]*dynamodb.AttributeValue)
// 		for k, mp := range m {
// 			var ddbAttr dynamodb.AttributeValue
// 			bytes, _ := json.Marshal(mp)
// 			json.Unmarshal(bytes, &ddbAttr)
// 			values[k] = &ddbAttr
// 		}
// 		var item Item
// 		dynamodbattribute.UnmarshalMap(values, &item.Item)
// 		list = append(list, item)
// 	}
//
// 	for _, m := range list {
// 		for k, v := range m.Item {
// 			fmt.Println(k, v)
// 		}
// 	}
//
// 	return buf.String(), nil
// }
