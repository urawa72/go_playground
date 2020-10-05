package gui

import (
	"encoding/json"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type ItemDetail struct {
	*tview.Table
}

func NewItemDetail() *ItemDetail {
	id := &ItemDetail{
		Table: tview.NewTable().Select(0, 0).SetSelectable(true, true),
	}
    id.SetBorder(true).SetTitle("Item Detail").SetTitleAlign(tview.AlignLeft)
	return id
}

func (id *ItemDetail) UpdateView(items *Items, item *Item) {
	t := id.Clear()

	r := 0
	t.SetCell(r, 0, &tview.TableCell{
		Text:				items.HashKey,
		Align:				tview.AlignLeft,
		Color:				tcell.ColorYellow,
		BackgroundColor:	tcell.ColorDefault,
	})
	t.SetCell(r, 1, tview.NewTableCell(item.Data[items.HashKey].(string)))
	if items.SortKey != "" {
		r++
		t.SetCell(r, 0, &tview.TableCell{
			Text:				items.SortKey,
			Align:				tview.AlignLeft,
			Color:				tcell.ColorYellow,
			BackgroundColor:	tcell.ColorDefault,
		})
		t.SetCell(r, 1, tview.NewTableCell(item.Data[items.SortKey].(string)))
	}
	r++
	for i, key := range items.Attributes {
		if key.(string) == items.HashKey || key.(string) == items.SortKey {
			continue
		}
		t.SetCell(r+i, 0, &tview.TableCell{
			Text:				key.(string),
			Align:				tview.AlignLeft,
			Color:				tcell.ColorYellow,
			BackgroundColor:	tcell.ColorDefault,
		})
		if item.Data[key.(string)] == nil {
			t.SetCell(r+i, 1, tview.NewTableCell(""))
		} else {
			json, err := json.Marshal(item.Data[key.(string)])
			if err != nil {
				panic(err)
			}
			t.SetCell(r+i, 1, tview.NewTableCell(string(json)))
		}
	}
}
