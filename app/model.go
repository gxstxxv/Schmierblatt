package app

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	schmierblatt        textarea.Model
	commandline         textinput.Model
	filemenu            list.Model
	files               []string
	selected_file_index int
	open_file_index     int
	focus               map[string]bool
	width, height       int
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

		case m.focus["filemenu"]:
			cmd = m.handleFilemenuInput(msg)

		case m.focus["global"]:
			cmd = m.handleGlobalInput(msg)

		}

	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
		m.updateModelSizes(m.width, m.height)

	}

	return m, cmd

}

var borderStyle = lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#616161"))
var filemenuStyle = lipgloss.NewStyle().Margin(1, 2)

func (m *Model) View() string {

	var view string

	switch {

	case !m.focus["filemenu"]:
		view = lipgloss.JoinVertical(
			lipgloss.Left,
			borderStyle.Render(m.schmierblatt.View()),
			borderStyle.Render(m.commandline.View()),
		)

	case m.focus["filemenu"]:
		view = lipgloss.JoinHorizontal(
			lipgloss.Top,
			filemenuStyle.Render(m.filemenu.View()),
			borderStyle.Render(m.schmierblatt.View()),
		)

	}

	return view

}

func (m *Model) handleGlobalInput(msg tea.KeyMsg) tea.Cmd {

	switch {

	case key.Matches(msg, keys.colon):
		m.changeFocusTo("commandline")

	case key.Matches(msg, keys.insert):
		m.changeFocusTo("schmierblatt")

	case key.Matches(msg, keys.tab):
		m.changeFocusTo("filemenu")
		m.updateModelSizes(m.width, m.height)

	}

	return nil

}

func (m *Model) handleFilemenuInput(msg tea.KeyMsg) tea.Cmd {

	var cmd tea.Cmd

	switch {

	case key.Matches(msg, keys.tab):
		m.changeFocusTo("global")
		m.updateModelSizes(m.width, m.height)

	case key.Matches(msg, keys.up):
		if m.selected_file_index <= 0 {
			return nil
		}
		m.selected_file_index -= 1

	case key.Matches(msg, keys.down):
		if m.selected_file_index >= len(m.files)-1 {
			return nil
		}
		m.selected_file_index += 1

	case key.Matches(msg, keys.right):
		if m.open_file_index != m.selected_file_index {
			return m.changeSchmierblattValue(msg)
		}
		m.changeFocusTo("global")
		m.updateModelSizes(m.width, m.height)

	}

	m.filemenu, cmd = m.filemenu.Update(msg)

	return cmd

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
		writeFile(m.schmierblatt.Value(), m.files[m.selected_file_index])
		m.commandline.Placeholder = "Schmierblatt has been saved!"

	case "wq":
		writeFile(m.schmierblatt.Value(), m.files[m.selected_file_index])
		return tea.Quit

	case "q":
		return tea.Quit

	}

	return nil

}

const FilemenuWidth = 22
const CommandlineHeight = 3

func (m *Model) updateModelSizes(width, height int) {

	var i, j int

	if m.focus["filemenu"] {
		i = FilemenuWidth
		j = CommandlineHeight
	}

	m.setSchmierblattSize(width-2-i, height-5+j)
	m.setCommandlineSize(width - 5 - i)
	m.setFilemenuSize(width-1, height-2)

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

	case "global", "filemenu":
		m.commandline.Blur()
		m.schmierblatt.Blur()

	}

	m.focus[name] = true

}

func (m *Model) resetCommandline() {

	m.commandline.Placeholder = ""
	m.commandline.Reset()

}

func (m *Model) changeSchmierblattValue(msg tea.KeyMsg) tea.Cmd {

	var cmd tea.Cmd

	m.open_file_index = m.selected_file_index
	m.filemenu.Select(m.selected_file_index)
	m.schmierblatt.SetValue(readFile(m.files[m.selected_file_index]))
	m.schmierblatt, cmd = m.schmierblatt.Update(msg)

	return cmd

}

func (m *Model) setSchmierblattSize(width, height int) {

	m.schmierblatt.SetHeight(height)
	m.schmierblatt.SetWidth(width)

}

func (m *Model) setCommandlineSize(width int) {

	m.commandline.Width = width
	m.commandline.CharLimit = width

}

func (m *Model) setFilemenuSize(width, height int) {

	m.filemenu.SetSize(width, height)

}
