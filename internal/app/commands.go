package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/gxstxxv/schmierblatt/internal/logger"
)

func (m *Model) handleCommand(value string) tea.Cmd {
	switch value {
	case "w":
		if err := writeFile(m.schmierblatt.Value(), m.files[m.selectedFileIndex]); err != nil {
			// m.commandline.Placeholder = "Error saving Schmierblatt!"
			logger.Error("Failed to save file", "error", err)
		} else {
			// m.commandline.Placeholder = "Schmierblatt has been saved!"
			m.filemenu = initList(m.files, m.selectedFileIndex)
			logger.Info("File saved successfully", "file", m.files[m.selectedFileIndex])
		}
	case "wq":
		if err := writeFile(m.schmierblatt.Value(), m.files[m.selectedFileIndex]); err != nil {
			// m.commandline.Placeholder = "Error saving Schmierblatt!"
			logger.Error("Failed to save file", "error", err)
			return nil
		}
		logger.Info("File saved successfully, quitting", "file", m.files[m.selectedFileIndex])
		return tea.Quit
	case "q":
		logger.Info("Quitting application")
		return tea.Quit
	default:
		logger.Info("Unknown command", "command", value)
	}
	return nil
}
