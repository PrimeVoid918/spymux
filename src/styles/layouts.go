package styles

import (
	// tea "github.com/charmbracelet/bubbletea"
	// "strings"

	"github.com/charmbracelet/lipgloss"
)

func FlexRow(elements ...string) string {
	return lipgloss.JoinHorizontal(lipgloss.Bottom, elements...)
}

func FlexCol(elements ...string) string {
	return lipgloss.JoinVertical(lipgloss.Bottom, elements...)
}
