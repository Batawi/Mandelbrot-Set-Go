package main

import (
	"math/cmplx"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Camera struct {
	camPos       pixel.Vec
	camMoveSpeed float64
	camZoom      float64
	// camZoomSpeed float64
}

func (cam Camera) MoveUp() {

}

var canvas *pixelgl.Canvas
var DeltaTime *float64

func FractalManagerInit(bounds pixel.Rect, deltaTime *float64) *pixelgl.Canvas {
	canvas = pixelgl.NewCanvas(bounds)
	DeltaTime = deltaTime

	return canvas
}

func FractalManagerUpdate() {

	canvas.SetBounds(pixel.R(0, 0, 50, 400))
	pixels := canvas.Pixels()

	for i := 0; i < int(canvas.Bounds().H()); i++ {
		for j := 0; j < int(canvas.Bounds().W()); j++ {

			coord := complex(
				MapValueToRange(float64(j), 0, canvas.Bounds().W(), -2, 2),
				MapValueToRange(float64(i), 0, canvas.Bounds().H(), -2, 2))

			escaping_point := coord

			c := 0
			for c = 0; c < 20; c++ {
				if cmplx.Abs(escaping_point) >= 4 { //4

					pixels[i*4*int(canvas.Bounds().W())+j*4] = 5
					pixels[i*4*int(canvas.Bounds().W())+j*4+1] = 5
					pixels[i*4*int(canvas.Bounds().W())+j*4+2] = 5
					pixels[i*4*int(canvas.Bounds().W())+j*4+3] = 255
					break
				}
				// escaping_point = cmplx.Pow(escaping_point, complex(2, 0)) + coord
				escaping_point = escaping_point*escaping_point + coord
			}

			if c == 20 {
				pixels[i*4*int(canvas.Bounds().W())+j*4] = 200
				pixels[i*4*int(canvas.Bounds().W())+j*4+1] = 200
				pixels[i*4*int(canvas.Bounds().W())+j*4+2] = 200
				pixels[i*4*int(canvas.Bounds().W())+j*4+3] = 255
			}

		}
	}
	canvas.SetPixels(pixels)
}
