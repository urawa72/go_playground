package gui

import (
	"github.com/rivo/tview"
)

type Info struct {
	*tview.TextView
}

func NewInfo() *Info {
	i := &Info{
		TextView: tview.NewTextView().SetTextAlign(tview.AlignRight).SetDynamicColors(true),
	}
	i.SetTitleAlign(tview.AlignLeft).SetTitle("Info").SetBorder(true)
	i.SetWrap(false)
	return i
}
