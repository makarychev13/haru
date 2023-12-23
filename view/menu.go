package view

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/makarychev13/haru/pkg/cursor"
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
			Align(lipgloss.Left)

	styleVersion = lipgloss.
			NewStyle().
			Foreground(lipgloss.Color("#9B9B9B"))
)

type Menu struct {
	cur    *cursor.Cursor[string]
	cursor int
}

func NewMenu() Menu {
	return Menu{
		cur: cursor.NewCursor[string](
			"New vault",
			"Open vault",
			"About",
		),
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
			m.cur.Inc()
		case tea.KeyUp.String():
			m.cur.Dec()
		}
	}

	return m, nil
}

func (m Menu) View() string {
	var sb strings.Builder

	sb.WriteString(header)
	sb.WriteString("\n\n")

	var sbChoice strings.Builder

	for i, v := range m.cur.AllValues() {
		var b string

		if i == m.cur.Index() {
			b = style.SelectedButton.Render(v)
		} else {
			b = style.UnselectedButton.Render(v)
		}

		sbChoice.WriteString(b)

		if i != len(m.cur.AllValues())-1 {
			sbChoice.WriteString("\n")
		}
	}

	sb.WriteString(styleChoice.Render(sbChoice.String()))
	sb.WriteString("\n\n")
	sb.WriteString(styleVersion.Render(version))

	return styleWrap.Render(sb.String())
}
