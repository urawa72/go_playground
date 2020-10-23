package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iot"
	"github.com/aws/aws-sdk-go/service/iotdataplane"
)

func main() {
    sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := iot.New(sess, aws.NewConfig().WithRegion("ap-northeast-1"))

	descResp, err := svc.DescribeEndpoint(&iot.DescribeEndpointInput{})
	if err != nil {
		panic(err)
	}

	dataSvc := iotdataplane.New(sess, &aws.Config{
		Endpoint: descResp.EndpointAddress,
	})

	result, err := svc.ListThings(&iot.ListThingsInput{})
	if err != nil {
		panic(err)
	}

	output, err := dataSvc.GetThingShadow(&iotdataplane.GetThingShadowInput{
		ThingName: result.Things[0].ThingName,
	})

	var buf bytes.Buffer
	if err := json.Indent(&buf, []byte(string(output.Payload)), "", "  "); err != nil {
		panic(err)
	}
	fmt.Println(buf.String())
}
