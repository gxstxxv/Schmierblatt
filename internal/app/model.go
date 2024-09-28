package app

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
)

const (
	FilemenuWidth     = 22
	CommandlineHeight = 3
)

type Model struct {
	schmierblatt      textarea.Model
	commandline       textinput.Model
	filemenu          list.Model
	files             []string
	selectedFileIndex int
	openFileIndex     int
	focus             map[string]bool
	width, height     int
}

func InitModel() *Model {
	files, err := readDir()
	if err != nil {
		logger.Error("Error reading directory", "error", err)
		return nil
	}

	return &Model{
		schmierblatt:      initTextarea(),
		commandline:       initTextinput(),
		filemenu:          initList(),
		files:             files,
		selectedFileIndex: 0,
		openFileIndex:     0,
		focus: map[string]bool{
			"schmierblatt": false,
			"commandline":  false,
			"filemenu":     false,
			"global":       true,
		},
	}
}

func initTextarea() textarea.Model {
	t := textarea.New()
	t.CharLimit = 0
	t.Blur()
	files, err := readDir()
	if err != nil {
		logger.Error("Error reading directory", "error", err)
		return t
	}
	content, err := readFile(files[0])
	if err != nil {
		logger.Error("Error reading file", "error", err)
		return t
	}
	t.SetValue(content)
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

	files, err := readDir()
	if err != nil {
		logger.Error("Error reading directory", "error", err)
		return l
	}
	descs, err := readDesc()
	if err != nil {
		logger.Error("Error reading descriptions", "error", err)
		return l
	}

	items := make([]list.Item, 0, len(files))
	for i, f := range files {
		items = append(items, file{title: f, desc: descs[i]})
	}
	l.SetItems(items)
	return l
}
