package models

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type MainModel struct {
	Title   string
	Choices []string
	Cursor  int
}

func InitModel() MainModel {
	choices := []string{"Iniciar Sesion", "Registrarse", "Salir de la Aplicacion"}
	return MainModel{Title: "Real-Time Chat App with WebSockets", Choices: choices}
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
		case "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m MainModel) View() string {
	s := m.Title + "\nMenu principal:\n\n"
	for i, choice := range m.Choices {
		cursor := " "
		if m.Cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}
	return s
}
