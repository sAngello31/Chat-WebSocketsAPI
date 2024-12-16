package models

import (
	"client_websockets/colors"
	"client_websockets/services"
	"strings"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/gorilla/websocket"
)

var (
	width          = 30
	heightTextArea = 3
	heightViewport = 5
)

type ChatModel struct {
	Conn     *websocket.Conn
	Viewport viewport.Model
	TextArea textarea.Model
	Message  []string
}

func InitChatModel(userA, userB string) ChatModel {
	ta := textarea.New()
	ta.Placeholder = "Enviar un mensaje..."
	ta.Focus()

	ta.Prompt = "┃ "
	ta.CharLimit = 280

	ta.SetHeight(heightTextArea)
	ta.SetWidth(width)

	ta.FocusedStyle.CursorLine = lipgloss.NewStyle()
	ta.ShowLineNumbers = false
	ta.KeyMap.InsertNewline.SetEnabled(false)

	vp := viewport.New(width, heightViewport)
	vp.SetContent("Se inicio un nuevo chat.")

	return ChatModel{
		Viewport: vp,
		TextArea: ta,
		Message:  []string{},
		Conn:     services.ConnectChat(userA, userB),
	}
}

func (m ChatModel) Init() tea.Cmd { return m.TextArea.Cursor.BlinkCmd() }

func (m ChatModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			m.Conn.Close()
			return InitMenuModel(), nil
		case "enter":
			v := m.TextArea.Value()
			if v == "" {
				return m, nil
			}
			m.sendMsg(v)
			return m, nil
		default:
			var cmd tea.Cmd
			m.TextArea, cmd = m.TextArea.Update(msg)
			return m, cmd
		}
	}

	return m, nil
}

func (m ChatModel) View() string {
	var b strings.Builder
	b.WriteString(m.Viewport.View())
	b.WriteRune('\n')
	b.WriteRune('\n')
	b.WriteString(m.TextArea.View())
	b.WriteRune('\n')
	b.WriteByte('\n')
	return b.String()
}

func (m *ChatModel) sendMsg(msg string) {
	m.Message = append(m.Message, colors.FocusedStyle.Render("You: ")+msg)
	m.Viewport.SetContent(strings.Join(m.Message, "\n"))
	m.TextArea.Reset()
	m.Viewport.GotoBottom()
	m.Conn.WriteMessage(websocket.TextMessage, []byte(msg))
}
