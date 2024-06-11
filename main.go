package main

import (
	"Schmierblatt/app"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {

	app.InitPath()

	p := tea.NewProgram(
		app.InitModel(),
		tea.WithAltScreen(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error while starting Program: %v", err)
		os.Exit(1)
	}

}
