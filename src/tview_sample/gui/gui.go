package gui

import (
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
		g.globalKeybind(event)
		return event
	})
}

func (g *Gui) recordsKeybindings() {
    g.Records.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		g.globalKeybind(event)
		return event
	})
 	// g.Records.SetSelectionChangedFunc(func(row, col int) {
	// 	if row < 1 {
	// 		return
	// 	}
	// 	g.ItemDetail.UpdateView(g)
	// })
}

func (g *Gui) itemDetailKeybinding() {
	g.ItemDetail.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		g.globalKeybind(event)
		return event
	})
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
		AddItem(g.Records, 0, 1, 1, 1, 0, 0, true).
		AddItem(g.ItemDetail, 1, 1, 2, 1, 0, 0, true)

	grid := tview.NewGrid().
		SetRows(1, 0, 2).
		SetColumns(30).
		AddItem(g.Header, 0, 0, 1, 1, 0, 0, true).
		AddItem(mainGrid, 1, 0, 1, 2, 0, 0, true).
		AddItem(g.Footer, 2, 0, 1, 2, 0, 0, true)

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
