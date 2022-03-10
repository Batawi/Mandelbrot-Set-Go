package fractal_manager

import (
	"math"

	"github.com/faiface/pixel"
	"github.com/lucasb-eyer/go-colorful"
)

func calculateColor(iterations uint64, iterlimit uint64, x, y float64) pixel.RGBA {
	if iterations == iterlimit {
		return pixel.RGBA{10, 10, 10, 255}
	}

	var s float64 //smooth coef

	s = float64(iterations) + 1.0 - math.Log(math.Log(math.Sqrt(x*x+y*y)))/math.Log(2)

	s /= (float64(iterlimit))

	// c := colorful.Hsl(s*360.0, 1, 100)
	c := colorful.Hsv(s*360.0, 1, 199)

	// return pixel.RGBA{s * 250, 10, 10, 255}
	// return pixel.RGBA{c.R, c.G, c.B, 255}
	return pixel.RGBA{c.R, c.G, c.B, 215}
}
