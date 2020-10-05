package gui

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/rivo/tview"
)

type Tables struct {
	*tview.Table
	NameArray []string
}

func (t *Tables) Selected() string {
	row, _ := t.GetSelection()
	return t.NameArray[row]
}

func NewTables() *Tables {
	t := &Tables{
		Table: tview.NewTable().Select(0, 0).SetFixed(1, 1).SetSelectable(true, false),
	}
    t.SetBorder(true).SetTitle("Tables").SetTitleAlign(tview.AlignLeft)

    input := &dynamodb.ListTablesInput{}
    result, _ := Client.ListTables(input)

    table := t.Clear()
	for i, name := range result.TableNames {
		table.SetCell(i, 0, tview.NewTableCell(*name))
		t.NameArray = append(t.NameArray, *name)
	}

	return t
}

