package app

import (
	"github.com/charmbracelet/bubbles/key"
)

type KeyMap struct {
	esc    key.Binding
	quit   key.Binding
	colon  key.Binding
	enter  key.Binding
	insert key.Binding
}

var keys = KeyMap{

	esc: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "focus textarea"),
	),
	colon: key.NewBinding(
		key.WithKeys(":"),
		key.WithHelp(":", "focus commandline"),
	),
	enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "perform command"),
	),
	quit: key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("ctrl+c", "quit schmierblatt"),
	),
	insert: key.NewBinding(
		key.WithKeys("i"),
		key.WithHelp("i", "enter insertmode"),
	),
}
