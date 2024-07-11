package styles

import "github.com/charmbracelet/lipgloss"

func ErrorText() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("#ffffff")).
		Background(lipgloss.Color("#ad363e"))
}

func SuccessText() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("#ffffff")).
		Background(lipgloss.Color("#35572e"))
}

func ListItem() lipgloss.Style {
	return lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("63"))
}
