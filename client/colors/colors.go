package colors

import "github.com/charmbracelet/lipgloss"

var (
	FocusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	BlurStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	NoStyle      = lipgloss.NewStyle()
)
