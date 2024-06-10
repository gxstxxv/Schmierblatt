package main

import (
	tea "github.com/charmbracelet/bubbletea"
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
