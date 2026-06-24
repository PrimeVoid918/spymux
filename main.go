package main

import (
	"flag"
	"fmt"
	"os"
	"spymux/src/config"
	"spymux/src/spybin"
	"spymux/src/spydir"
	"spymux/src/tui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	dirMode := flag.Bool("d", false, "Launch spydir directly")
	binMode := flag.Bool("b", false, "Launch spybin directly")
	flag.Parse()

	theme, themeErr := config.LoadSystemTheme()
	if themeErr != nil {
		fmt.Printf("Error Loading theme: %s", themeErr)
	}

	if *dirMode {
		runSpyDir(theme)
		return
	}
	if *binMode {
		runSpyBin(theme)
		return
	}

	p := tea.NewProgram(tui.InitialPickerModel(), tea.WithAltScreen())
	m, err := p.Run()
	if err != nil {
		fmt.Printf("Error running selection interface: %v\n", err)
		os.Exit(1)
	}

	chosenMode := m.(tui.PickerModel).Choice
	switch chosenMode {
	case tui.ModeSpyDir:
		runSpyDir(theme)
	case tui.ModeSpyBin:
		runSpyBin(theme)
	default:
		os.Exit(0)
	}
}

func runSpyDir(theme *config.AppTheme) {
	p := tea.NewProgram(spydir.InitialModel(theme), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("spydir failed: %v\n", err)
		os.Exit(1)
	}
}

func runSpyBin(theme *config.AppTheme) {
	p := tea.NewProgram(spybin.InitialModel(theme), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("spybin failed: %v\n", err)
		os.Exit(1)
	}
}
