package fractal_manager

import (
	"math"

	"github.com/Batawi/Mandelbrot-Set-Go/utils"
	"github.com/faiface/pixel"
)

func CameraMoveUp() {
	fractalBounds = fractalBounds.Moved(pixel.V(0, windowBounds.H()*moveSpeed))
	machineState = update
}

func CameraMoveDown() {
	fractalBounds = fractalBounds.Moved(pixel.V(0, windowBounds.H()*moveSpeed*-1))
	machineState = update
}

func CameraMoveRight() {
	fractalBounds = fractalBounds.Moved(pixel.V(windowBounds.H()*moveSpeed, 0))
	machineState = update
}

func CameraMoveLeft() {
	fractalBounds = fractalBounds.Moved(pixel.V(windowBounds.H()*moveSpeed*-1, 0))
	machineState = update
}

func CameraMove(v pixel.Vec) {
	v.X = utils.MapValueToRange(v.X, 0, windowBounds.W(), 0, fractalBounds.W())
	v.Y = utils.MapValueToRange(v.Y, 0, windowBounds.H(), 0, fractalBounds.H())

	fractalBounds = fractalBounds.Moved(v)
	machineState = update
}

func CameraMoveCenter(v pixel.Vec) {
	v.X = utils.MapValueToRange(v.X, 0, windowBounds.W(), fractalBounds.Min.X, fractalBounds.Max.X)
	v.Y = utils.MapValueToRange(v.Y, 0, windowBounds.H(), fractalBounds.Min.Y, fractalBounds.Max.Y)

	fractalBounds = fractalBounds.Moved(v.Sub(fractalBounds.Center()))
	machineState = update
}

func CameraZoom(zoomCounts float64) {
	scale := math.Pow(camZoomSpeed, zoomCounts)
	fractalBounds = utils.ScaleRect(fractalBounds, scale)
	machineState = update
}

func IterationsUp() {
	iterationsLimit += iterationsJump
	machineState = update
}

func IterationsDown() {
	if iterationsLimit > iterationsJump {
		iterationsLimit -= iterationsJump
		machineState = update
	}
}

func UpdateWinBounds(r pixel.Rect) {
	windowBounds = r
	fractalBounds = utils.ScaleRectToRect(windowBounds, fractalBounds)
	Canvas.SetBounds(r)
	machineState = update
}
