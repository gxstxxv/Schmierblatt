package main

import (
	"os"
  "path/filepath"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
)

func init_path() {

	ex, err := os.Executable()
	if err != nil {
			panic(err)
	}

	path = filepath.Dir(ex) 

}

func init_textarea() textarea.Model {

	t := textarea.New()
	t.Focus()
	t.SetValue(read_file())

	return t

}

func init_textinput() textinput.Model {

	t := textinput.New()
	t.Blur()
	t.Prompt = ": "
	
	return t

}

func init_model() model {

	return model{
 		schmierblatt: init_textarea(),
		commandline: init_textinput(),
		focus: map[string]bool{
			"schmierblatt": true,
			"commandline": false,
		},
	}

}
