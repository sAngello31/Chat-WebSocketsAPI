package models

import tea "github.com/charmbracelet/bubbletea"

type ExitModel struct {
	Message string
}

func InitExitModel() ExitModel {
	return ExitModel{Message: "Vuelve Pronto!"}
}

func (m ExitModel) Init() tea.Cmd {
	return nil
}

func (m ExitModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		return m, tea.Quit
	}
	return m, nil
}

func (m ExitModel) View() string {
	return m.Message + "\nPresione cualquier tecla para para salir"
}
