package gui

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/rivo/tview"
)

type Text1 struct {
	*tview.TextView
}

func NewText1() *Text1 {
	tx := &Text1{
		TextView: tview.NewTextView().SetDynamicColors(true),
	}
    tx.SetTitleAlign(tview.AlignLeft).SetTitle("Main 1").SetBorder(true)
	tx.SetWrap(false)

	return tx
}

func (tx *Text1) UpdateView(g *Gui) {
	tableName := g.TableList.Selected()
	tx.scanItems(tableName)
	// tx.SetText()
}

func (tx *Text1) scanItems(name string) {

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

	for _, item := range items {
		for k, v := range item.Data {
			fmt.Println(k, v)
			// tx.SetText(k)
            // switch v := v.(type) {
			// case string:
			// 	tx.SetText(v)
			// case []string:
			// 	tx.SetText(v[0])
			// }
		}
	}


}

type Item struct {
	Data map[string]interface{}
}
