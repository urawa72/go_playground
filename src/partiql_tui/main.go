package main

import (
	// "os"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"

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
	// os.Exit(start())
	res, _ := RunCmd()
	fmt.Println(res)
}

type Item struct {
	Item map[string]interface{}
}

type Items struct {
	Items []map[string]interface{}
}

func RunCmd() (string, error) {
	sql := "select * from books"
	buf := bytes.Buffer{}
	cmd := exec.Command("aws", "dynamodb", "execute-statement", "--statement", sql)
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	if err := cmd.Run(); err != nil {
		return "", errors.New(buf.String())
	}

	jsonStr := []byte(buf.String())
	var items Items
	if err := json.Unmarshal(jsonStr, &items); err != nil {
		fmt.Println("Errrr!")
	}
	fmt.Println(items)

	return buf.String(), nil
}
