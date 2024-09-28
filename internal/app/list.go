package app

type file struct {
	title string
	desc  string
}

func (f file) Title() string       { return f.title }
func (f file) Description() string { return f.desc }
func (f file) FilterValue() string { return f.title }
