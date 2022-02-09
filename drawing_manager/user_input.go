package drawing_manager

import (
	"fmt"

	"github.com/Batawi/Mandelbrot-Set-Go/fractal_manager"
	"github.com/faiface/pixel/pixelgl"
)

func checkUserInput(win *pixelgl.Window) {

	if win.Pressed(pixelgl.KeyUp) {
		fractal_manager.CameraMoveUp()
	}
	if win.Pressed(pixelgl.MouseButton1) {
		fmt.Println("siema")
	}
}
