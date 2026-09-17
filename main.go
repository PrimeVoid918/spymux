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

	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		os.Exit(1)
	}

	if *dirMode {
		runSpyDir(theme)
		return
	}
	if *binMode {
		runSpyBin(theme, cfg.Spybin)
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
		runSpyBin(theme, cfg.Spybin)
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

func runSpyBin(theme *config.AppTheme, cfg config.SpybinConfig) {
	p := tea.NewProgram(
		spybin.InitialModel(theme, cfg),
		tea.WithAltScreen(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Printf("spybin failed: %v\n", err)
		os.Exit(1)
	}
}
