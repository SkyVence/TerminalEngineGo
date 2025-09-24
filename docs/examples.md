# Examples

This document provides code examples for common use cases with TerminalEngineGo.

## Basic Counter Application

A simple counter that increments and decrements with key presses:

```go
package main

import (
    "fmt"
    "github.com/skyvence/TerminalEngineGo"
)

type counterModel struct {
    count int
}

func (m counterModel) Init() engine.Msg {
    return nil
}

func (m counterModel) Update(msg engine.Msg) (engine.Model, engine.Cmd) {
    switch msg := msg.(type) {
    case engine.KeyMsg:
        switch msg.Rune {
        case 'q':
            return m, func() engine.Msg { return engine.Quit() }
        case '+', '=':
            m.count++
        case '-':
            m.count--
        case 'r':
            m.count = 0
        }
    }
    return m, nil
}

func (m counterModel) View() string {
    return fmt.Sprintf(`
Counter: %d

Controls:
  +/= : Increment
  -   : Decrement  
  r   : Reset
  q   : Quit
`, m.count)
}

func main() {
    p := engine.NewProgram(counterModel{})
    p.Run()
}
```

## Timer Application

An application that updates every second:

```go
package main

import (
    "fmt"
    "time"
    "github.com/skyvence/TerminalEngineGo"
)

type timerModel struct {
    startTime time.Time
    current   time.Time
}

func (m timerModel) Init() engine.Msg {
    m.startTime = time.Now()
    m.current = m.startTime
    return engine.TickNow()
}

func (m timerModel) Update(msg engine.Msg) (engine.Model, engine.Cmd) {
    switch msg := msg.(type) {
    case engine.KeyMsg:
        if msg.Rune == 'q' {
            return m, func() engine.Msg { return engine.Quit() }
        }
    case engine.TickMsg:
        m.current = msg.Time
        return m, engine.Tick(time.Second)
    }
    return m, nil
}

func (m timerModel) View() string {
    elapsed := m.current.Sub(m.startTime)
    return fmt.Sprintf(`
Timer: %s
Current Time: %s

Press 'q' to quit
`, elapsed.Truncate(time.Second), m.current.Format("15:04:05"))
}

func main() {
    p := engine.NewProgram(timerModel{})
    p.Run()
}
```

## Simple Animation

Using the built-in Animation type:

```go
package main

import (
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
    p.Run()
}
```

## Navigation Menu

A simple menu with arrow key navigation:

```go
package main

import (
    "fmt"
    "github.com/skyvence/TerminalEngineGo"
)

type menuModel struct {
    options  []string
    selected int
}

func (m menuModel) Init() engine.Msg {
    m.options = []string{"Option 1", "Option 2", "Option 3", "Quit"}
    return nil
}

func (m menuModel) Update(msg engine.Msg) (engine.Model, engine.Cmd) {
    switch msg := msg.(type) {
    case engine.KeyMsg:
        switch msg.Rune {
        case 'q':
            return m, func() engine.Msg { return engine.Quit() }
        case '↑':
            if m.selected > 0 {
                m.selected--
            }
        case '↓':
            if m.selected < len(m.options)-1 {
                m.selected++
            }
        case '\r', '\n': // Enter key
            if m.selected == len(m.options)-1 { // Quit option
                return m, func() engine.Msg { return engine.Quit() }
            }
            // Handle other selections here
        }
    }
    return m, nil
}

func (m menuModel) View() string {
    s := "Select an option:\n\n"
    
    for i, option := range m.options {
        cursor := " "
        if i == m.selected {
            cursor = ">"
        }
        s += fmt.Sprintf("%s %s\n", cursor, option)
    }
    
    s += "\nUse ↑/↓ to navigate, Enter to select, q to quit"
    return s
}

func main() {
    p := engine.NewProgram(menuModel{})
    p.Run()
}
```

## Full Screen Application

Using alternate screen mode:

```go
package main

import (
    "fmt"
    "strings"
    "github.com/skyvence/TerminalEngineGo"
)

type fullScreenModel struct {
    width, height int
    message       string
}

func (m fullScreenModel) Init() engine.Msg {
    m.message = "Full Screen Mode!"
    return nil
}

func (m fullScreenModel) Update(msg engine.Msg) (engine.Model, engine.Cmd) {
    switch msg := msg.(type) {
    case engine.KeyMsg:
        if msg.Rune == 'q' {
            return m, func() engine.Msg { return engine.Quit() }
        }
    case engine.SizeMsg:
        m.width = msg.Width
        m.height = msg.Height
    }
    return m, nil
}

func (m fullScreenModel) View() string {
    // Center the message
    padding := strings.Repeat(" ", (m.width-len(m.message))/2)
    centered := padding + m.message
    
    // Create vertical padding
    verticalPadding := m.height / 2
    content := strings.Repeat("\n", verticalPadding) + centered
    
    // Add footer
    footer := fmt.Sprintf("\n\nTerminal size: %dx%d\nPress 'q' to quit", m.width, m.height)
    
    return content + footer
}

func main() {
    p := engine.NewProgram(fullScreenModel{}, engine.WithAltScreen())
    p.Run()
}
```

## Game Example

A simple game with player movement:

```go
package main

import (
    "fmt"
    "github.com/skyvence/TerminalEngineGo"
)

type gameModel struct {
    playerX, playerY int
    width, height    int
    score           int
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
        m.width = msg.Width - 2  // Account for borders
        m.height = msg.Height - 5 // Account for UI
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
    p := engine.NewProgram(gameModel{})
    p.Run()
}
```

These examples demonstrate the core patterns and capabilities of TerminalEngineGo. You can find complete, runnable versions in the `examples/` directory of the repository.