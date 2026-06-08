// src/tui/picker.go
package tui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type MenuOption string

const (
	ModeSpyDir MenuOption = "SPYDIR (Directories)"
	ModeSpyBin MenuOption = "SPYBIN (Applications)"
)

var options = []MenuOption{ModeSpyDir, ModeSpyBin}

var (
	activeStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("0")).Background(lipgloss.Color("12")).Padding(0, 1).Bold(true)
	inactiveStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("7")).Background(lipgloss.Color("8")).Padding(0, 1)
	titleStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("14")).Bold(true)
)

type PickerModel struct {
	Cursor int
	Choice MenuOption
}

func InitialPickerModel() PickerModel {
	return PickerModel{Cursor: 0, Choice: ""}
}

func (m PickerModel) Init() tea.Cmd { return nil }

func (m PickerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			return m, tea.Quit
		case "left", "h", "shift+tab":
			m.Cursor = (m.Cursor - 1 + len(options)) % len(options)
		case "right", "l", "tab":
			m.Cursor = (m.Cursor + 1) % len(options)
		case "enter", "space":
			m.Choice = options[m.Cursor]
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m PickerModel) View() string {
	var b strings.Builder
	b.WriteString(titleStyle.Render("🚀 Select Spymux Engine Target:") + "\n\n")

	var renderedOptions []string
	for i, opt := range options {
		if i == m.Cursor {
			renderedOptions = append(renderedOptions, activeStyle.Render(string(opt)))
		} else {
			renderedOptions = append(renderedOptions, inactiveStyle.Render(string(opt)))
		}
	}
	b.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, renderedOptions...) + "\n")
	return b.String()
}
