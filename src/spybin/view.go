package spybin

import (
	"fmt"
	"spymux/src/styles"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func searchBar(m Model) string {
	stats := fmt.Sprintf("(%d/%d applications)", len(m.Filtered), len(m.Apps))
	statsElement := lipgloss.NewStyle().Foreground(m.Theme.Color(12)).Render(stats)

	searchIcon := lipgloss.NewStyle().Foreground(m.Theme.Color(10)).Render(" █")
	query := lipgloss.NewStyle().Underline(true).Render(m.Query)

	layout := styles.FlexRow(fmt.Sprintf("%s %s %s", searchIcon, query, statsElement))

	return layout
}

func selectionList(m Model) string {
	var result strings.Builder

	//! original value is -2
	maxVisibleRows := m.Height - 3
	if maxVisibleRows <= 0 {
		// TODO: Handle terminals that are too small to render the application list.
		// Instead of falling back to an arbitrary number of rows, render a clear
		// "Terminal too small" message, similar to btop.
		maxVisibleRows = 10
	}

	startIdx := 0
	if m.Index >= maxVisibleRows {
		startIdx = m.Index - maxVisibleRows + 1
	}

	endIdx := startIdx + maxVisibleRows
	if endIdx > len(m.Filtered) {
		endIdx = len(m.Filtered)
	}

	selectorIcon := lipgloss.NewStyle().Foreground(m.Theme.Color(14)).Render(" █")
	selectedRow := lipgloss.NewStyle().Foreground(m.Theme.Color(14)).Underline(true)
	notSelectedRow := lipgloss.NewStyle().Foreground(m.Theme.Color(7))

	for i := startIdx; i < endIdx; i++ {
		app := m.Filtered[i]
		content := " " + app.DisplayName
		remaining := m.Width - lipgloss.Width(content)

		//! should remove the trailing /n in the end of the entry
		if i == m.Index {
			if remaining > 0 {
				content += strings.Repeat(" ", remaining)
			}
			result.WriteString(fmt.Sprintf("%s%s", selectorIcon, selectedRow.Render(content)))
			result.WriteByte('\n')
		} else {
			result.WriteString(notSelectedRow.Render(fmt.Sprintf("   %s", app.DisplayName)))
			result.WriteByte('\n')
		}
	}

	layout := styles.FlexCol(result.String())

	return layout
}
