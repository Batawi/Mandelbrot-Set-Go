package main

import (
	"github.com/Batawi/Mandelbrot-Set-Go/drawing_manager"
	"github.com/faiface/pixel/pixelgl"
)

func main() {
	pixelgl.Run(drawing_manager.Run)
}
