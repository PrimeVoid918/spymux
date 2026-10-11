package spybin

import (
	"fmt"
	"spymux/src/styles"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func searchBar(m Model) string {
	icons := m.Icons
	appIcon := lipgloss.NewStyle().Padding(0, 1).Render(icons.System.Package)

	edges := lipgloss.NewStyle()
	iconAndAppCount := lipgloss.NewStyle().
		Foreground(m.Theme.Color(0)).
		Background(m.Theme.Color(3)).
		Bold(true).
		Render(fmt.Sprintf("%s%-3d", appIcon, len(m.Filtered)))
	leftEdge := edges.Foreground(m.Theme.Color(3)).Render(icons.System.Powerline.SlantRightTop)
	rightEdge := edges.Foreground(m.Theme.Color(3)).Render(icons.System.Powerline.TriangleRight)
	stats := fmt.Sprintf("%s%s%s: ", leftEdge, iconAndAppCount, rightEdge)

	query := lipgloss.NewStyle().Underline(true).Render(m.Query)

	sideItem := lipgloss.NewStyle().Foreground(m.Theme.Color(12)).Render(strconv.Itoa(len(m.Apps)))
	leftWidth := stats + query
	rightWidth := sideItem

	gapMaker := styles.ItemPadding(m.Width, leftWidth, rightWidth)
	return gapMaker
}

func selectionList(m Model) string {
	systemIcons := m.Icons.System

	reservedLines := 2
	maxVisibleRows := m.Height - reservedLines
	if maxVisibleRows <= 0 {
		// TODO: Handle terminals that are too small to render the application list.
		maxVisibleRows = 10
	}

	start := 0
	if m.Index >= maxVisibleRows {
		start = m.Index - maxVisibleRows + 1
	}

	end := start + maxVisibleRows
	if end > len(m.Filtered) {
		end = len(m.Filtered)
	}

	appIconContainer := lipgloss.NewStyle().Padding(0, 1)
	edge := lipgloss.NewStyle()
	leftEdge := edge.Foreground(m.Theme.Accent()).Render(systemIcons.Powerline.LeftBubble)
	rightEdge := edge.Render(systemIcons.Powerline.SlantLeftTop)

	var list strings.Builder
	renderedCount := 0

	if len(m.Filtered) == 0 {
		list.WriteString("  No matching apps found")
		renderedCount = 1
	} else {
		for i := start; i < end; i++ {
			app := m.Filtered[i]

			var lineContainer string
			if i == m.Index {
				appIcon := appIconContainer.
					Foreground(m.Theme.Color(0)).
					Background(m.Theme.Accent()).
					Render(app.AppIcon)
				nameContainer := app.Name
				rightPart := lipgloss.NewStyle().Foreground(m.Theme.Accent()).Underline(true).Bold(true).Render(fmt.Sprintf("%s%s", rightEdge, nameContainer))
				lineContainer = styles.FlexRow(leftEdge, appIcon, rightPart)
			} else {
				appIcon := appIconContainer.
					Foreground(m.Theme.Color(2)).
					Render(app.AppIcon)
				lineContainer = styles.FlexRow(" ", appIcon, " ", app.Name)
			}
			if i != end && i != start {
				list.WriteByte('\n')
			}

			list.WriteString(lineContainer)
			renderedCount++
		}
	}

	for renderedCount < maxVisibleRows {
		list.WriteString("\n")
		renderedCount++
	}

	layout := styles.FlexCol(list.String())
	return layout
}
