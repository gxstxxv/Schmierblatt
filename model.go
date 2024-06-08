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
		if m.focus["schmierblatt"] {

			switch {
				
			case key.Matches(msg, keys.esc):
				m.schmierblatt.Blur()
				m.focus["schmierblatt"] = false

			}

			m.schmierblatt, cmd = m.schmierblatt.Update(msg)

		} else if m.focus["commandline"] {

			switch {
				
			case key.Matches(msg, keys.esc):
				m.commandline = reset_textinput(m.commandline)
				m = change_focus_to(m, "schmierblatt")

			case key.Matches(msg, keys.enter):
				value := m.commandline.Value()
				m.commandline = reset_textinput(m.commandline)

				switch value {

				case "w":
					write_file(m.schmierblatt.Value())
					m.commandline.Placeholder = "Schmierblatt has been saved!"

				case "wq":
					write_file(m.schmierblatt.Value())
					return m, tea.Quit

				case "q":
					return m, tea.Quit

				}

				m = change_focus_to(m, "schmierblatt")

			}

			m.commandline, cmd = m.commandline.Update(msg)

		} else if !m.focus["schmierblatt"] {
			
			switch {
				
			case key.Matches(msg, keys.colon):
				m = change_focus_to(m, "commandline")

			case key.Matches(msg, keys.insert):
				m.schmierblatt.Focus()
				m.focus["schmierblatt"] = true

			}

			m.schmierblatt, cmd = m.schmierblatt.Update(msg)

		} 

		switch {

		case key.Matches(msg, keys.quit):
			return m, tea.Quit

		}
	
	case tea.WindowSizeMsg:
		width := msg.Width
		height := msg.Height

		m.schmierblatt = set_textarea_size(m.schmierblatt, width-2, height-5)
		m.commandline = set_textinput_size(m.commandline, width-5)
	
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
