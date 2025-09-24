# Getting Started with TerminalEngineGo

## Introduction

TerminalEngineGo is a terminal game engine that follows the Model-View-Update (MVU) architecture pattern, similar to The Elm Architecture. This design makes it easy to build interactive terminal applications with predictable state management.

## Core Concepts

### Model
The Model represents your application's state. It should contain all the data your application needs to render and respond to user input.

### View
The View function takes your model and returns a string representation of what should be displayed on the terminal.

### Update
The Update function handles messages (events) and returns a new model and optionally a command to execute.

### Messages
Messages represent events in your application - user input, timer ticks, or custom events.

### Commands
Commands are functions that can produce messages asynchronously, such as timers or external API calls.

## Basic Application Structure

```go
package main

import "github.com/skyvence/TerminalEngineGo"

// Define your model
type model struct {
    // Your application state here
    counter int
    message string
}

// Initialize your model
func (m model) Init() engine.Msg {
    // Return initial message or nil
    return nil
}

// Handle messages and update state
func (m model) Update(msg engine.Msg) (engine.Model, engine.Cmd) {
    switch msg := msg.(type) {
    case engine.KeyMsg:
        switch msg.Rune {
        case 'q':
            return m, func() engine.Msg { return engine.Quit() }
        case '+':
            m.counter++
        case '-':
            m.counter--
        }
    }
    return m, nil
}

// Render your application
func (m model) View() string {
    return fmt.Sprintf("Counter: %d\nPress +/- to change, q to quit\n", m.counter)
}

func main() {
    initialModel := model{counter: 0, message: "Welcome!"}
    program := engine.NewProgram(initialModel)
    
    if err := program.Run(); err != nil {
        fmt.Printf("Error: %v\n", err)
        os.Exit(1)
    }
}
```

## Program Options

You can customize your program with options:

```go
// Enable alternate screen mode (full screen)
program := engine.NewProgram(model, engine.WithAltScreen())
```

## Input Handling

The engine automatically converts terminal input into messages:

- Regular keys become `KeyMsg` with the corresponding rune
- Arrow keys are converted to special runes: `↑`, `↓`, `←`, `→`
- Ctrl+C sends a `QuitMsg`

## Timers and Animation

Use the built-in timer functions for animations:

```go
func (m model) Init() engine.Msg {
    return engine.Tick(time.Second) // Start a 1-second timer
}

func (m model) Update(msg engine.Msg) (engine.Model, engine.Cmd) {
    switch msg := msg.(type) {
    case engine.TickMsg:
        // Handle timer tick
        m.frame = (m.frame + 1) % len(m.frames)
        return m, engine.Tick(200 * time.Millisecond) // Next frame
    }
    return m, nil
}
```

## Next Steps

- Check out the [API Reference](api-reference.md) for detailed function documentation
- Explore [Examples](examples.md) for practical implementations
- Look at the example projects in the `examples/` directory