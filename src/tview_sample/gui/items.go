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

type Records struct {
	*tview.Table
	Items	[]Item
}

type KeyDetail struct {
	KeyType string
	AttributeType *string
}

var header = []string{
	"HASH",
	"SORT",
}

func NewRecords() *Records {
	r := &Records{
		Table: tview.NewTable().Select(0, 0).SetSelectable(true, true),
	}
    r.SetBorder(true).SetTitleAlign(tview.AlignLeft)
	return r
}

func (r *Records) Selecting() *Item {
	if len(r.Items) == 0 {
		return nil
	}
	row, _ := r.GetSelection()
	if row < 0 {
		return nil
	}
	fmt.Println(row)
	return &r.Items[row-1]
}

func (r *Records) UpdateView(g *Gui) {
	tableName := g.TableList.Selected()
	r.scanItems(tableName)
	g.nextPanel()
}

func (r *Records) scanItems(name string) {
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

    r.Items = []Item{}

	for _, item := range result.Items {
		var i Item
		err = dynamodbattribute.UnmarshalMap(item, &i.Data)
		if err != nil {
			panic(err)
		}
		r.Items = append(r.Items, i)
	}

	keys := mapset.NewSet()
	for _, item := range r.Items {
		for k, _ := range item.Data {
			keys.Add(k)
		}
	}
	keyArray := keys.ToSlice()
	sort.Slice(keyArray, func(i, j int) bool { return keyArray[i].(string) <  keyArray[j].(string) })

	t := r.Clear()

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

	for i, item := range r.Items {
		c := 0
		t.SetCell(i+1, c, tview.NewTableCell(item.Data[hashKey].(string)))
		if sortKey != "" {
			c++
			t.SetCell(i+1, c, tview.NewTableCell(item.Data[sortKey].(string)))
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
				t.SetCell(i+1, c+j, tview.NewTableCell(string(json)))
			}
		}
	}

	// for i, item := range r.Items {
	// 	var hKey string
	// 	var sKey string
	// 	for k, v := range item.Data {
	// 		j, err := json.Marshal(v)
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		if hashKey == string(k) {
	// 			hKey = string(j)
	// 		}
	// 		if sortKey == string(k) {
	// 			sKey = string(j)
	// 		}
	// 	}
	// 	t.SetCell(i+1, 0, tview.NewTableCell(hKey))
	// 	t.SetCell(i+1, 1, tview.NewTableCell(sKey))
	// }
}

type Item struct {
	Data map[string]interface{}
}
