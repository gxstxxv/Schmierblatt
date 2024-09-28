package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gxstxxv/schmierblatt/internal/app"
	"github.com/gxstxxv/schmierblatt/internal/logger"
)

func main() {
	if err := logger.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize logger: %v\n", err)
		os.Exit(1)
	}

	if err := app.LoadConfig(); err != nil {
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
