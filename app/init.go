package app

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/charmbracelet/bubbles/list"

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
	path = "/Users/gg1/Documents/Code/Go/Schmierblatt/"

}

func initTextarea() textarea.Model {

	t := textarea.New()
	t.CharLimit = 0
	t.Blur()
	files := readDir()
	t.SetValue(readFile(files[0]))

	return t

}

func initTextinput() textinput.Model {

	t := textinput.New()
	t.Blur()
	t.Prompt = ": "

	return t

}

func initList() list.Model {

	l := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	l.SetShowTitle(false)
	l.SetShowHelp(false)
	l.SetFilteringEnabled(false)
	l.DisableQuitKeybindings()

	files := readDir()
	descs := readDesc()

	items := []list.Item{}

	for i, f := range files {
		item := file{
			title: f,
			desc:  descs[i],
		}
		items = append(items, item)
	}

	l.SetItems(items)

	return l

}

func InitModel() *Model {

	return &Model{
		schmierblatt: initTextarea(),
		commandline:  initTextinput(),
		filemenu:     initList(),
		files:        readDir(),
		selected_file_index:   0,
		open_file_index: 0,
		focus: map[string]bool{
			"schmierblatt": false,
			"commandline":  false,
			"filemenu":     false,
			"global":       true,
		},
	}

}
