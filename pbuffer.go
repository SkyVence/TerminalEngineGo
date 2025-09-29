package engine

import (
	"fmt"
	"math"
	"strings"
)

type Color int

const (
	ColorBlack Color = iota
	ColorRed
	ColorGreen
	ColorYellow
	ColorBlue
	ColorMagenta
	ColorCyan
	ColorWhite
)

type Pixel struct {
	Char rune
	FG   Color
	BG   Color
}

type PixelBuffer struct {
	Width  int
	Height int
	Data   [][]Pixel
}

func NewPixelBuffer(width, height int) *PixelBuffer {
	data := make([][]Pixel, height)
	for i := range data {
		data[i] = make([]Pixel, width)
	}
	return &PixelBuffer{
		Width:  width,
		Height: height,
		Data:   data,
	}
}

func (pb *PixelBuffer) SetPixel(x, y int, p Pixel) {
	if x >= 0 && x < pb.Width && y >= 0 && y < pb.Height {
		pb.Data[y][x] = p
	}
}

func (pb *PixelBuffer) FillRect(x, y, w, h int, p Pixel) {
	for i := y; i < y+h && i < pb.Height; i++ {
		for j := x; j < x+w && j < pb.Width; j++ {
			pb.Data[i][j] = p
		}
	}
}

func (pb *PixelBuffer) DrawLine(x1, y1, x2, y2 int, p Pixel) {
	dx := int(math.Abs(float64(x1 - x2)))
	dy := int(math.Abs(float64(y1 - y2)))
	sx := 1
	if x1 > x2 {
		sx = -1
	}
	sy := 1
	if y1 > y2 {
		sy = -1
	}

	x, y := x1, y1
	if dx >= dy {
		d := 2*dy - dx
		for i := 0; i <= dx; i++ {
			pb.SetPixel(x, y, p)
			if d > 0 {
				y += sy
				d -= 2 * dx
			}
			d += 2 * dy
			x += sx
		}
	} else {
		d := 2*dx - dy
		for i := 0; i <= dy; i++ {
			pb.SetPixel(x, y, p)
			if d > 0 {
				x += sx
				d -= 2 * dy
			}
			d += 2 * dx
			y += sy
		}
	}
}

func (pb *PixelBuffer) RenderToTerminal() string {
	var output strings.Builder

	for y := 0; y < pb.Height; y++ {
		for x := 0; x < pb.Width; x++ {
			pixel := pb.Data[y][x]

			// Set colors
			output.WriteString(fmt.Sprintf("\x1b[38;5;%dm\x1b[48;5;%dm",
				uint8(pixel.FG), uint8(pixel.BG)))

			// Write character
			output.WriteRune(pixel.Char)

			// Reset colors
			output.WriteString("\x1b[0m")
		}
		output.WriteString("\n")
	}

	return output.String()
}
