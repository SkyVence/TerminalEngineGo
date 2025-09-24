# API Reference

## Core Types

### Model Interface

The Model interface defines the contract for your application's state management:

```go
type Model interface {
    Init() Msg
    Update(msg Msg) (Model, Cmd)
    View() string
}
```

- **Init()**: Called once when the program starts. Return an initial message or nil.
- **Update(msg Msg)**: Handle incoming messages and return new model state and optional command.
- **View()**: Return a string representation of your model for rendering.

### Message Types

#### KeyMsg
```go
type KeyMsg struct {
    Rune rune
}
```
Represents keyboard input. Special keys are converted to Unicode symbols:
- `↑` - Up arrow
- `↓` - Down arrow  
- `←` - Left arrow
- `→` - Right arrow

#### QuitMsg
```go
type QuitMsg struct{}
```
Signals the program should exit. Send this to quit the application.

#### TickMsg
```go
type TickMsg struct {
    Time time.Time
}
```
Timer message containing the time when the tick occurred.

#### SizeMsg
```go
type SizeMsg struct {
    Width  int
    Height int
}
```
Terminal size change notification.

### Commands

Commands are functions that return messages:

```go
type Cmd func() Msg
```

#### Built-in Commands

**Quit()**
```go
func Quit() Msg
```
Returns a QuitMsg to exit the program.

**Tick(duration)**
```go
func Tick(d time.Duration) Cmd
```
Returns a command that sends a TickMsg after the specified duration.

**TickNow()**
```go
func TickNow() Cmd
```
Returns a command that sends a TickMsg immediately.

## Program

### NewProgram
```go
func NewProgram(model Model, opts ...ProgramOption) *Program
```
Creates a new Program with the given model and options.

### Program Methods

**Run()**
```go
func (p *Program) Run() error
```
Starts the program main loop. Blocks until the program exits.

**GetSize()**
```go
func (p *Program) GetSize() (int, int)
```
Returns the current terminal dimensions (width, height).

**GetRenderer()**
```go
func (p *Program) GetRenderer() Renderer  
```
Returns the renderer instance for advanced usage.

### Program Options

**WithAltScreen()**
```go
func WithAltScreen() ProgramOption
```
Enables alternate screen buffer for full-screen applications.

## Game Interface

For game development, you can use the Game interface which is compatible with Model:

```go
type Game interface {
    Init() Msg
    Update(Msg) (Model, Cmd)
    View() string
}
```

**Wrap(game)**
```go
func Wrap(g Game) Model
```
Converts a Game into a Model for use with NewProgram.

## Animation

### Animation Type
```go
type Animation struct {
    Frames []string
    // private fields
}
```

**NewAnimation(frames)**
```go
func NewAnimation(frames []string) Animation
```
Creates a new animation with the given frames and default 200ms speed.

**LoadAnimationFile(filename)**
```go
func LoadAnimationFile(filename string) ([]string, error)
```
Loads animation frames from a file. Frames should be separated by "---".

### Animation Methods

**Init()**
```go
func (a Animation) Init() Cmd
```
Returns initial tick command to start the animation.

**Update(msg)**
```go
func (a Animation) Update(msg Msg) (Animation, Cmd)
```
Advances animation frame on TickMsg and returns next tick command.

**View()**
```go
func (a Animation) View() string
```
Returns current animation frame as string.

## Localization

### LocalizationManager

**GetLocalizationManager()**
```go
func GetLocalizationManager() *LocalizationManager
```
Returns the singleton localization manager instance.

**SetLanguage(lang)**
```go
func (lm *LocalizationManager) SetLanguage(lang string) error
```
Loads and sets the catalog for the specified language.

**Text(key, args...)**
```go
func (lm *LocalizationManager) Text(key string, args ...any) string
```
Retrieves localized text with placeholder replacement.

**GetCurrentLanguage()**
```go
func (lm *LocalizationManager) GetCurrentLanguage() string
```
Returns the currently set language code.

**GetSupportedLanguages()**
```go
func (lm *LocalizationManager) GetSupportedLanguages() ([]string, error)
```
Returns available language codes from assets/interface/ directory.

## Renderer (Advanced)

The renderer handles terminal output and can be accessed for advanced usage:

### Renderer Methods

**Write(s string)**
Buffers content for rendering.

**ShowCursor() / HideCursor()**
Controls cursor visibility.

**SetWindowTitle(title string)**
Sets the terminal window title.

**EnterAltScreen() / ExitAltScreen()**
Switches between normal and alternate screen buffers.

**SetCursor(x, y int)**
Positions cursor at specific coordinates (alternate screen only).

## Global Functions

**SetGlobalRenderer(renderer)**
```go
func SetGlobalRenderer(renderer Renderer)
```
Sets the global renderer instance.

**GetGlobalRenderer()**
```go
func GetGlobalRenderer() Renderer
```
Returns the global renderer instance.