package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Tui struct {
	App			*tview.Application
	QueryView	*QueryView
	ResultView	*ResultView
	Info		*InfoView
	Pages		*tview.Pages
	Panels
}

type Panels struct {
	Current	int
	Panels	[]tview.Primitive
}

func New() *Tui {
	// NewClient()
	queryView := NewQueryView()
	resultView := NewResultView()
	infoView := NewInfoView()

	t := &Tui {
		App:		tview.NewApplication(),
		QueryView: 	queryView,
		ResultView:	resultView,
		Info:		infoView,
	}

	t.Panels = Panels{
		Panels: []tview.Primitive{
			queryView,
			resultView,
		},
	}

	return t
}

func (t *Tui) queryKeybindings() {
 	t.QueryView.SetDoneFunc(func(key tcell.Key) {
		switch key {
		case tcell.KeyEscape:
			// g.App.Stop()
		case tcell.KeyEnter:
			text := t.QueryView.GetText()
			t.QueryView.Query = text
			t.ResultView.UpdateView(t)
			// g.nextPanel()
		}
	}).SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		// g.GrobalKeybind(event)
		return event
	})
}

func (t *Tui) Run() error {
	// result, err := runCmd()
	// if err != nil {
	// 	panic("error")
	// }
	// fmt.Println(result)

	grid := tview.NewGrid().
		SetRows(3, 0, 4).
		AddItem(t.QueryView, 0, 0, 1, 1, 0, 0, true).
		AddItem(t.ResultView, 1, 0, 1, 1, 0, 0, true).
		AddItem(t.Info, 2, 0, 1, 1, 0, 0, false)

	t.Pages = tview.NewPages().AddAndSwitchToPage("main", grid, true)

	t.queryKeybindings()
	// t.tableListKeybind()
	// t.itemsKeybindings()
	// t.infoKeybinding()
	// t.itemDetailKeybinding()

	if err := t.App.SetRoot(t.Pages, true).SetFocus(t.QueryView).Run(); err != nil {
		t.App.Stop()
		return err
	}

	return nil
}
