package gui

import (
	"encoding/json"
	"strings"

	"github.com/rivo/tview"
)

type ItemDetail struct {
	*tview.TextView
}

func NewItemDetail() *ItemDetail {
	i := &ItemDetail{
		TextView: tview.NewTextView().SetTextAlign(tview.AlignRight).SetDynamicColors(true),
	}
	i.SetTitleAlign(tview.AlignLeft).SetTitle("Info").SetBorder(true)
	i.SetWrap(false)
	return i
}

func (i *ItemDetail) UpdateView(g *Gui) {
	item := g.Records.Selecting()
	if item != nil {
		i.ShowItemDetail(g, item)
	}
}

func (i *ItemDetail) ShowItemDetail(g *Gui, item *Item) {
	var detail []string
	for k, v := range item.Data {
		j, err := json.Marshal(v)
		if err != nil {
			panic(err)
		}
		text := string(k) + " " + string(j)
		detail = append(detail, text)
	}

	text := strings.Join(detail, "\n")

	g.App.QueueUpdateDraw(func() {
		i.SetText(text)
	})
}
