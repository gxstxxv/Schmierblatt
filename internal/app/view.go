package app

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	borderStyle      = lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#616161"))
	filemenuStyle    = lipgloss.NewStyle().Margin(1, 2)
	commandlineStyle = lipgloss.NewStyle().Inherit(borderStyle)
)

func (m *Model) View() string {
	if !m.focus["filemenu"] {
		return lipgloss.JoinVertical(
			lipgloss.Left,
			borderStyle.Render(m.schmierblatt.View()),
			commandlineStyle.Width(m.width-2).Render(m.commandline.View()),
		)
	}
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		filemenuStyle.Render(m.filemenu.View()),
		borderStyle.Render(m.schmierblatt.View()),
	)
}
