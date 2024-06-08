package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {

	init_path()

	p := tea.NewProgram(
		init_model(), 
		tea.WithAltScreen(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error while starting Program: %v", err)
		os.Exit(1)
	}

}

