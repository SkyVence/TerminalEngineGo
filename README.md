# TerminalEngineGo

![Version](https://img.shields.io/github/v/tag/SkyVence/TerminalEngineGo?label=version&sort=semver)
![Go Version](https://img.shields.io/github/go-mod/go-version/SkyVence/TerminalEngineGo)
![License](https://img.shields.io/github/license/SkyVence/TerminalEngineGo)
![Build Status](https://img.shields.io/github/actions/workflow-status/SkyVence/TerminalEngineGo/release.yml)

A simple, elegant terminal game engine written in Go that provides a foundation for building interactive terminal applications and games.

## Features

- **Event-driven architecture** with Model-View-Update pattern
- **Terminal input handling** with support for arrow keys and special characters
- **Flexible rendering system** with double buffering and ANSI escape sequences
- **Animation support** with frame-based animations
- **Localization support** with JSON-based language files
- **Alternate screen mode** for full-screen applications
- **Simple API** that's easy to learn and use

## Features planned

- **Compositing** multiple layer for ui's
- **Mouse support**

## Installation

```bash
go get github.com/skyvence/TerminalEngineGo
```

## Quick Start

```go
package main

import (
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
    return "Hello, Terminal Engine! Press 'q' to quit.\n"
}

func main() {
    m := model{content: "Hello World"}
    p := engine.NewProgram(m)
    p.Run()
}
```

## Documentation

For detailed documentation, see the [docs](./docs/) directory:

- [Getting Started](./docs/getting-started.md)
- [API Reference](./docs/api-reference.md)
- [Examples](./docs/examples.md)

## Examples

Check out the [examples](./examples/) directory for practical implementations:

- [Hello World](./examples/hello-world/) - Basic terminal application
- [Animation Demo](./examples/animation/) - Frame-based animation example
- [Game Example](./examples/game/) - Simple interactive game

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Changelog

See [CHANGELOG.md](CHANGELOG.md) for version history and changes.
