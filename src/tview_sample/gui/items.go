package gui

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/deckarep/golang-set"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type Items struct {
	*tview.Table
	ItemArray	[]Item
}

type Item struct {
	Data map[string]interface{}
}

var header = []string{
	"HASH",
	"SORT",
}

func NewItems() *Items {
	i := &Items{
		Table: tview.NewTable().Select(0, 0).SetSelectable(true, true),
	}
    i.SetBorder(true).SetTitle("Items").SetTitleAlign(tview.AlignLeft)
	return i
}

func (i *Items) Selecting() *Item {
	if len(i.ItemArray) == 0 {
		return nil
	}
	row, _ := i.GetSelection()
	if row < 0 {
		return nil
	}
	fmt.Println(row)
	return &i.ItemArray[row-1]
}

func (i *Items) UpdateView(g *Gui) {
	tableName := g.Tables.Selected()
	i.scanItems(tableName)
	g.nextPanel()
}

func (items *Items) scanItems(name string) {
	input := &dynamodb.DescribeTableInput{
		TableName: aws.String(name),
	}
    tableDetail, _ := Client.DescribeTable(input)
	schema := tableDetail.Table.KeySchema
	var hashKey string
	var sortKey string
	for _, s := range schema {
		if *s.KeyType == "HASH" {
			hashKey = *s.AttributeName
		}
		if *s.KeyType == "SORT" {
			sortKey = *s.AttributeName
		}
	}

	params := &dynamodb.ScanInput{
		TableName: aws.String(name),
		Limit: aws.Int64(100),
	}

	result, err := Client.Scan(params)
	if err != nil {
		panic(err)
	}

    items.ItemArray = []Item{}

	for _, item := range result.Items {
		var tmp Item
		err = dynamodbattribute.UnmarshalMap(item, &tmp.Data)
		if err != nil {
			panic(err)
		}
		items.ItemArray = append(items.ItemArray, tmp)
	}

	keys := mapset.NewSet()
	for _, item := range items.ItemArray {
		for k, _ := range item.Data {
			keys.Add(k)
		}
	}
	keyArray := keys.ToSlice()
	sort.Slice(keyArray, func(i, j int) bool { return keyArray[i].(string) <  keyArray[j].(string) })

	t := items.Clear()

	c := 0
	t.SetCell(0, c, &tview.TableCell{
		Text:				hashKey,
		NotSelectable:		true,
		Align:				tview.AlignLeft,
		Color:				tcell.ColorYellow,
		BackgroundColor:	tcell.ColorDefault,
	})
	if sortKey != "" {
		c++
		t.SetCell(0, c, &tview.TableCell{
			Text:				sortKey,
			NotSelectable:		true,
			Align:				tview.AlignLeft,
			Color:				tcell.ColorYellow,
			BackgroundColor:	tcell.ColorDefault,
		})
	}
	for i, h := range keyArray {
		t.SetCell(0, c+i, &tview.TableCell{
			Text:				h.(string),
			NotSelectable:		true,
			Align:				tview.AlignLeft,
			Color:				tcell.ColorYellow,
			BackgroundColor:	tcell.ColorDefault,
		})
	}

	for i, item := range items.ItemArray {
		c := 0
		t.SetCell(i+1, c, tview.NewTableCell(item.Data[hashKey].(string)).SetMaxWidth(20))
		if sortKey != "" {
			c++
			t.SetCell(i+1, c, tview.NewTableCell(item.Data[sortKey].(string)).SetMaxWidth(20))
		}
		for j, key := range keyArray {
			if key == hashKey || key == sortKey {
				continue
			}
			if item.Data[key.(string)] == nil {
				t.SetCell(i+1, c+j, tview.NewTableCell(""))
			} else {
				json, err := json.Marshal(item.Data[key.(string)])
				if err != nil {
					panic(err)
				}
				t.SetCell(i+1, c+j, tview.NewTableCell(string(json)).SetMaxWidth(20))
			}
		}
	}
}
