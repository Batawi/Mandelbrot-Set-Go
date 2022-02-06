package main

import (
	"fmt"
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

var canvas *pixelgl.Canvas

func fractalManagerInit(bounds pixel.Rect) *pixelgl.Canvas {
	canvas = pixelgl.NewCanvas(bounds)

	return canvas
}

func fractalManagerUpdate() {

	fmt.Println(canvas.Bounds().Max.X)
	fmt.Println(canvas.Bounds().Max.Y)

	pixels := canvas.Pixels()

	for i := 0; i < int(canvas.Bounds().Max.Y); i++ {
		for j := 0; j < int(canvas.Bounds().Max.X); j++ {

			coord := complex(
				mapValueToRange(float64(j), 0, canvas.Bounds().Max.X, -2, 2),
				mapValueToRange(float64(i), 0, canvas.Bounds().Max.Y, -2, 2))

			escaping_point := coord

			c := 0
			for c = 0; c < 20; c++ {
				if cmplx.Abs(escaping_point) >= 4 { //4

					pixels[i*4*int(canvas.Bounds().Max.X)+j*4] = 5
					pixels[i*4*int(canvas.Bounds().Max.X)+j*4+1] = 5
					pixels[i*4*int(canvas.Bounds().Max.X)+j*4+2] = 5
					pixels[i*4*int(canvas.Bounds().Max.X)+j*4+3] = 255
					break
				}
				// escaping_point = cmplx.Pow(escaping_point, complex(2, 0)) + coord
				escaping_point = escaping_point*escaping_point + coord
			}

			if c == 20 {
				pixels[i*4*int(canvas.Bounds().Max.X)+j*4] = 200
				pixels[i*4*int(canvas.Bounds().Max.X)+j*4+1] = 200
				pixels[i*4*int(canvas.Bounds().Max.X)+j*4+2] = 200
				pixels[i*4*int(canvas.Bounds().Max.X)+j*4+3] = 255
			}

		}
	}
	canvas.SetPixels(pixels)
}
