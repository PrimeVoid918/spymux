package spybin

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	config "spymux/src/config"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// var (
// 	promptStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("10")).Bold(true)  // Neon Green
// 	matchStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("14")).Faint(true) // Cyan match count
// 	selectedRowStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("12")).Bold(true)
// 	normalRowStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("7"))
// 	borderStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("8"))
// )

type AppEntry struct {
	DisplayName string
	ExecCmd     string
}

type Model struct {
	Theme *config.AppTheme

	PromptStyle      lipgloss.Style
	MatchStyle       lipgloss.Style
	SelectedRowStyle lipgloss.Style
	NormalRowStyle   lipgloss.Style
	BorderStyle      lipgloss.Style

	Apps     []AppEntry
	Filtered []AppEntry
	Index    int
	Query    string
	Width    int
	Height   int
}

func InitialModel(theme *config.AppTheme) Model {
	apps := ScanDesktopFiles()

	promptStyle := lipgloss.NewStyle().Foreground(theme.Color(10)).Bold(true)
	matchStyle := lipgloss.NewStyle().Foreground(theme.Color(14)).Faint(true)
	selectedRowStyle := lipgloss.NewStyle().Foreground(theme.Color(12)).Bold(true)
	normalRowStyle := lipgloss.NewStyle().Foreground(theme.Color(7))
	borderStyle := lipgloss.NewStyle().Foreground(theme.Color(8))
	return Model{
		Apps:     apps,
		Filtered: apps,

		PromptStyle:      promptStyle,
		MatchStyle:       matchStyle,
		SelectedRowStyle: selectedRowStyle,
		NormalRowStyle:   normalRowStyle,
		BorderStyle:      borderStyle,
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
				launchApp(m.Filtered[m.Index].ExecCmd)
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

	var doc strings.Builder

	stats := m.MatchStyle.Render(fmt.Sprintf("(%d/%d applications)", len(m.Filtered), len(m.Apps)))
	doc.WriteString(fmt.Sprintf("%s %s  %s\n", m.PromptStyle.Render(""), m.Query, stats))

	hrWidth := m.Width
	if hrWidth > 60 {
		hrWidth = 60
	}
	doc.WriteString(m.BorderStyle.Render(strings.Repeat("─", hrWidth)) + "\n")

	maxVisibleRows := m.Height - 4
	if maxVisibleRows <= 0 {
		maxVisibleRows = 5
	}

	startIdx := 0
	if m.Index >= maxVisibleRows {
		startIdx = m.Index - maxVisibleRows + 1
	}

	endIdx := startIdx + maxVisibleRows
	if endIdx > len(m.Filtered) {
		endIdx = len(m.Filtered)
	}

	for i := startIdx; i < endIdx; i++ {
		app := m.Filtered[i]
		if i == m.Index {
			doc.WriteString(m.SelectedRowStyle.Render(fmt.Sprintf(" ▸ %s", app.DisplayName)) + "\n")
		} else {
			doc.WriteString(m.NormalRowStyle.Render(fmt.Sprintf("   %s", app.DisplayName)) + "\n")
		}
	}

	return doc.String()
}

func (m *Model) filterApps() {
	m.Index = 0
	if m.Query == "" {
		m.Filtered = m.Apps
		return
	}

	var out []AppEntry
	q := strings.ToLower(m.Query)
	for _, app := range m.Apps {
		if strings.Contains(strings.ToLower(app.DisplayName), q) {
			out = append(out, app)
		}
	}
	m.Filtered = out
}

func ScanDesktopFiles() []AppEntry {
	home, _ := os.UserHomeDir()
	searchPaths := []string{
		"/usr/share/applications",
		"/usr/local/share/applications",
		filepath.Join(home, ".local/share/applications"),
	}

	var results []AppEntry
	seen := make(map[string]bool)

	for _, path := range searchPaths {
		_ = filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
			if err != nil || info.IsDir() || !strings.HasSuffix(info.Name(), ".desktop") {
				return nil
			}

			pLower := strings.ToLower(p)
			if strings.Contains(pLower, "wine") || strings.Contains(pLower, "programs") ||
				strings.Contains(pLower, "uninstall") || strings.Contains(pLower, "chrome apps") ||
				strings.Contains(pLower, "google maps") {
				return nil
			}

			name, execCmd, valid := parseDesktopFile(p)
			if !valid {
				return nil
			}

			switch name {
			case "Base", "Math", "Draw", "Writer", "Calc", "Impress":
				return nil
			}
			if checkContainsBlacklist(name) {
				return nil
			}

			icon := getIcon(name)
			displayName := fmt.Sprintf("%s %s", icon, name)

			if !seen[displayName] {
				seen[displayName] = true
				results = append(results, AppEntry{DisplayName: displayName, ExecCmd: execCmd})
			}
			return nil
		})
	}
	return results
}

func parseDesktopFile(filePath string) (string, string, bool) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", "", false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var name, execCmd string
	var inMainSection, skipDisplay bool

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			if line == "[Desktop Entry]" {
				inMainSection = true
			} else {
				inMainSection = false
			}
			continue
		}

		if !inMainSection {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		val := strings.TrimSpace(parts[1])

		switch key {
		case "NoDisplay", "Terminal":
			if val == "true" {
				skipDisplay = true
			}
		case "Name":
			if name == "" {
				name = val
			}
		case "Exec":
			if execCmd == "" {
				idx := strings.Index(val, " %")
				if idx != -1 {
					execCmd = val[:idx]
				} else {
					execCmd = val
				}
			}
		}
	}

	if skipDisplay || name == "" || execCmd == "" {
		return "", "", false
	}
	return name, execCmd, true
}

func checkContainsBlacklist(name string) bool {
	low := strings.ToLower(name)
	blacklists := []string{"avahi", "qt6", "qt5", "assistant", "designer", "linguist", "openjdk", "java", "xwayland", "console", "shell", "runtime"}
	for _, b := range blacklists {
		if strings.Contains(low, b) {
			return true
		}
	}
	return false
}

func getIcon(name string) string {
	low := strings.ToLower(name)
	switch {
	case strings.Contains(low, "browser"), strings.Contains(low, "zen"), strings.Contains(low, "chrome"), strings.Contains(low, "tor"):
		return "󰈹"
	case strings.Contains(low, "terminal"), strings.Contains(low, "kitty"):
		return ""
	case strings.Contains(low, "steam"), strings.Contains(low, "game"):
		return "󰊴"
	case strings.Contains(low, "code"), strings.Contains(low, "obsidian"):
		return "󱞂"
	case strings.Contains(low, "torrent"), strings.Contains(low, "qbittorrent"):
		return "󱘖"
	case strings.Contains(low, "bluetooth"), strings.Contains(low, "overskride"):
		return "󰂯"
	case strings.Contains(low, "file"), strings.Contains(low, "ncdu"):
		return "󱏒"
	case strings.Contains(low, "audio"), strings.Contains(low, "mixer"), strings.Contains(low, "wiremix"):
		return "󰓃"
	default:
		return "󰲋"
	}
}

func launchApp(execCmd string) {
	if os.Getenv("HYPRLAND_INSTANCE_SIGNATURE") != "" {
		_ = exec.Command("hyprctl", "dispatch", "exec", "--", execCmd).Run()
		return
	}

	cmd := exec.Command("sh", "-c", execCmd)

	cmd.Stdout = nil
	cmd.Stderr = nil
	cmd.Stdin = nil

	_ = cmd.Start()
	if cmd.Process != nil {
		_ = cmd.Process.Release()
	}
}
