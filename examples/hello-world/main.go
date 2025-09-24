package main

import (
	"fmt"
	"os"

	"github.com/skyvence/TerminalEngineGo"
)

type model struct {
	content string
}

func (m model) Init() engine.Msg {
	return nil
}

func (m model) Update(msg engine.Msg) (engine.Model, engine.Cmd) {
	switch msg := msg.(type) {
	case engine.KeyMsg:
		if msg.Rune == 'q' {
			return m, func() engine.Msg { return engine.Quit() }
		}
	}
	return m, nil
}

func (m model) View() string {
	return "Hello, Terminal Engine!\nPress 'q' to quit.\n"
}

func main() {
	m := model{content: "Hello World"}
	p := engine.NewProgram(m)
	
	if err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}