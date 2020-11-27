package tui

import (
	"github.com/rivo/tview"
)

type Tui struct {
	App			*tview.Application
	QueryView	*QueryView
	Results		*Items
	Info		*InfoView
	Pages		*tview.Pages
	Panels
}

type Panels struct {
	Current	int
	Panels	[]tview.Primitive
}

// func (g *Gui) switchPanel(p tview.Primitive) *tview.Application {
// 	return g.App.SetFocus(p)
// }
//
// func (g *Gui) nextPanel() {
// 	idx := (g.Panels.Current + 1) % len(g.Panels.Panels)
// 	g.Panels.Current = idx
// 	g.switchPanel(g.Panels.Panels[g.Panels.Current])
// }
//
// func (g *Gui) prevPanel() {
// 	g.Panels.Current--
//
// 	if g.Panels.Current < 0 {
// 		g.Current = len(g.Panels.Panels) - 1
// 	} else {
// 		idx := (g.Panels.Current) % len(g.Panels.Panels)
// 		g.Panels.Current = idx
// 	}
// 	g.switchPanel(g.Panels.Panels[g.Panels.Current])
// }
//
// func convertParams(f *tview.Form) {
// 	cnt := f.GetFormItemCount()
// 	for i := 0; i < cnt; i++ {
// 		item := f.GetFormItem(i)
// 		switch item.(type) {
// 		case *tview.InputField:
// 			text := item.(*tview.InputField).GetText()
// 			fmt.Println(len(text))
// 		case *tview.DropDown:
// 			label, index := item.(*tview.DropDown).GetCurrentOption()
// 			fmt.Println(label, index)
// 		}
// 	}
// }
//
// func (g *Gui) PutItem(message, doneLabel string, primitive tview.Primitive) {
// 	options := []string{"String", "Number", "Binary"}
// 	form := tview.NewForm()
// 	form.SetBorder(true)
// 	form.SetTitleAlign(tview.AlignLeft)
// 	form.SetTitle("Put Item")
//     form.AddInputField("Hash Key       ", "", 80, nil, nil).
//     	AddInputField("Sort Key       ", "", 80, nil, nil).
// 		AddButton("Add", func() {
// 			form.AddInputField("Attribute Name", "", 80, nil, nil).
// 				AddDropDown("Attribute Type", options, 0, func(option string, index int) {
// 				}).
// 				AddInputField("Value", "", 80, nil, nil)
// 		}).
// 		AddButton("Execute", func() {
// 			convertParams(form)
// 			g.CloseAndSwitchPanel("form", g.Tables)
// 		}).
// 		AddButton("Cancel", func() {
// 			g.CloseAndSwitchPanel("form", g.Tables)
// 		})
//
// 	g.Pages.AddAndSwitchToPage("form", g.Modal(form, 50, 29), true).ShowPage("main")
// }
//
// func (g *Gui) CloseAndSwitchPanel(removePrimitive string, primitive tview.Primitive) {
// 	g.Pages.RemovePage(removePrimitive).ShowPage("main")
// 	g.switchPanel(primitive)
// }

// func (t *Tui) Modal(p tview.Primitive, width, height int) tview.Primitive {
// 	return tview.NewGrid().
// 		SetColumns(0, width, 0).
// 		SetRows(0, height, 0).
// 		AddItem(p, 1, 1, 1, 1, 0, 0, true)
// }
//

func New() *Tui {
	queryView := NewQueryView()
	results := NewResults()
	infoView := NewInfoView()
	NewClient()

	t := &Tui {
		App:		tview.NewApplication(),
		QueryView: queryView,
		Results:	results,
		Info:		infoView,
	}

	t.Panels = Panels{
		Panels: []tview.Primitive{
			queryView,
			results,
		},
	}

	return t
}

func (t *Tui) Run() error {
	grid := tview.NewGrid().
		SetRows(3, 0, 4).
		AddItem(t.QueryView, 0, 0, 1, 1, 0, 0, true).
		AddItem(t.Results, 1, 0, 1, 1, 0, 0, true).
		AddItem(t.Info, 2, 0, 1, 1, 0, 0, false)

	t.Pages = tview.NewPages().AddAndSwitchToPage("main", grid, true)

	// t.tableListKeybind()
	// t.itemsKeybindings()
	// g.infoKeybinding()
	// t.itemDetailKeybinding()

	if err := t.App.SetRoot(t.Pages, true).SetFocus(t.QueryView).Run(); err != nil {
		t.App.Stop()
		return err
	}

	return nil
}
