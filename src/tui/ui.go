package tui

import (
	"fmt"
	"os/exec"
	"strings"

	"spymux/src/config" // ⚠️ Change "my-tui-launcher" to match your go.mod module name!

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	selectedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("0")).
			Background(lipgloss.Color("12")).
			Bold(true)

	normalStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("7"))

	faintStyle = lipgloss.NewStyle().
			Faint(true)

	matchCountStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("2")).
			Bold(true)
)

type Model struct {
	Apps      []config.App
	ModeIndex int

	Dirs     []string
	DirIndex int

	Query string

	Width  int
	Height int
}

func LoadDirs() []string {
	out, _ := exec.Command("zoxide", "query", "-l").Output()
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	return lines
}

func InitialModel() Model {
	return Model{
		Apps: config.LoadConfig(),
		Dirs: LoadDirs(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func padRight(s string, width int) string {
	if len(s) >= width {
		return s
	}
	return s + strings.Repeat(" ", width-len(s))
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		return m, nil

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit

		case tea.KeyRunes:
			m.Query += msg.String()
			m.DirIndex = 0

		case tea.KeyBackspace:
			if len(m.Query) > 0 {
				m.Query = m.Query[:len(m.Query)-1]
			}
			m.DirIndex = 0

		default:
			switch msg.String() {
			case "left":
				m.ModeIndex = (m.ModeIndex - 1 + len(m.Apps)) % len(m.Apps)

			case "right":
				m.ModeIndex = (m.ModeIndex + 1) % len(m.Apps)

			case "up":
				dirs := m.FilteredDirs()
				if len(dirs) > 0 {
					m.DirIndex = (m.DirIndex - 1 + len(dirs)) % len(dirs)
				}

			case "down":
				dirs := m.FilteredDirs()
				if len(dirs) > 0 {
					m.DirIndex = (m.DirIndex + 1) % len(dirs)
				}

			case "enter":
				dirs := m.FilteredDirs()
				if len(dirs) == 0 || len(m.Apps) == 0 {
					return m, nil
				}
				launch(m.Apps[m.ModeIndex], dirs[m.DirIndex])
				return m, tea.Quit
			}
		}
	}
	return m, nil
}

func (m Model) View() string {
	dirs := m.FilteredDirs()
	var s strings.Builder

	modeBar := ""
	for i, app := range m.Apps {
		if i == m.ModeIndex {
			token := " " + app.Name + " "
			modeBar += selectedStyle.Render(token) + " "
		} else {
			token := "[" + app.Name + "]"
			modeBar += normalStyle.Render(token) + " "
		}
	}
	s.WriteString(padRight(modeBar, m.Width) + "\n")

	matchesStr := matchCountStyle.Render(fmt.Sprintf("%d", len(dirs)))
	stats := fmt.Sprintf(": %s (%s matches)", m.Query, matchesStr)
	s.WriteString(padRight(stats, m.Width) + "\n")

	s.WriteString(faintStyle.Render(strings.Repeat("─", m.Width)) + "\n")

	for i, d := range dirs {
		var line string
		if i == m.DirIndex {
			rawLine := fmt.Sprintf("> %s", d)
			line = selectedStyle.Render(padRight(rawLine, m.Width))
		} else {
			rawLine := fmt.Sprintf("  %s", d)
			line = normalStyle.Render(padRight(rawLine, m.Width))
		}
		s.WriteString(line + "\n")
	}

	return s.String()
}

func (m Model) FilteredDirs() []string {
	if m.Query == "" {
		return m.Dirs
	}
	var out []string
	q := strings.ToLower(m.Query)
	for _, d := range m.Dirs {
		if strings.Contains(strings.ToLower(d), q) {
			out = append(out, d)
		}
	}
	return out
}

func launch(app config.App, dir string) {
	var fullCmd string
	if strings.Contains(app.Cmd, "{dir}") {
		fullCmd = strings.ReplaceAll(app.Cmd, "{dir}", dir)
	} else {
		fullCmd = fmt.Sprintf("%s %s", app.Cmd, dir)
	}
	exec.Command("hyprctl", "dispatch", "exec", fullCmd).Run()
}
