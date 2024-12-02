package utilsmodel

import (
	"client_websockets/colors"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
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
