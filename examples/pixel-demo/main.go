package main

import (
	"github.com/skyvence/TerminalEngineGo/engine"
)

type pixelModel struct {
	width  int
	height int
}

func (m pixelModel) Init() engine.Msg {
	return nil
}

func (m pixelModel) Update(msg engine.Msg) (engine.Model, engine.Cmd) {
	switch msg.(type) {
	case engine.KeyMsg:
		return m, engine.Quit
	}
	return m, nil
}

func (m pixelModel) View() string {
	return "This should not be called in pixel mode"
}

func (m pixelModel) PixelView() *engine.PixelBuffer {
	buffer := engine.NewPixelBuffer(m.width, m.height)

	// Draw a simple border
	for x := 0; x < m.width; x++ {
		buffer.SetPixel(x, 0, engine.Pixel{Char: '█', FG: engine.ColorWhite, BG: engine.ColorBlack})
		buffer.SetPixel(x, m.height-1, engine.Pixel{Char: '█', FG: engine.ColorWhite, BG: engine.ColorBlack})
	}
	for y := 0; y < m.height; y++ {
		buffer.SetPixel(0, y, engine.Pixel{Char: '█', FG: engine.ColorWhite, BG: engine.ColorBlack})
		buffer.SetPixel(m.width-1, y, engine.Pixel{Char: '█', FG: engine.ColorWhite, BG: engine.ColorBlack})
	}

	// Draw a filled rectangle in the center
	centerX, centerY := m.width/2, m.height/2
	rectWidth, rectHeight := 10, 5
	for y := centerY - rectHeight/2; y < centerY+rectHeight/2; y++ {
		for x := centerX - rectWidth/2; x < centerX+rectWidth/2; x++ {
			buffer.SetPixel(x, y, engine.Pixel{Char: '█', FG: engine.ColorRed, BG: engine.ColorBlue})
		}
	}

	return buffer
}

func main() {
	width, height := 80, 24
	model := pixelModel{width: width, height: height}

	p := engine.NewProgram(model, engine.WithBase())
	if err := p.Run(); err != nil {
		panic(err)
	}
}
