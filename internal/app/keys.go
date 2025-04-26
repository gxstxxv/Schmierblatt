package app

import (
	"github.com/charmbracelet/bubbles/key"
)

type KeyMap struct {
	Esc    key.Binding
	Quit   key.Binding
	Colon  key.Binding
	Enter  key.Binding
	Insert key.Binding
	Tab    key.Binding
	Up     key.Binding
	Down   key.Binding
	Right  key.Binding
	Open   key.Binding
}

var Keys = KeyMap{
	Esc: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "Exit current mode"),
	),
	Colon: key.NewBinding(
		key.WithKeys(":"),
		key.WithHelp(":", "Enter command mode"),
	),
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "Confirm"),
	),
	Open: key.NewBinding(
		key.WithKeys("o"),
		key.WithHelp("o", "Open"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("ctrl+c", "Quit"),
	),
	Insert: key.NewBinding(
		key.WithKeys("i"),
		key.WithHelp("i", "Enter insert mode"),
	),
	Tab: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "Switch focus"),
	),
	Up: key.NewBinding(
		key.WithKeys("k", "up"),
		key.WithHelp("k/↑", "Move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("j", "down"),
		key.WithHelp("j/↓", "Move down"),
	),
	Right: key.NewBinding(
		key.WithKeys("l", "right"),
		key.WithHelp("l/→", "Move right"),
	),
}
