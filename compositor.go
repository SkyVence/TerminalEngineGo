package engine

type Layer struct {
	Buffer *PixelBuffer
	ZIndex int
	Alpha  float32 // 0.0 (transparent) to 1.0 (opaque)
}

type Compositor struct {
	Layers []*Layer
}

func (c *Compositor) AddLayer(layer *Layer) {
	c.Layers = append(c.Layers, layer)
	// Sort layers by ZIndex --> Look for optimized sorting algorithm
}

func (c *Compositor) Composite() *PixelBuffer {
	result := NewPixelBuffer(c.Layers[0].Buffer.Width, c.Layers[0].Buffer.Height)

	for _, layer := range c.Layers {
		for y := 0; y < layer.Buffer.Height; y++ {
			for x := 0; x < layer.Buffer.Width; x++ {
				pixel := layer.Buffer.Data[y][x]
				if layer.Alpha > 0.5 {
					result.Data[y][x] = pixel
				}
			}
		}
	}
	return result
}
