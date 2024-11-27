package models

import tea "github.com/charmbracelet/bubbletea"

type RegisterModel struct{}

func InitRegisterModel() RegisterModel {
	return RegisterModel{}
}

func (m RegisterModel) Init() tea.Cmd {
	return nil
}

func (m RegisterModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return nil, nil
}

func (m RegisterModel) View() string {
	return ""
}
