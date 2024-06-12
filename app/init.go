package app

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
)

func InitPath() {

	ex, err := os.Executable()
	if err != nil {
		fmt.Printf("Error while finding Schmierblatt-Path: %v", err)
		os.Exit(1)
	}

	path = filepath.Dir(ex)

}

func initTextarea() textarea.Model {

	t := textarea.New()
	t.CharLimit = 0
	t.Focus()
	t.SetValue(readFile())

	return t

}

func initTextinput() textinput.Model {

	t := textinput.New()
	t.Blur()
	t.Prompt = ": "

	return t

}

func InitModel() *Model {

	return &Model{
		schmierblatt: initTextarea(),
		commandline:  initTextinput(),
		focus: map[string]bool{
			"schmierblatt": true,
			"commandline":  false,
			"global":       false,
		},
	}

}
