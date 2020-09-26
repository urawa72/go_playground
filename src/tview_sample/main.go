package main

import (
	"os"

	"github.com/urawa72/tview_sample/gui"
)

func start() int {
	if err := gui.New().Run(); err != nil {
		return 1
	}

	return 0
}

func main() {
	os.Exit(start())
}
