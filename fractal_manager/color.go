package fractal_manager

import (
	"math"

	"github.com/faiface/pixel"
)

func calculateColor(iterations uint64, iterlimit uint64, x, y float64) pixel.RGBA {
	if iterations == iterlimit {
		return pixel.RGBA{10, 10, 10, 255}
	}

	var s float64 //smooth coef

	s = float64(iterations) + 1.0 - math.Log(math.Log(math.Sqrt(x*x+y*y)))/math.Log(2)

	s /= (float64(iterlimit))

	return pixel.RGBA{s * 250, 10, 10, 255}
}
