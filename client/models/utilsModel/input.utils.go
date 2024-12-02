package utilsmodel

import (
	"client_websockets/colors"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

/*
func setInputs() {

}
*/
func PrintInputs(inputs *[]textinput.Model, bool_inputs *[]bool) string {
	var b strings.Builder
	for i := range *inputs {
		b.WriteString((*inputs)[i].View())
		if !(*bool_inputs)[i] {
			b.WriteString(colors.ErrorStyle.Render(" (Obligatorio)"))
		}
		b.WriteRune('\n')
	}
	b.WriteRune('\n')
	b.WriteString(colors.BlurStyle.Render("Presiona 'esc' para retroceder"))
	b.WriteRune('\n')
	b.WriteString(colors.BlurStyle.Render("Presiona 'enter' para aceptar"))
	return b.String()
}

func UpdateCursor(inputs *[]textinput.Model, cursor int) tea.Cmd {
	cmds := make([]tea.Cmd, len(*inputs))
	for i := range *inputs {
		if i == cursor {
			cmds[i] = (*inputs)[i].Focus()
			(*inputs)[i].PromptStyle = colors.FocusedStyle
			(*inputs)[i].TextStyle = colors.FocusedStyle
			continue
		}
		(*inputs)[i].Blur()
		(*inputs)[i].PromptStyle = colors.NoStyle
		(*inputs)[i].TextStyle = colors.NoStyle
	}
	return tea.Batch(cmds...)
}
