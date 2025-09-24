package main

import (
	"fmt"
	"os"

	"github.com/skyvence/TerminalEngineGo"
)

type animationModel struct {
	animation engine.Animation
}

func (m animationModel) Init() engine.Msg {
	frames := []string{
		"( ˘ ³˘)♥",
		"( ˘ ³˘)♡",
		"( ˘ ³˘)♥",
		"( ˘ ³˘)♡",
	}
	m.animation = engine.NewAnimation(frames)
	return m.animation.Init()()
}

func (m animationModel) Update(msg engine.Msg) (engine.Model, engine.Cmd) {
	switch msg := msg.(type) {
	case engine.KeyMsg:
		if msg.Rune == 'q' {
			return m, func() engine.Msg { return engine.Quit() }
		}
	case engine.TickMsg:
		var cmd engine.Cmd
		m.animation, cmd = m.animation.Update(msg)
		return m, cmd
	}
	return m, nil
}

func (m animationModel) View() string {
	return fmt.Sprintf(`
%s

Animated heart! Press 'q' to quit.
`, m.animation.View())
}

func main() {
	p := engine.NewProgram(animationModel{})
	
	if err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}