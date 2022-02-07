package main

import (
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const INITIAL_WINDOW_WIDTH = 1024
const INITIAL_WINDOW_HEIGHT = 786

func windowInit() *pixelgl.Window {
	cfg := pixelgl.WindowConfig{
		Title:  "MandelBrot, Set, Go!",
		Bounds: pixel.R(0, 0, INITIAL_WINDOW_WIDTH, INITIAL_WINDOW_HEIGHT),
		VSync:  true,
		// Undecorated: true,
		// Position:    pixel.V(100, 100),
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	return win
}

func run() {
	var DeltaTime float64

	win := windowInit()
	canvas := FractalManagerInit(win.Bounds(), &DeltaTime)
	FractalManagerUpdate()

	last := time.Now()
	for !win.Closed() {
		DeltaTime = time.Since(last).Seconds()
		last = time.Now()

		win.Clear(colornames.Skyblue)

		checkUserInput(win)

		canvas.Draw(win, StretchToFit(canvas, win))

		win.Update()
	}
}

// every target has SetMatric method which might be usefull someday
// win.SetMatrix(pixel.IM.Moved(pixel.V(0, 0)))

// example of IMDRAW
// imd := imdraw.New(nil)
// imd.Clear()
// imd.Color = pixel.RGB(1, 0, 0)
// imd.Push(win.Bounds().Center().Add(pixel.V(0, 0)))
// imd.Circle(900, 0)
// imd.Draw(canvas)
