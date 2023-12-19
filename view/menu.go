package view

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/makarychev13/haru/style"
)

const (
	version = "v0.0.1"
	header  = `██╗  ██╗ █████╗ ██████╗ ██╗   ██╗
██║  ██║██╔══██╗██╔══██╗██║   ██║
███████║███████║██████╔╝██║   ██║
██╔══██║██╔══██║██╔══██╗██║   ██║
██║  ██║██║  ██║██║  ██║╚██████╔╝
╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═╝ ╚═════╝`
)

var (
	styleWrap = lipgloss.
			NewStyle().
			Align(lipgloss.Center).
			Padding(2)

	styleChoice = lipgloss.
			NewStyle().

		//Padding(1). //TODO: WTF???
		Align(lipgloss.Left)

	styleVersion = lipgloss.
			NewStyle().
			Foreground(lipgloss.Color("#9B9B9B"))
)

type Menu struct {
	items  []string
	cursor int
}

func NewMenu() Menu {
	return Menu{
		items: []string{
			"Новое хранилище",
			"Открыть хранилище",
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

	var sbChoice strings.Builder

	for i, v := range m.items {
		var b string

		if i == int(m.cursor) {
			b = style.SelectedButton.Render(v)
		} else {
			b = v
		}

		sbChoice.WriteString(b)

		if i != len(m.items)-1 {
			sbChoice.WriteString("\n")
		}
	}

	s := sb.String() +
		styleChoice.Render(sbChoice.String()) +
		"\n\n" +
		styleVersion.Render(version)

	return styleWrap.Render(s)
}
