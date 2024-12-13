package models

import (
	"client_websockets/colors"
	modeldata "client_websockets/model_data"
	"client_websockets/services"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type MenuModel struct {
	User   *modeldata.User
	List   *[]modeldata.User
	Cursor int
}

func InitMenuModel() MenuModel {
	return MenuModel{
		User: services.GetUserData(),
		List: services.GetAllUsers(),
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
		case "up":
			if m.Cursor > 0 {
				m.Cursor--
			}
		case "down":
			if m.Cursor < len(*m.List)-1 {
				m.Cursor++
			}
		}
	}
	return m, nil
}

func (m MenuModel) View() string {
	var b strings.Builder
	b.WriteString("Bienvenido: " + m.User.Username + "\n")
	b.WriteString("Tu numero de contacto: # " + m.User.Description() + "\n\n")
	b.WriteString(colors.BlurStyle.Render("Presiona 'esc' para cerrar sesión"))
	b.WriteRune('\n')
	b.WriteString(colors.BlurStyle.Render("Presiona 'enter' para iniciar una conversación"))
	b.WriteRune('\n')
	b.WriteString("Seleccione con quien quiere hablar:")
	b.WriteRune('\n')

	for i, user := range *m.List {
		cursor := " "
		if i == m.Cursor {
			cursor = ">"
			b.WriteString(colors.FocusedStyle.Render(cursor))
			b.WriteString(colors.FocusedStyle.Render(" " + createLineMenu(&user)))
		} else {
			b.WriteString(cursor)
			b.WriteString(" " + createLineMenu(&user))
		}

		b.WriteRune('\n')
	}
	b.WriteRune('\n')
	return b.String()
}

func createLineMenu(user *modeldata.User) string {
	time_spaces := 10 - len(user.Title())
	spaces := strings.Repeat(" ", time_spaces)
	return user.Title() + spaces + " | " + user.Description()
}
