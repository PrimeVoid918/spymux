package config

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	PromptStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("10")).Bold(true)  // Neon Green
	MatchStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("14")).Faint(true) // Cyan match count
	SelectedRowStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("12")).Bold(true)
	NormalRowStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("7"))
	BorderStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("8"))
)
