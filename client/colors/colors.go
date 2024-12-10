package colors

import "github.com/charmbracelet/lipgloss"

var (
	FocusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	BlurStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	ErrorStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("9"))
	SuccesStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("10")).Bold(true)
	DocStyle     = lipgloss.NewStyle().Margin(1, 2)
	NoStyle      = lipgloss.NewStyle()
)
