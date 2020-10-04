package gui

import (
	"fmt"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type Gui struct {
	App			*tview.Application
	Pages		*tview.Pages
	TableList	*TableList
	Records		*Records
	ItemDetail	*ItemDetail
	Header		*tview.TextView
	Footer		*tview.TextView
	Panels
}

type Panels struct {
	Current	int
	Panels	[]tview.Primitive
}

func (g *Gui) switchPanel(p tview.Primitive) *tview.Application {
	return g.App.SetFocus(p)
}

func (g *Gui) nextPanel() {
	idx := (g.Panels.Current + 1) % len(g.Panels.Panels)
	g.Panels.Current = idx
	g.switchPanel(g.Panels.Panels[g.Panels.Current])
}

func (g *Gui) globalKeybind(event *tcell.EventKey) {
	switch event.Key() {
	case tcell.KeyTab:
		g.nextPanel()
	}
}

func (g *Gui) tableListKeybind() {
	g.TableList.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEnter {
			g.Records.UpdateView(g)
		}
        switch event.Rune() {
        case 'P':
			g.PutItem("Modal Test", "OK", g.TableList)
		}
		g.globalKeybind(event)
		return event
	})
}


func (g *Gui) recordsKeybindings() {
    g.Records.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		g.globalKeybind(event)
		return event
	})
}

func (g *Gui) itemDetailKeybinding() {
	g.ItemDetail.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		g.globalKeybind(event)
		return event
	})
}

func convertParams(f *tview.Form) {
	cnt := f.GetFormItemCount()
	for i := 0; i < cnt; i++ {
		item := f.GetFormItem(i)
		switch item.(type) {
		case *tview.InputField:
			text := item.(*tview.InputField).GetText()
			fmt.Println(len(text))
		case *tview.DropDown:
			label, index := item.(*tview.DropDown).GetCurrentOption()
			fmt.Println(label, index)
		}
	}
}

func (g *Gui) PutItem(message, doneLabel string, primitive tview.Primitive) {
	options := []string{"String", "Number", "Binary"}
	form := tview.NewForm()
	form.SetBorder(true)
	form.SetTitleAlign(tview.AlignLeft)
	form.SetTitle("Put Item")
    form.AddInputField("Hash Key       ", "", 80, nil, nil).
    	AddInputField("Sort Key       ", "", 80, nil, nil).
		AddButton("Add", func() {
			form.AddInputField("Attribute Name", "", 80, nil, nil).
				AddDropDown("Attribute Type", options, 0, func(option string, index int) {
				}).
				AddInputField("Value", "", 80, nil, nil)
		}).
		AddButton("Execute", func() {
			convertParams(form)
			g.CloseAndSwitchPanel("form", g.TableList)
		}).
		AddButton("Cancel", func() {
			g.CloseAndSwitchPanel("form", g.TableList)
		})

	g.Pages.AddAndSwitchToPage("form", g.Modal(form, 50, 29), true).ShowPage("main")
}

func (g *Gui) CloseAndSwitchPanel(removePrimitive string, primitive tview.Primitive) {
	g.Pages.RemovePage(removePrimitive).ShowPage("main")
	g.switchPanel(primitive)
}

func (g *Gui) Modal(p tview.Primitive, width, height int) tview.Primitive {
	return tview.NewGrid().
		SetColumns(0, width, 0).
		SetRows(0, height, 0).
		AddItem(p, 1, 1, 1, 1, 0, 0, true)
}

func New() *Gui {
	header := tview.NewTextView().SetDynamicColors(true)
    header.SetTitleAlign(tview.AlignLeft).SetTitle("Header")

	footer := tview.NewTextView().SetDynamicColors(true)
    footer.SetTitleAlign(tview.AlignLeft).SetTitle("Footer")

	NewDbClient()

	tableList := NewTableList()

	records := NewRecords()

	itemDetail := NewItemDetail()

	g := &Gui {
		App:		tview.NewApplication(),
		TableList:	tableList,
		Records: 	records,
		ItemDetail:	itemDetail,
		Header: 	header,
		Footer:		footer,
	}

	g.Panels = Panels{
		Panels: []tview.Primitive{
			tableList,
			records,
			itemDetail,
		},
	}

	return g
}

func (g *Gui) Run() error {
	mainGrid := tview.NewGrid().SetRows(0, 0, 0).SetColumns(30, 0).
		AddItem(g.TableList, 0, 0, 3, 1, 0, 0, true).
		AddItem(g.Records, 0, 1, 2, 1, 0, 0, true).
		AddItem(g.ItemDetail, 2, 1, 1, 1, 0, 0, true)

	grid := tview.NewGrid().
		SetRows(0).
		SetColumns(30).
		AddItem(mainGrid, 0, 0, 1, 2, 0, 0, true)

	g.Pages = tview.NewPages().AddAndSwitchToPage("main", grid, true)

	g.tableListKeybind()
	g.recordsKeybindings()
	g.itemDetailKeybinding()

	if err := g.App.SetRoot(g.Pages, true).SetFocus(g.TableList).Run(); err != nil {
		g.App.Stop()
		return err
	}

	return nil
}
