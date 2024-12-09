package models

import (
	"client_websockets/colors"
	modeldata "client_websockets/model_data"
	"client_websockets/services"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type MenuModel struct {
	User     modeldata.User
	Contacts []modeldata.User
	Cursor   int
}

func InitMenuModel() MenuModel {
	return MenuModel{
		User: *services.GetUserData(),
	}
}

func (m MenuModel) Init() tea.Cmd {
	return nil
}

func (m MenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			return InitModel(), nil
		case "esc":
			return InitModel(), nil
		}
	}
	return m, nil
}

func (m MenuModel) View() string {
	contact_number := strconv.Itoa(m.User.ContactNumber)
	var b strings.Builder
	b.WriteString("Bienvenido: " + m.User.Username + "\n")
	b.WriteString("Tu numero de contacto: # " + contact_number + "\n\n")
	b.WriteString("Seleccione con quien quiere hablar:\n")

	b.WriteRune('\n')
	b.WriteString(colors.BlurStyle.Render("Presiona 'esc' para cerrar sesión"))
	b.WriteRune('\n')
	b.WriteString(colors.BlurStyle.Render("Presiona 'enter' para iniciar una conversación"))
	return b.String()
}
