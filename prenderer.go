package engine

import "io"

type PixelRenderer struct {
	*StandardRenderer
	compositor *Compositor
}

func (pr *PixelRenderer) RenderPixels(buffer *PixelBuffer) {
	ansiString := buffer.RenderToTerminal()
	pr.Write(ansiString)
}

func NewPixelRenderer(out io.Writer) Renderer {
	sr := NewRenderer(out).(*StandardRenderer)
	pr := &PixelRenderer{
		StandardRenderer: sr,
		compositor:       &Compositor{},
	}
	return pr
}
