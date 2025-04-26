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
	if m.focus["commandline"] {
		return lipgloss.JoinVertical(
			lipgloss.Left,
			borderStyle.Render(m.schmierblatt.View()),
			commandlineStyle.Width(m.width-2).Render(m.commandline.View()),
		)
	}
	if m.focus["filemenu"] {
		return lipgloss.JoinHorizontal(
			lipgloss.Top,
			filemenuStyle.Render(m.filemenu.View()),
			borderStyle.Render(m.schmierblatt.View()),
		)
	}
	return borderStyle.Height(m.height - 2).Render(m.schmierblatt.View())
}
