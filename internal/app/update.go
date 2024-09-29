package app

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

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
		m.updateModelSizes()
	}

	return m, cmd
}

func (m *Model) handleGlobalInput(msg tea.KeyMsg) tea.Cmd {
	switch {
	case key.Matches(msg, Keys.Colon):
		m.changeFocusTo("commandline")
	case key.Matches(msg, Keys.Insert):
		m.changeFocusTo("schmierblatt")
	case key.Matches(msg, Keys.Tab):
		m.changeFocusTo("filemenu")
		m.updateModelSizes()
	}
	return nil
}

func (m *Model) handleFilemenuInput(msg tea.KeyMsg) tea.Cmd {
	var cmd tea.Cmd

	switch {
	case key.Matches(msg, Keys.Tab):
		m.changeFocusTo("global")
		m.updateModelSizes()
	case key.Matches(msg, Keys.Up):
		if m.selectedFileIndex > 0 {
			m.selectedFileIndex--
		}
	case key.Matches(msg, Keys.Down):
		if m.selectedFileIndex < len(m.files)-1 {
			m.selectedFileIndex++
		}
	case key.Matches(msg, Keys.Right):
		if m.openFileIndex != m.selectedFileIndex {
			return m.changeSchmierblattValue(msg)
		}
		m.changeFocusTo("global")
		m.updateModelSizes()
	}

	m.filemenu, cmd = m.filemenu.Update(msg)
	return cmd
}

func (m *Model) handleSchmierblattInput(msg tea.KeyMsg) tea.Cmd {
	var cmd tea.Cmd

	switch {
	case key.Matches(msg, Keys.Esc):
		m.changeFocusTo("global")
	case key.Matches(msg, Keys.Tab):
		m.schmierblatt.InsertString("\t")
		return nil

	}

	m.schmierblatt, cmd = m.schmierblatt.Update(msg)
	return cmd
}

func (m *Model) handleCommandlineInput(msg tea.KeyMsg) tea.Cmd {
	var cmd tea.Cmd

	switch {
	case key.Matches(msg, Keys.Esc):
		m.resetCommandline()
		m.changeFocusTo("global")
	case key.Matches(msg, Keys.Enter):
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
