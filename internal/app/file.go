package app

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gxstxxv/schmierblatt/internal/logger"
)

func readDir() ([]string, error) {
	entries, err := os.ReadDir(config.App.AssetsPath)
	if err != nil {
		return nil, fmt.Errorf("error reading directory: %w", err)
	}

	files := make([]string, 0, len(entries))
	for _, entry := range entries {
		files = append(files, entry.Name())
	}

	return files, nil
}

func readDesc() ([]string, error) {
	files, err := readDir()
	if err != nil {
		return nil, err
	}

	descs := make([]string, 0, len(files))
	descLength := FilemenuWidth - 6

	for _, file := range files {
		filePath := filepath.Join(config.App.AssetsPath, file)
		f, err := os.Open(filePath)
		if err != nil {
			return nil, fmt.Errorf("error opening file %s: %w", file, err)
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		var desc string
		if scanner.Scan() {
			desc = scanner.Text()
		}

		if err := scanner.Err(); err != nil {
			return nil, fmt.Errorf("error reading file %s: %w", file, err)
		}

		if len(desc) > descLength {
			desc = desc[:descLength-3] + "..."
		}

		descs = append(descs, desc)
	}

	return descs, nil
}

func readFile(file string) (string, error) {
	data, err := os.ReadFile(filepath.Join(config.App.AssetsPath, file))
	if err != nil {
		return "", fmt.Errorf("error while reading file %s: %w", file, err)
	}
	return string(data), nil
}

func writeFile(data, fileName string) error {
	filePath := filepath.Join(config.App.AssetsPath, fileName)
	err := os.WriteFile(filePath, []byte(data), 0644)
	if err != nil {
		logger.Error("Failed to write file", "file", fileName, "error", err)
		return err
	}
	logger.Info("File written successfully", "file", fileName)
	return nil
}
