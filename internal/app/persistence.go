package app

import (
	"os"

	"github.com/gxstxxv/schmierblatt/internal/logger"
)

func saveLastOpenedFile(fileName string) {
	err := os.WriteFile(config.App.LastOpenedFilePath, []byte(fileName), 0644)
	if err != nil {
		logger.Error("Failed to save last opened file", "error", err)
	}
}

func getLastOpenedFile() string {
	content, err := os.ReadFile(config.App.LastOpenedFilePath)
	if err != nil {
		if !os.IsNotExist(err) {
			logger.Error("Failed to read last opened file", "error", err)
		}
		return ""
	}
	return string(content)
}

func getLastOpenedFileIndex(lastOpenedFile string, files []string) int {
	index := 0
	if lastOpenedFile != "" {
		for i, file := range files {
			if file == lastOpenedFile {
				index = i
				break
			}
		}
	}
	return index
}
