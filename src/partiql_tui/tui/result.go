package tui

import (
	// "encoding/json"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"

	// "github.com/aws/aws-sdk-go/aws"
	// "github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	// "github.com/deckarep/golang-set"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type ResultView struct {
	*tview.Table
	ItemArray	[]Item
	HashKey		string
	SortKey		string
	Attributes	[]interface{}
}

type Item struct {
	Items []map[string]interface{}
}

var header = []string{
	"HASH",
	"SORT",
}

func NewResultView() *ResultView {
	rv := &ResultView{
		Table: tview.NewTable().Select(0, 0).SetSelectable(true, true),
	}
    rv.SetBorder(true).SetTitle("Results").SetTitleAlign(tview.AlignLeft)
	return rv
}

func (rv *ResultView) UpdateView(t *Tui) {
	table := rv.Clear()

	result, err := rv.RunCmd(t.QueryView.Query)
	if err != nil {
		panic("error")
	}

	table.SetCell(0, 0, &tview.TableCell{
		Text:				result,
		NotSelectable:		true,
		Align:				tview.AlignLeft,
		Color:				tcell.ColorYellow,
		BackgroundColor:	tcell.ColorDefault,
	})
}

func (rv *ResultView) RunCmd(sql string) (string, error) {
	buf := bytes.Buffer{}
	cmd := exec.Command("aws", "dynamodb", "execute-statement", "--statement", sql)
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	if err := cmd.Run(); err != nil {
		return "", errors.New(buf.String())
	}

	jsonStr := []byte(buf.String())
	var item Item
	if err := json.Unmarshal(jsonStr, &item); err != nil {
		fmt.Println("Errrr!")
	}

	for _, m := range item.Items {
		fmt.Println(m)
	}

	return buf.String(), nil
}

// func (rv *ResultView) executeQuery(sql string) {
//
// }
