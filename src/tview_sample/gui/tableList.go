package gui

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/rivo/tview"
)

type TableList struct {
	*tview.Table
	Tables []string
}

func (t *TableList) Selected() string {
	row, _ := t.GetSelection()
	return t.Tables[row]
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
		t.Tables = append(t.Tables, *name)
	}

	return t
}

