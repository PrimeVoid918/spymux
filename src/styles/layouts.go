package styles

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func ItemPadding(termWidth int, left string, right string, padSym ...string) string {
	// Default to a space if no padding character is provided
	padChar := " "
	if len(padSym) > 0 && padSym[0] != "" {
		padChar = padSym[0]
	}

	usedWidth := lipgloss.Width(left) + lipgloss.Width(right)
	gap := termWidth - usedWidth
	if gap < 0 {
		gap = 0
	}

	return left + strings.Repeat(padChar, gap) + right
}

func FlexRow(elements ...string) string {
	return lipgloss.JoinHorizontal(lipgloss.Bottom, elements...)
}

func FlexCol(elements ...string) string {
	return lipgloss.JoinVertical(lipgloss.Bottom, elements...)
}
