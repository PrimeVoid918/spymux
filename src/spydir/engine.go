package spydir

import (
	"fmt"
	"spymux/src/config"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"os"
	"os/exec"
	"strings"
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

	Dirs     []ZoxideDirs
	DirIndex int

	Err error

	Query string

	Width  int
	Height int
}

type ZoxideDirs struct {
	Score float64
	Path  string
}

func InitialModel() Model {
	dirs, err := getZoxideDirs()
	apps := config.LoadConfig()

	return Model{
		Apps:  apps,
		Dirs:  dirs,
		Err:   err,
		Width: 80,
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
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "backspace":
			if len(m.Query) > 0 {
				m.Query = m.Query[:len(m.Query)-1]
			}
			m.DirIndex = 0
			return m, nil

		case "left", "right", "up", "down", "enter":
			return handleKeyNavigation(m, msg.String())
		}

		if msg.Type == tea.KeyRunes {
			m.Query += msg.String()
			m.DirIndex = 0
			return m, nil
		}
	}
	return m, nil
}

func (model Model) View() string {
	dirs := model.FilteredDirs()
	var strBuilder strings.Builder
	modeBar := ""

	for i, app := range model.Apps {
		if i == model.ModeIndex {
			token := " " + app.Name + " "
			modeBar += selectedStyle.Render(token) + " "
		} else {
			token := "[" + app.Name + "]"
			modeBar += normalStyle.Render(token) + " "
		}
	}
	strBuilder.WriteString(padRight(modeBar, model.Width) + "\n")

	matchesStr := matchCountStyle.Render(fmt.Sprintf("%d", len(dirs)))
	stats := fmt.Sprintf(": %s (%s matches)", model.Query, matchesStr)
	strBuilder.WriteString(padRight(stats, model.Width) + "\n")

	strBuilder.WriteString(faintStyle.Render(strings.Repeat("─", model.Width)) + "\n")

	reservedLines := 4
	maxVisibleDirs := model.Height - reservedLines
	if maxVisibleDirs < 1 {
		maxVisibleDirs = 1
	}

	start := 0
	if model.DirIndex >= maxVisibleDirs {
		start = model.DirIndex - maxVisibleDirs + 1
	}
	end := start + maxVisibleDirs
	if end > len(dirs) {
		end = len(dirs)
	}

	for i := start; i < end; i++ {
		dir := dirs[i]
		var line string
		if i == model.DirIndex {
			rawLine := fmt.Sprintf("> %.2f %s", dir.Score, dir.Path)
			line = selectedStyle.Render(padRight(rawLine, model.Width))
		} else {
			rawLine := fmt.Sprintf(" %.2f %s", dir.Score, dir.Path)
			line = normalStyle.Render(padRight(rawLine, model.Width))
		}
		strBuilder.WriteString(line + "\n")
	}

	return strings.TrimSuffix(strBuilder.String(), "\n")
}

func (model Model) FilteredDirs() []ZoxideDirs {
	if model.Query == "" {
		return model.Dirs
	}
	out := []ZoxideDirs{}
	query := strings.ToLower(model.Query)
	for _, dir := range model.Dirs {
		if strings.Contains(strings.ToLower(dir.Path), query) {
			out = append(out, dir)
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

	if os.Getenv("HYPRLAND_INSTANCE_SIGNATURE") != "" {
		_ = exec.Command("hyprctl", "dispatch", "exec", fullCmd).Run()
		return
	}

	cmd := exec.Command("sh", "-c", fullCmd)

	cmd.Stdout = nil
	cmd.Stderr = nil
	cmd.Stdin = nil

	_ = cmd.Start()
	if cmd.Process != nil {
		_ = cmd.Process.Release()
	}
}

func padRight(s string, width int) string {
	visibleLen := lipgloss.Width(s)
	if visibleLen >= width {
		return s
	}
	return s + strings.Repeat(" ", width-visibleLen)
}

func handleKeyNavigation(model Model, msg string) (tea.Model, tea.Cmd) {
	switch msg {
	case "left":
		model.ModeIndex = (model.ModeIndex - 1 + len(model.Apps)) % len(model.Apps)
	case "right":
		model.ModeIndex = (model.ModeIndex + 1) % len(model.Apps)
	case "up":
		dirs := model.FilteredDirs()
		if len(dirs) > 0 {
			model.DirIndex = (model.DirIndex - 1 + len(dirs)) % len(dirs)
		}
	case "down":
		dirs := model.FilteredDirs()
		if len(dirs) > 0 {
			model.DirIndex = (model.DirIndex + 1) % len(dirs)
		}
	case "enter":
		dirs := model.FilteredDirs()
		if len(dirs) == 0 || len(model.Apps) == 0 {
			return model, nil
		}
		launch(model.Apps[model.ModeIndex], dirs[model.DirIndex].Path)
		return model, tea.Quit
	}
	return model, nil
}

func getZoxideDirs() ([]ZoxideDirs, error) {
	out, err := exec.Command(
		"zoxide",
		"query",
		"--list",
		"--score",
	).Output()

	if err != nil {
		return nil, err
	}

	var dirs []ZoxideDirs

	for _, line := range strings.Split(strings.TrimSpace(string(out)), "\n") {
		parts := strings.Fields(line)

		if len(parts) < 2 {
			continue
		}

		score, err := strconv.ParseFloat(parts[0], 64)
		if err != nil {
			continue
		}

		dirs = append(dirs, ZoxideDirs{
			Score: score,
			Path:  parts[1],
		})
	}

	return dirs, nil
}

func readDirContents(dir string) ([]os.DirEntry, error) {
	return os.ReadDir(dir)
}
