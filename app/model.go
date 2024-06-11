package app

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var borderStyle = lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#616161"))

type Model struct {
	schmierblatt textarea.Model
	commandline  textinput.Model
	focus        map[string]bool
}

func (m *Model) Init() tea.Cmd {

	return nil

}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch {

		case m.focus["schmierblatt"]:
			cmd = m.handleSchmierblattInput(msg)

		case m.focus["commandline"]:
			cmd = m.handleCommandlineInput(msg)

		case m.focus["global"]:
			cmd = m.handleGlobalInput(msg)

		}

	case tea.WindowSizeMsg:
		m.handleWindowInput(msg.Width, msg.Height)

	}

	return m, cmd

}

func (m *Model) View() string {

	return lipgloss.JoinVertical(
		lipgloss.Center,
		borderStyle.Render(m.schmierblatt.View()),
		borderStyle.Render(m.commandline.View()),
	)

}

func (m *Model) handleGlobalInput(msg tea.KeyMsg) tea.Cmd {

	switch {

	case key.Matches(msg, keys.colon):
		m.changeFocusTo("commandline")

	case key.Matches(msg, keys.insert):
		m.changeFocusTo("schmierblatt")

	}

	return nil

}

func (m *Model) handleSchmierblattInput(msg tea.KeyMsg) tea.Cmd {

	var cmd tea.Cmd

	switch {

	case key.Matches(msg, keys.esc):
		m.changeFocusTo("global")

	}

	m.schmierblatt, cmd = m.schmierblatt.Update(msg)

	return cmd

}

func (m *Model) handleCommandlineInput(msg tea.KeyMsg) tea.Cmd {

	var cmd tea.Cmd

	switch {

	case key.Matches(msg, keys.esc):
		m.resetCommandline()
		m.changeFocusTo("global")

	case key.Matches(msg, keys.enter):
		value := m.commandline.Value()
		m.resetCommandline()

		if cmd = m.handleCommand(value); cmd != nil {
			return cmd
		}

		m.changeFocusTo("global")

	}

	m.commandline, cmd = m.commandline.Update(msg)

	return cmd

}

func (m *Model) handleCommand(value string) tea.Cmd {

	switch value {

	case "w":
		writeFile(m.schmierblatt.Value())
		m.commandline.Placeholder = "Schmierblatt has been saved!"

	case "wq":
		writeFile(m.schmierblatt.Value())
		return tea.Quit

	case "q":
		return tea.Quit

	}

	return nil

}

func (m *Model) handleWindowInput(width, height int) {

	m.setSchmierblattSize(width-2, height-5)
	m.setCommandlineSize(width - 5)

}

func (m *Model) changeFocusTo(name string) {

	for i := range m.focus {
		m.focus[i] = false
	}

	switch name {

	case "schmierblatt":
		m.commandline.Blur()
		m.schmierblatt.Focus()

	case "commandline":
		m.schmierblatt.Blur()
		m.commandline.Focus()

	case "global":
		m.commandline.Blur()
		m.schmierblatt.Blur()

	}

	m.focus[name] = true

}

func (m *Model) resetCommandline() {

	m.commandline.Placeholder = ""
	m.commandline.Reset()

}

func (m *Model) setSchmierblattSize(width, height int) {

	m.schmierblatt.SetHeight(height)
	m.schmierblatt.SetWidth(width)

}

func (m *Model) setCommandlineSize(width int) {

	m.commandline.Width = width
	m.commandline.CharLimit = width

}
