package drawing_manager

import (
	"github.com/Batawi/Mandelbrot-Set-Go/fractal_manager"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

var savedMousePos pixel.Vec

func checkUserInput(win *pixelgl.Window) {

	// Check if application window has been resized
	if prevWinBounds != win.Bounds() {
		fractal_manager.UpdateWinBounds(win.Bounds())
		prevWinBounds = win.Bounds()
	}

	// Check input for camera movement
	if win.Pressed(pixelgl.KeyUp) {
		fractal_manager.CameraMoveUp()
	}

	if win.Pressed(pixelgl.KeyDown) {
		fractal_manager.CameraMoveDown()
	}

	if win.Pressed(pixelgl.KeyRight) {
		fractal_manager.CameraMoveRight()
	}

	if win.Pressed(pixelgl.KeyLeft) {
		fractal_manager.CameraMoveLeft()
	}

	if win.JustPressed(pixelgl.MouseButton1) {
		savedMousePos = win.MousePosition()

	} else if win.JustReleased(pixelgl.MouseButton1) {

		if savedMousePos != win.MousePosition() {
			fractal_manager.CameraMove(savedMousePos.Sub(win.MousePosition()))
		}
	}

	if win.JustPressed(pixelgl.MouseButton2) {
		fractal_manager.CameraMoveCenter(win.MousePosition())
	}

	scrollCounts := win.MouseScroll().Y
	if scrollCounts != 0 {
		fractal_manager.CameraZoom(scrollCounts)
	}

	if win.JustPressed(pixelgl.KeyZ) {
		fractal_manager.IterationsUp()
	}

	if win.JustPressed(pixelgl.KeyX) {
		fractal_manager.IterationsDown()
	}
}
