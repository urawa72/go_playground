package tui

import "github.com/rivo/tview"

type InfoView struct {
	*tview.TextView
}

func NewInfoView() *InfoView {
	info := &InfoView{
		TextView: tview.NewTextView().SetTextAlign(tview.AlignLeft).SetDynamicColors(true),
	}
	info.SetTitleAlign(tview.AlignLeft)
 	moveNavi := "[red::b]j[white]: move down, [red]k[white]: move up, [red]h[white]: move left, [red]l[white]: move right, [red]g[white]: move to top, [red]G[white]: move to bottom, [red]Ctrl-f[white]: next page [red]Ctrl-b[white]: previous page, [red]Ctrl-c[white]: stop pst"
	info.SetText(moveNavi)
	return info
}
