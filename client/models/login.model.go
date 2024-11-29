package models

import tea "github.com/charmbracelet/bubbletea"

type LoginModel struct {
}

func InitLoginModel() LoginModel {
	return LoginModel{}
}

func (m LoginModel) Init() tea.Cmd {
	return nil
}

func (m LoginModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return m, tea.Quit
		}

	}
	return m, nil
}

func (m LoginModel) View() string {
	return "Login"
}
