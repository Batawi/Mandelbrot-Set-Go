package main

import (
	"math/cmplx"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

var canvas *pixelgl.Canvas

func fractalManagerInit(bounds pixel.Rect) *pixelgl.Canvas {
	canvas = pixelgl.NewCanvas(bounds)

	return canvas
}

func fractalManagerUpdate() {
	pixels := canvas.Pixels()

	var y float64
	var x float64

	for i := 0; i < INITIAL_WINDOW_HEIGHT; i++ {
		for j := 0; j < INITIAL_WINDOW_WIDTH; j++ {

			y = scale_value_to_range(float64(i), 0, INITIAL_WINDOW_HEIGHT, -2, 2)
			x = scale_value_to_range(float64(j), 0, INITIAL_WINDOW_WIDTH, -2, 2)

			coord := complex(x, y)
			escaping_point := coord

			c := 0
			for c = 0; c < 20; c++ {
				if cmplx.Abs(escaping_point) >= 4 { //4

					pixels[i*4*INITIAL_WINDOW_WIDTH+j*4] = 0
					pixels[i*4*INITIAL_WINDOW_WIDTH+j*4+1] = 0
					pixels[i*4*INITIAL_WINDOW_WIDTH+j*4+2] = 0
					pixels[i*4*INITIAL_WINDOW_WIDTH+j*4+3] = 255
					break
				}
				// escaping_point = cmplx.Pow(escaping_point, complex(2, 0)) + coord
				escaping_point = escaping_point*escaping_point + coord
			}

			if c == 20 {
				pixels[i*4*INITIAL_WINDOW_WIDTH+j*4] = 255
				pixels[i*4*INITIAL_WINDOW_WIDTH+j*4+1] = 255
				pixels[i*4*INITIAL_WINDOW_WIDTH+j*4+2] = 255
				pixels[i*4*INITIAL_WINDOW_WIDTH+j*4+3] = 255
			}

		}
	}
	canvas.SetPixels(pixels)
}
