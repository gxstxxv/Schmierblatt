package app

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/gxstxxv/schmierblatt/internal/logger"
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

func (m *Model) Init() tea.Cmd {
	return nil
}

func InitModel() *Model {
	files, err := readDir()
	if err != nil {
		logger.Error("Error reading directory", "error", err)
		files = []string{}
	}

	lastOpenedFile := getLastOpenedFile()
	lastOpenedFileIndex := getLastOpenedFileIndex(lastOpenedFile, files)

	m := &Model{
		schmierblatt:      initTextarea(files, lastOpenedFileIndex),
		commandline:       initTextinput(),
		filemenu:          initList(files),
		files:             files,
		selectedFileIndex: lastOpenedFileIndex,
		openFileIndex:     lastOpenedFileIndex,
		focus: map[string]bool{
			"schmierblatt": false,
			"commandline":  false,
			"filemenu":     false,
			"global":       true,
		},
	}

	return m
}

func initTextarea(files []string, lastOpenedFileIndex int) textarea.Model {
	t := textarea.New()
	t.CharLimit = 0
	t.Blur()
	if len(files) > 0 {
		content, err := readFile(files[lastOpenedFileIndex])
		if err != nil {
			logger.Error("Error reading file", "error", err)
		} else {
			t.SetValue(content)
		}
	}
	return t
}

func initTextinput() textinput.Model {
	t := textinput.New()
	t.Blur()
	t.Prompt = ": "
	return t
}

func initList(files []string) list.Model {
	l := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	l.SetShowTitle(false)
	l.SetShowHelp(false)
	l.SetFilteringEnabled(false)
	l.DisableQuitKeybindings()

	if len(files) > 0 {
		descs, err := readDesc()
		if err != nil {
			logger.Error("Error reading descriptions", "error", err)
			descs = make([]string, len(files)) // Use empty descriptions
		}

		items := make([]list.Item, 0, len(files))
		for i, f := range files {
			items = append(items, file{title: f, desc: descs[i]})
		}
		l.SetItems(items)
	}

	return l
}
