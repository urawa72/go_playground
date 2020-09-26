package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
    "gopkg.in/AlecAivazis/survey.v1"
)

func main() {
    sess := session.Must(session.NewSession())

    svc := dynamodb.New(
        sess,
        aws.NewConfig().WithRegion("ap-northeast-1"),
    )

    input := &dynamodb.ListTablesInput{}
    result, err := svc.ListTables(input)
    if err != nil {
        panic(err)
    }
    fmt.Println(result)

    color := ""
    prompt := &survey.Select{
        Message: "Choose a color:",
        Options: []string{"red", "blue", "green"},
    }
    survey.AskOne(prompt, &color, nil)

    fmt.Println(color)

}
