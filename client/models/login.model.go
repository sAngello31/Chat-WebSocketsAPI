package models

import (
	"client_websockets/colors"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type LoginModel struct {
	FocusIndex int
	Inputs     []textinput.Model
	BoolInputs []bool
}

func InitLoginModel() LoginModel {
	m := LoginModel{
		Inputs:     make([]textinput.Model, 2),
		BoolInputs: make([]bool, 2),
	}
	var t textinput.Model
	for i := range m.Inputs {
		t = textinput.New()
		t.Cursor.Style = colors.FocusedStyle
		t.CharLimit = 32
		switch i {
		case 0:
			t.Placeholder = "Nickname*"
			t.Focus()
			t.PromptStyle = colors.FocusedStyle
			t.TextStyle = colors.FocusedStyle
		case 1:
			t.Placeholder = "Password*"
			t.EchoMode = textinput.EchoPassword
			t.EchoCharacter = '*'
		}
		m.Inputs[i] = t
	}
	return m
}

func (m LoginModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m LoginModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return InitModel(), nil
		case "enter":
			return InitModel(), nil
		case "up", "down":
			if m.FocusIndex > 0 && msg.String() == "up" {
				m.FocusIndex--
			} else if m.FocusIndex < len(m.Inputs)-1 && msg.String() == "down" {
				m.FocusIndex++
			}
			cmd := m.updateCursor()
			return m, cmd
		}
	}
	cmd := m.updateInputs(msg)
	return m, cmd
}

func (m LoginModel) View() string {
	var b strings.Builder
	b.WriteString("Iniciar Sesión\n")
	for i := range m.Inputs {
		b.WriteString(m.Inputs[i].View())
		if !m.BoolInputs[i] {
			b.WriteString(colors.ErrorStyle.Render(" (Obligatorio)"))
		}
		b.WriteRune('\n')
	}
	b.WriteRune('\n')
	b.WriteString(colors.BlurStyle.Render("Seleccione 'esc' para retroceder"))
	b.WriteRune('\n')
	b.WriteString(colors.BlurStyle.Render("Seleccione 'enter' para aceptar\n"))

	return b.String()
}

func (m LoginModel) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.Inputs))
	for i := range m.Inputs {
		m.Inputs[i], cmds[i] = m.Inputs[i].Update(msg)
		if len(m.Inputs[i].Value()) > 0 {
			m.BoolInputs[i] = true
		} else {
			m.BoolInputs[i] = false
		}
	}
	return tea.Batch(cmds...)
}

func (m LoginModel) updateCursor() tea.Cmd {
	cmds := make([]tea.Cmd, len(m.Inputs))
	for i := range m.Inputs {
		if i == m.FocusIndex {
			cmds[i] = m.Inputs[i].Focus()
			m.Inputs[i].PromptStyle = colors.FocusedStyle
			m.Inputs[i].TextStyle = colors.FocusedStyle
			continue
		}
		m.Inputs[i].Blur()
		m.Inputs[i].PromptStyle = colors.NoStyle
		m.Inputs[i].TextStyle = colors.NoStyle

	}
	return tea.Batch(cmds...)
}
