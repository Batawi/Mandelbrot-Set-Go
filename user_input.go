package main

import (
	"fmt"

	"github.com/faiface/pixel/pixelgl"
)

func checkUserInput(win *pixelgl.Window) {

	if win.Pressed(pixelgl.KeyLeft) {
		// camPos.X -= camSpeed * dt
		// cam.moveLeft()
		fmt.Println("siema")
	}
}
