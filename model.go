package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

var border_style = lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#616161"))

type model struct {

	schmierblatt textarea.Model
	commandline textinput.Model
	focus map[string]bool

}

func (m model) Init() tea.Cmd {

	return nil

}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch {

		case m.focus["schmierblatt"]:
			cmd = m.handle_schmierblatt_input(msg)

		case m.focus["commandline"]:
			cmd = m.handle_commandline_input(msg)

		case m.focus["global"]:
			cmd = m.handle_global_input(msg)

		}

	case tea.WindowSizeMsg:
		m.handle_window_input(msg.Width, msg.Height)
	
	}

	return m, cmd

}

func (m model) View() string {

	return lipgloss.JoinVertical(
		lipgloss.Center,
		border_style.Render(m.schmierblatt.View()),
		border_style.Render(m.commandline.View()),
	)

}

func (m *model) handle_global_input(msg tea.KeyMsg) tea.Cmd {

	switch {
		
	case key.Matches(msg, keys.colon):
		m.change_focus_to("commandline")

	case key.Matches(msg, keys.insert):
		m.change_focus_to("schmierblatt")

	}

	return nil

}

func (m *model) handle_schmierblatt_input(msg tea.KeyMsg) tea.Cmd {
	
	var cmd tea.Cmd

	switch {

	case key.Matches(msg, keys.esc):
		m.change_focus_to("global")

	}

	m.schmierblatt, cmd = m.schmierblatt.Update(msg)

	return cmd

}

func (m *model) handle_commandline_input(msg tea.KeyMsg) tea.Cmd {

	var cmd tea.Cmd
	
	switch {
		
	case key.Matches(msg, keys.esc):
		m.reset_commandline()
		m.change_focus_to("global")

	case key.Matches(msg, keys.enter):
		value := m.commandline.Value()
		m.reset_commandline()

		if cmd = m.handle_command(value); cmd != nil {
			return cmd
		}

		m.change_focus_to("global")

	}

	m.commandline, cmd = m.commandline.Update(msg)

	return cmd

}

func (m *model) handle_command(value string) tea.Cmd {

	switch value {

	case "w":
		write_file(m.schmierblatt.Value())
		m.commandline.Placeholder = "Schmierblatt has been saved!"

	case "wq":
		write_file(m.schmierblatt.Value())
		return tea.Quit

	case "q":
		return tea.Quit

	}

	return nil

}

func (m *model) handle_window_input(width, height int) {

	m.set_schmierblatt_size(width-2, height-5)
	m.set_commandline_size(width-5)

}

func (m *model) change_focus_to(name string) {

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

func (m *model) reset_commandline() {

	m.commandline.Placeholder = ""
	m.commandline.Reset()

}

func (m *model) set_schmierblatt_size(width, height int) {

	m.schmierblatt.SetHeight(height)
	m.schmierblatt.SetWidth(width)

}

func (m *model) set_commandline_size(width int) {

	m.commandline.Width = width
	m.commandline.CharLimit = width

}
