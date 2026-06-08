package main

import (
	"flag"
	"fmt"
	"os"
	"spymux/src/spybin"
	"spymux/src/spydir"
	"spymux/src/tui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// Define fast CLI execution flags
	dirMode := flag.Bool("d", false, "Launch spydir directly")
	binMode := flag.Bool("b", false, "Launch spybin directly")
	flag.Parse()

	// 1. Instant execution paths via flags
	if *dirMode {
		runSpyDir()
		return
	}
	if *binMode {
		runSpyBin()
		return
	}

	// 2. Fallback execution path: Show picker TUI
	p := tea.NewProgram(tui.InitialPickerModel(), tea.WithAltScreen())
	m, err := p.Run()
	if err != nil {
		fmt.Printf("Error running selection interface: %v\n", err)
		os.Exit(1)
	}

	// Route based on what the user chose in the fallback UI
	chosenMode := m.(tui.PickerModel).Choice
	switch chosenMode {
	case tui.ModeSpyDir:
		runSpyDir()
	case tui.ModeSpyBin:
		runSpyBin()
	default:
		os.Exit(0)
	}
}

func runSpyDir() {
	p := tea.NewProgram(spydir.InitialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("spydir failed: %v\n", err)
		os.Exit(1)
	}
}

func runSpyBin() {
	p := tea.NewProgram(spybin.InitialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("spybin failed: %v\n", err)
		os.Exit(1)
	}
}
