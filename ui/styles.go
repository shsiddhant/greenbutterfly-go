package ui

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"
)

var (
	ColorPrimary = lipgloss.Color("5")   // Purple
	ColorMuted   = lipgloss.Color("244") // Grey
	ColorGreen   = lipgloss.Color("2")   // Green
	ColorYellow  = lipgloss.Color("3")   // Yellow
	ColorSad     = lipgloss.Color("12")  // Blue
	ColorLove    = lipgloss.Color("9")   //Red
)

var (
	TitleStyle = lipgloss.NewStyle().Bold(true).Foreground(ColorLove)
	DimStyle   = lipgloss.NewStyle().Foreground(ColorMuted)

	headerTextStyle = lipgloss.NewStyle().Bold(true).Foreground(ColorPrimary)
	PastDaysStyle   = lipgloss.NewStyle().Bold(false).Foreground(ColorSad)
	QuoteStyle      = lipgloss.NewStyle().Italic(true).Foreground(ColorLove)
)

func RenderHeader(title string, width int) string {
	text := fmt.Sprintf(" %s ", strings.ToUpper(title))

	return lipgloss.PlaceHorizontal(
		width,
		lipgloss.Left,
		headerTextStyle.Render(text),
		lipgloss.WithWhitespaceChars("─"),
	) + "\n"
}
