package gui

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/rivo/tview"
)

type TableList struct {
	*tview.Table
	dbTables []string
}

func (t *TableList) Selected() string {
	row, _ := t.GetSelection()
	return t.dbTables[row]
}

func NewTableList() *TableList {
	t := &TableList{
		Table: tview.NewTable().Select(0, 0).SetFixed(1, 1).SetSelectable(true, false),
	}
    t.SetBorder(true).SetTitle("tests").SetTitleAlign(tview.AlignLeft)

    input := &dynamodb.ListTablesInput{}
    result, _ := Client.ListTables(input)

    table := t.Clear()
	for i, name := range result.TableNames {
		table.SetCell(i, 0, tview.NewTableCell(*name))
		t.dbTables = append(t.dbTables, *name)
	}

	return t
}

