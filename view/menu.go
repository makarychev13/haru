package view

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/makarychev13/haru/style"
)

const (
	header = `██╗  ██╗ █████╗ ██████╗ ██╗   ██╗
██║  ██║██╔══██╗██╔══██╗██║   ██║
███████║███████║██████╔╝██║   ██║
██╔══██║██╔══██║██╔══██╗██║   ██║
██║  ██║██║  ██║██║  ██║╚██████╔╝
╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═╝ ╚═════╝`
)

var (
	styleWrap = lipgloss.
		NewStyle().
		Bold(true).
		Align(lipgloss.Center).
		Padding(2)
)

type Menu struct {
	items  []string
	cursor int
}

func NewMenu() Menu {
	return Menu{
		items: []string{
			"Открыть хранилище",
			"Добавить хранилище",
			"Настройки",
			"Выход",
		},
		cursor: 0,
	}
}

func (m Menu) Init() tea.Cmd {
	return nil
}

func (m Menu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		case tea.KeyCtrlC.String():
			return m, tea.Quit
		case tea.KeyDown.String():
			m.cursor++
			if m.cursor == len(m.items) {
				m.cursor = 0
			}
		case tea.KeyUp.String():
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(m.items) - 1
			}
		}
	}

	return m, nil
}

func (m Menu) View() string {
	var sb strings.Builder

	sb.WriteString(header)
	sb.WriteString("\n\n")

	for i, v := range m.items {
		var b string

		if i == int(m.cursor) {
			b = style.SelectedButton.Render(v)
		} else {
			b = v
		}

		sb.WriteString(b)

		if i != len(m.items)-1 {
			sb.WriteString("\n")
		}
	}

	return styleWrap.Render(sb.String())
}
