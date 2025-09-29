package main

import (
	"fmt"
	"os"

	engine "github.com/skyvence/TerminalEngineGo"
)

type gameModel struct {
	playerX, playerY int
	width, height    int
	score            int
}

func (m gameModel) Init() engine.Msg {
	m.playerX, m.playerY = 5, 5
	m.width, m.height = 20, 10
	return nil
}

func (m gameModel) Update(msg engine.Msg) (engine.Model, engine.Cmd) {
	switch msg := msg.(type) {
	case engine.KeyMsg:
		switch msg.Rune {
		case 'q':
			return m, func() engine.Msg { return engine.Quit() }
		case '↑', 'w':
			if m.playerY > 0 {
				m.playerY--
				m.score++
			}
		case '↓', 's':
			if m.playerY < m.height-1 {
				m.playerY++
				m.score++
			}
		case '←', 'a':
			if m.playerX > 0 {
				m.playerX--
				m.score++
			}
		case '→', 'd':
			if m.playerX < m.width-1 {
				m.playerX++
				m.score++
			}
		}
	case engine.SizeMsg:
		if msg.Width > 2 && msg.Height > 5 {
			m.width = msg.Width - 2   // Account for borders
			m.height = msg.Height - 5 // Account for UI
		}
	}
	return m, nil
}

func (m gameModel) View() string {
	// Create game board
	board := make([][]rune, m.height)
	for i := range board {
		board[i] = make([]rune, m.width)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}

	// Place player
	if m.playerY < len(board) && m.playerX < len(board[0]) {
		board[m.playerY][m.playerX] = '@'
	}

	// Render board
	result := fmt.Sprintf("Score: %d\n\n", m.score)
	for _, row := range board {
		result += string(row) + "\n"
	}

	result += "\nUse WASD or arrow keys to move, q to quit"
	return result
}

func main() {
	p := engine.NewProgram(engine.Wrap(gameModel{}), engine.WithAltScreen())

	if err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
