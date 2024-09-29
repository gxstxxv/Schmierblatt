package main

import (
	"fmt"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gxstxxv/schmierblatt/internal/app"
	"github.com/gxstxxv/schmierblatt/internal/logger"
)

func main() {
	// Get the executable's directory
	exePath, err := os.Executable()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get executable path: %v\n", err)
		os.Exit(1)
	}
	exeDir := filepath.Dir(exePath)

	// Change the working directory to the executable's directory
	if err := os.Chdir(exeDir); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to change working directory: %v\n", err)
		os.Exit(1)
	}

	// Initialize logger with absolute path
	logPath := filepath.Join(exeDir, "schmierblatt.log")
	if err := logger.Init(logPath); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize logger: %v\n", err)
		os.Exit(1)
	}

	// Load config with absolute path
	configPath := filepath.Join(exeDir, "config.yaml")
	if err := app.LoadConfig(configPath); err != nil {
		logger.Error("Failed to load configuration", "error", err)
		os.Exit(1)
	}

	p := tea.NewProgram(
		app.InitModel(),
		tea.WithAltScreen(),
	)

	if _, err := p.Run(); err != nil {
		logger.Error("Error while running program", "error", err)
		os.Exit(1)
	}
}
