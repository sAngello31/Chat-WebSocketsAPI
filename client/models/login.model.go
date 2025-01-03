package models

import (
	"client_websockets/colors"
	modeldata "client_websockets/model_data"
	utilsmodel "client_websockets/models/utilsModel"
	"client_websockets/services"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type LoginModel struct {
	FocusIndex int
	Inputs     []textinput.Model
	BoolInputs []bool
	IsWrong    bool
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
			return m.enterLogin(), nil
		case "up", "down":
			if m.FocusIndex > 0 && msg.String() == "up" {
				m.FocusIndex--
			} else if m.FocusIndex < len(m.Inputs)-1 && msg.String() == "down" {
				m.FocusIndex++
			}
			cmd := utilsmodel.UpdateCursor(&m.Inputs, m.FocusIndex)
			return m, cmd
		}
	}
	cmd := m.updateInputs(msg)
	return m, cmd
}

func (m LoginModel) View() string {
	var b strings.Builder
	b.WriteString("Iniciar Sesión\n")
	s := utilsmodel.PrintInputs(&m.Inputs, &m.BoolInputs)
	b.WriteString(s)
	if m.IsWrong {
		b.WriteString(colors.ErrorStyle.Render("Error: user y/o contraseña incorrectos"))
	}
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

func (m LoginModel) enterLogin() tea.Model {
	data := modeldata.UserLogin{
		Username: m.Inputs[0].Value(),
		Password: m.Inputs[1].Value(),
	}
	status, token := services.Login(&data)
	if status != 200 {
		m.IsWrong = true
		return m
	}
	services.SaveToken(token)
	return InitMenuModel()
}
