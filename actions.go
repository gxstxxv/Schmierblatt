package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/key"
)

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
