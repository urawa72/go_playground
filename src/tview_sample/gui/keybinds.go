package gui

import "github.com/gdamore/tcell"

func (g *Gui) globalKeybind(event *tcell.EventKey) {
	switch event.Key() {
	case tcell.KeyTab:
		g.nextPanel()
	}
}

func (g *Gui) tableListKeybind() {
	g.Tables.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEnter {
			g.Items.UpdateView(g)
		}
        switch event.Rune() {
        case 'P':
			g.PutItem("Modal Test", "OK", g.Tables)
		}
		g.globalKeybind(event)
		return event
	})
}


func (g *Gui) itemsKeybindings() {
    g.Items.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		g.globalKeybind(event)
		return event
	})
}

func (g *Gui) infoKeybinding() {
    g.Info.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		g.globalKeybind(event)
		return event
	})
}
