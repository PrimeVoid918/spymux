package spybin

import (
	"fmt"
	"strings"

	config "spymux/src/config"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type AppEntry struct {
	Name        string
	DisplayName string
	ExecCmd     string
}

type Model struct {
	Theme  *config.AppTheme
	Config config.SpybinConfig

	Err error

	Apps     []AppEntry
	Filtered []AppEntry
	Index    int
	Query    string
	Width    int
	Height   int
}

func InitialModel(theme *config.AppTheme, config config.SpybinConfig) Model {
	apps, err := ScanDesktopFiles()

	return Model{
		Theme:    theme,
		Config:   config,
		Apps:     apps,
		Filtered: apps,
		Err:      err,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		return m, nil

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit

		case tea.KeyBackspace:
			if len(m.Query) > 0 {
				m.Query = m.Query[:len(m.Query)-1]
			}
			m.filterApps()
			return m, nil

		case tea.KeyEnter:
			if len(m.Filtered) > 0 && m.Index < len(m.Filtered) {
				target := m.Filtered[m.Index]
				confExecCmd := m.Config.Apps
				for i := range confExecCmd {
					if target.Name == confExecCmd[i].Name {
						launchApp(confExecCmd[i].Cmd)
						return m, tea.Quit
					}
				}
				launchApp(target.ExecCmd)
				return m, tea.Quit
			}
			return m, nil
		}

		switch msg.String() {
		case "up":
			if len(m.Filtered) > 0 {
				m.Index = (m.Index - 1 + len(m.Filtered)) % len(m.Filtered)
			}
		case "down":
			if len(m.Filtered) > 0 {
				m.Index = (m.Index + 1) % len(m.Filtered)
			}
		default:
			if msg.Type == tea.KeyRunes {
				m.Query += msg.String()
				m.filterApps()
			}
		}
	}
	return m, nil
}

func (m Model) View() string {
	if m.Width == 0 {
		return "Scanning system applications..."
	}

	if err := m.Err; err != nil {
		wrapWidth := m.Width
		if wrapWidth < 20 {
			wrapWidth = 20
		}

		return lipgloss.NewStyle().
			Bold(true).
			Foreground(m.Theme.Color(1)).
			Width(wrapWidth). // <--- Enforces automatic word wrapping
			Render(fmt.Sprintf("err: %s", err))
	}

	searchBars := searchBar(m)
	horizontalLine := lipgloss.NewStyle().Foreground(m.Theme.Color(10)).Render(strings.Repeat("━", m.Width))
	list := selectionList(m)

	renderedView := lipgloss.JoinVertical(
		lipgloss.Top,
		searchBars,
		horizontalLine,
		list,
	)
	return renderedView
}
