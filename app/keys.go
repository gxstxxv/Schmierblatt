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
	tab    key.Binding
	up     key.Binding
	down   key.Binding
	right  key.Binding
}

var keys = KeyMap{

	esc: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", ""),
	),
	colon: key.NewBinding(
		key.WithKeys(":"),
		key.WithHelp(":", ""),
	),
	enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", ""),
	),
	quit: key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("ctrl+c", ""),
	),
	insert: key.NewBinding(
		key.WithKeys("i"),
		key.WithHelp("i", ""),
	),
	tab: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", ""),
	),
	up: key.NewBinding(
		key.WithKeys("k", "Up"),
		key.WithHelp("k", ""),
	),
	down: key.NewBinding(
		key.WithKeys("j", "Down"),
		key.WithHelp("j", ""),
	),
	right: key.NewBinding(
		key.WithKeys("l", "Right"),
		key.WithHelp("l", ""),
	),
}
