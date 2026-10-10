package spydir

import (
	"fmt"
	config "spymux/src/config"
	icons "spymux/src/icons"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"os"
	"os/exec"
	"strings"
)

type Model struct {
	Theme  *config.AppTheme
	Config config.SpydirConfig
	Icons  icons.Icons

	SelectedStyle   lipgloss.Style
	NormalStyle     lipgloss.Style
	FaintStyle      lipgloss.Style
	MatchCountStyle lipgloss.Style

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

func InitialModel(theme *config.AppTheme, conf config.SpydirConfig) Model {
	dirs, err := getZoxideDirs()
	apps := conf.Apps

	selectedStyle := lipgloss.NewStyle().Foreground(theme.Color(0)).Background(theme.Accent()).Bold(true)
	normalStyle := lipgloss.NewStyle().Foreground(theme.Color(7))
	faintStyle := lipgloss.NewStyle().Foreground(theme.Color(2))
	matchCountStyle := lipgloss.NewStyle().Foreground(theme.Color(1))

	return Model{
		Theme:           theme,
		Apps:            apps,
		Icons:           icons.New(),
		Config:          conf,
		Dirs:            dirs,
		Err:             err,
		Width:           80,
		SelectedStyle:   selectedStyle,
		NormalStyle:     normalStyle,
		FaintStyle:      faintStyle,
		MatchCountStyle: matchCountStyle,
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
			m.DirIndex = 0
			return m, nil

		case tea.KeyLeft, tea.KeyRight, tea.KeyUp, tea.KeyDown, tea.KeyEnter:
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
	horizontalLine := lipgloss.NewStyle().Foreground(model.Theme.Color(10)).Render(strings.Repeat("━", model.Width))
	renderView := lipgloss.JoinVertical(
		lipgloss.Top,
		launchApps(model),
		searchBar(model),
		horizontalLine,
		dirList(model),
		footer(model),
	)
	return renderView
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
		targetDir := dirs[model.DirIndex].Path
		_ = exec.Command("zoxide", "add", targetDir).Run()

		launch(model.Apps[model.ModeIndex], targetDir)

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
