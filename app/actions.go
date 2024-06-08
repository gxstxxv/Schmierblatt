package main

import (
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
)

func change_focus_to(m model, name string) model {

	for i := range m.focus {
		m.focus[i] = false
	}

	switch name {

	case "schmierblatt":
		m.commandline.Blur()
		m.schmierblatt.Focus()

	case "commandline":
		m.schmierblatt.Blur()
		m.commandline.Focus()

	}

	m.focus[name] = true

	return m

}

func reset_textinput(t textinput.Model) textinput.Model {

	t.Placeholder = ""
	t.Reset()

	return t

}

func set_textarea_size(t textarea.Model, width, height int) textarea.Model {

	t.SetHeight(height)
	t.SetWidth(width)

	return t

}

func set_textinput_size(t textinput.Model, width int) textinput.Model {

	t.Width = width
	t.CharLimit = width

	return t

}
