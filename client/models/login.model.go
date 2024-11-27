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
	return nil, nil
}

func (m LoginModel) View() string {
	return ""
}
