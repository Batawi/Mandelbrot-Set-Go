package drawing_manager

import (
	"fmt"

	"github.com/Batawi/Mandelbrot-Set-Go/fractal_manager"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

var mousePos1 pixel.Vec
var mousePos2 pixel.Vec

func checkUserInput(win *pixelgl.Window) {

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
		mousePos1 = win.MousePosition()
	} else if win.JustReleased(pixelgl.MouseButton1) {
		mousePos2 = win.MousePosition()
		fmt.Println(mousePos2.Sub(mousePos1))
		fractal_manager.CameraMove(mousePos2.Sub(mousePos1))
	}

}
