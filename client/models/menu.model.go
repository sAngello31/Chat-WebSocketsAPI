package models

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type MenuModel struct {
}

func InitMenuModel() MenuModel {
	return MenuModel{}
}

func (m MenuModel) Init() tea.Cmd {
	return nil
}

func (m MenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m MenuModel) View() string {
	var b strings.Builder
	b.WriteString("Bienvenido: \n")
	return b.String()
}
