package gui

import (
	"encoding/json"
	"fmt"
	"strings"

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
	i.SetWrap(true)
	return i
}

func (i *Info) UpdateView(items *Items, item *Item) {
	var text []string
	c := 0
	text = append(text, fmt.Sprintf("[yellow]%s[white]\t%s", items.HashKey, item.Data[items.HashKey]))
	if items.SortKey != "" {
		c++
		text = append(text, fmt.Sprintf("[yellow]%s[white]\t%s", items.SortKey, item.Data[items.SortKey]))
	}
	for _, key := range items.Attributes {
		if key == items.HashKey || key == items.SortKey {
			continue
		}
		if item.Data[key.(string)] == nil {
			text = append(text, fmt.Sprintf("[yellow]%s[white]\t%s", key, ""))
		} else {
			json, err := json.Marshal(item.Data[key.(string)])
			if err != nil {
				panic(err)
			}
			text = append(text, fmt.Sprintf("[yellow]%s[white]\t%s", key, string(json)))
		}
	}
	i.SetText(strings.Join(text, "\n")).SetTextAlign(tview.AlignLeft)
}
