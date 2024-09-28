package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/gxstxxv/schmierblatt/internal/logger"
)

func (m *Model) updateModelSizes() {
	var i, j int
	if m.focus["filemenu"] {
		i, j = FilemenuWidth, CommandlineHeight
	}
	m.setSchmierblattSize(m.width-2-i, m.height-5+j)
	m.setCommandlineSize(m.width - 5 - i)
	m.setFilemenuSize(m.width-1, m.height-2)
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

func (m *Model) changeSchmierblattValue(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd

	m.openFileIndex = m.selectedFileIndex
	m.filemenu.Select(m.selectedFileIndex)
	content, err := readFile(m.files[m.selectedFileIndex])
	if err != nil {
		logger.Error("Failed to read file", "file", m.files[m.selectedFileIndex], "error", err)
		return nil
	}
	m.schmierblatt.SetValue(content)
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
