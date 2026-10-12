package spydir

import (
	"fmt"
	"path/filepath"
	"spymux/src/styles"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func launchApps(m Model) string {
	appContainer := lipgloss.NewStyle().Bold(true)
	edges := lipgloss.NewStyle().Foreground(m.Theme.Accent())
	icons := m.Icons.System.Powerline

	modeBar := ""
	if len(m.Apps) != 0 {
		for i, app := range m.Apps {
			if i == m.ModeIndex {
				left := edges.Render(icons.TriangleRight)
				right := edges.Render(icons.TriangleLeft)
				token := app.Name
				modeBar += right + appContainer.
					Foreground(m.Theme.Color(0)).
					Background(m.Theme.Accent()).
					Render(token) + left + " "
			} else {
				left := edges.Render(icons.ChevronLeft)
				right := edges.Render(icons.ChevronRight)
				token := app.Name
				modeBar += left + appContainer.
					Foreground(m.Theme.Accent()).
					Render(token) + right + " "
			}
		}
	} else {
		modeBar = "No applications configured. Add applications to ~/.config/spymux/config.toml."
	}

	layout := styles.FlexRow(modeBar)
	return layout
}

func searchBar(m Model) string {
	icons := m.Icons.System
	dirs := m.FilteredDirs()
	iconWrapper := lipgloss.NewStyle().Padding(0, 1)
	dirIcon := iconWrapper.Render(icons.DirectoryOpen)
	container := lipgloss.NewStyle()

	matchesStr := container.Render(fmt.Sprintf("%-4d", len(dirs)))
	layout := styles.FlexRow(dirIcon, matchesStr, ": ", m.Query)
	return layout
}

func dirList(m Model) string {
	dirs := m.FilteredDirs()

	reservedLines := 4
	maxVisibleDirs := m.Height - reservedLines
	if maxVisibleDirs < 1 {
		maxVisibleDirs = 1
	}

	start := 0
	if m.DirIndex >= maxVisibleDirs {
		start = m.DirIndex - maxVisibleDirs + 1
	}
	end := start + maxVisibleDirs
	if end > len(dirs) {
		end = len(dirs)
	}

	var list strings.Builder
	renderedCount := 0

	if len(dirs) == 0 {
		list.WriteString("  No matching directories found\n")
		renderedCount = 1
	} else {
		for i := start; i < end; i++ {
			dir := dirs[i]
			score := fmt.Sprintf("%4.0f", dir.Score)

			var lineContainer string
			if i == m.DirIndex {
				leftEdge := lipgloss.NewStyle().Foreground(m.Theme.Accent()).Render(m.Icons.System.Powerline.TriangleLeft)

				dirName := filepath.Base(dir.Path)
				dirPath := lipgloss.NewStyle().
					Foreground(m.Theme.Color(0)).
					Background(m.Theme.Accent()).
					Render(dirName + strings.Repeat(" ", m.Width))

				lineContainer = styles.FlexRow(score, " ", leftEdge+dirPath)
			} else {
				dirName := filepath.Base(dir.Path)
				dirPath := lipgloss.NewStyle().Render(dirName)
				lineContainer = styles.FlexRow(score, " ", dirPath)
			}
			if i != end && i != start {
				list.WriteString("\n")
			}
			list.WriteString(lineContainer)
			renderedCount++
		}
	}

	for renderedCount < maxVisibleDirs {
		list.WriteString("\n")
		renderedCount++
	}

	return list.String()
}

func footer(m Model) string {
	dirs := m.FilteredDirs()

	if len(dirs) <= 0 {
		return ""
	}

	dirPath := dirs[m.DirIndex].Path
	form := lipgloss.NewStyle().
		Foreground(m.Theme.FG()).
		Background(m.Theme.Color(2)).
		Render(dirPath + strings.Repeat(" ", m.Width))

	return styles.FlexRow(form)
}
