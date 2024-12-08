package models

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type MainModel struct {
	Title          string
	Choices        map[string]tea.Model
	ChoicesDisplay []string
	Cursor         int
}

func InitModel() MainModel {
	choices := []string{"Iniciar Sesión", "Registrar Nuevo Usuario", "Salir de la App"}
	choiceMap := make(map[string]tea.Model, len(choices))
	choiceMap[choices[0]] = InitLoginModel()
	choiceMap[choices[1]] = InitRegisterModel()
	choiceMap[choices[2]] = InitExitModel()
	return MainModel{Title: "Real-Time Chat App with WebSockets", Choices: choiceMap, ChoicesDisplay: choices}
}

func (m MainModel) Init() tea.Cmd {
	return tea.SetWindowTitle("Real-Time Chat App with WebSockets")
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			if m.Cursor > 0 {
				m.Cursor -= 1
			}
		case "down":
			if m.Cursor < len(m.Choices)-1 {
				m.Cursor += 1
			}
		case "enter":
			return m.Choices[m.ChoicesDisplay[m.Cursor]], nil
		}
	}
	return m, nil
}

func (m MainModel) View() string {
	var b strings.Builder
	b.WriteString(m.Title)
	b.WriteString("\nMenu Principal:\n\n")
	for i, choice := range m.ChoicesDisplay {
		cursor := " "
		if m.Cursor == i {
			cursor = ">"
		}
		b.WriteString(cursor)
		b.WriteString(" " + choice)
		b.WriteRune('\n')
	}
	return b.String()
}
