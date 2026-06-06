package main

import (
	"spymux/src/tui" // ⚠️ Change "my-tui-launcher" to match your go.mod module name!

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(tui.InitialModel())
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
