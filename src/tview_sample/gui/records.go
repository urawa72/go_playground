package gui

import (
    "encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/rivo/tview"
)

type Records struct {
	*tview.Table
}

func NewRecords() *Records {
	r := &Records{
		Table: tview.NewTable().Select(0, 0).SetSelectable(true, false),
	}
    r.SetBorder(true).SetTitleAlign(tview.AlignLeft)
	return r
}

func (r *Records) UpdateView(g *Gui) {
	tableName := g.TableList.Selected()
	r.scanItems(tableName)
}

func (r *Records) scanItems(name string) {

	params := &dynamodb.ScanInput{
		TableName: aws.String(name),
	}

	result, err := Client.Scan(params)
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

	t := r.Clear()
	for i, item := range items {
		c := 0
		for _, v := range item.Data {
			j, err := json.Marshal(v)
			if err != nil {
				panic(err)
			}
			t.SetCell(i+1, c, tview.NewTableCell(string(j)))
			c++
		}
	}
}

type Item struct {
	Data map[string]interface{}
}
