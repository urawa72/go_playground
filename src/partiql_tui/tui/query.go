package tui

import (
	// "github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type QueryView struct {
	*tview.InputField
	Query string
}

func NewQueryView() *QueryView {
	t := &QueryView{
		InputField: tview.NewInputField(),
	}
	t.SetTitle("Query").SetTitleAlign(tview.AlignLeft)
	t.SetBorder(true)
	t.SetFieldBackgroundColor(tcell.ColorBlack)

	return t
}
