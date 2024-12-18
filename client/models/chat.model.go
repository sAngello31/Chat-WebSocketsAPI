package models

import (
	"client_websockets/colors"
	modeldata "client_websockets/model_data"
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

type Message struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Content string `json:"content"`
}

type ChatModel struct {
	Conn     *websocket.Conn
	Viewport viewport.Model
	TextArea textarea.Model
	Users    map[string]*modeldata.User
	Message  []string
}

func InitChatModel(userA, userB *modeldata.User) ChatModel {
	vp := viewport.New(width, heightViewport)
	vp.SetContent("Se inicio un nuevo chat.")

	users := make(map[string]*modeldata.User, 2)
	users["From"] = userA
	users["To"] = userB

	return ChatModel{
		Viewport: vp,
		TextArea: makeTextAreaModel(),
		Message:  []string{},
		Users:    users,
		Conn:     services.ConnectChat(userA.Username, userB.Username),
	}
}

func (m ChatModel) Init() tea.Cmd { return m.listenMessage() }

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
	case Message:
		m.renderMsg(&msg)
	}

	return m, m.listenMessage()
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

func makeTextAreaModel() textarea.Model {
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
	return ta
}

func (m *ChatModel) renderMsg(msg *Message) {
	var from_user string
	if m.Users["From"].Username == msg.From {
		from_user = "You: "
	} else {
		from_user = msg.From + ": "
	}
	m.Message = append(m.Message, colors.FocusedStyle.Render(from_user)+msg.Content)
	m.Viewport.SetContent(strings.Join(m.Message, "\n"))
	m.TextArea.Reset()
	m.Viewport.GotoBottom()
}

func (m *ChatModel) sendMsg(msg string) {
	msgJSON := Message{
		From:    m.Users["From"].Username,
		To:      m.Users["To"].Username,
		Content: msg,
	}
	m.Conn.WriteJSON(msgJSON)
}

func (m *ChatModel) listenMessage() tea.Cmd {
	return func() tea.Msg {
		var msg Message
		err := m.Conn.ReadJSON(&msg)
		if err != nil {
			m.Conn.Close()
		}
		return msg
	}
}
