package gui

import (
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

// func (tx *Text1) UpdateView(g *Gui) {
// 	tableName := g.TableList.Selected()
// 	tx.scanItems(tableName)
// 	// tx.SetText()
// }
//
// func (tx *Text1) scanItems(name string) {
//
// 	params := &dynamodb.ScanInput{
// 		TableName: aws.String(name),
// 	}
//
// 	result, err := Client.Scan(params)
// 	if err != nil {
// 		panic(err)
// 	}
//
//     items := []Item{}
//
// 	for _, item := range result.Items {
// 		var i Item
// 		err = dynamodbattribute.UnmarshalMap(item, &i.Data)
// 		if err != nil {
// 			panic(err)
// 		}
// 		items = append(items, i)
// 	}
//
//     t := tview.NewTable().Select(0, 0).SetFixed(1, 1).SetSelectable(true, false)
//     t.SetBorder(true).SetTitleAlign(tview.AlignLeft)
//
// 	for _, item := range items {
// 		for _, v := range item.Data {
// 			fmt.Println(v)
// 			// tx.SetText(k)
//             // switch v := v.(type) {
// 			// case string:
// 			// 	tx.SetText(v)
// 			// case []string:
// 			// 	tx.SetText(v[0])
// 			// }
// 		}
// 	}
//
// }
//
// type Item struct {
// 	Data map[string]interface{}
// }