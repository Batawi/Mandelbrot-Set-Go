package drawing_manager

import (
	"time"

	"github.com/Batawi/Mandelbrot-Set-Go/fractal_manager"
	"github.com/Batawi/Mandelbrot-Set-Go/utils"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

// --- GLOBALS ---
const (
	initialWindowWidth  = 1024
	initialWindowHeight = 786
)

var (
	DeltaTime     float64
	prevWinBounds pixel.Rect
)

func windowInit() *pixelgl.Window {
	cfg := pixelgl.WindowConfig{
		Title:     "MandelBrot, Set, Go!",
		Bounds:    pixel.R(0, 0, initialWindowWidth, initialWindowHeight),
		VSync:     true,
		Resizable: true,
		// Undecorated: true,
		// Position:    pixel.V(100, 100),
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	return win
}

func Run() {

	win := windowInit()
	prevWinBounds = win.Bounds() // For checking if application window size has been changed
	win.SetSmooth(true)

	fractal_manager.Init(win.Bounds())

	prevTime := time.Now()
	for !win.Closed() {
		DeltaTime = time.Since(prevTime).Seconds()
		prevTime = time.Now()

		win.Clear(colornames.Skyblue)

		checkUserInput(win)
		fractal_manager.Update()

		fractal_manager.Canvas.Draw(win, utils.StretchToFit(fractal_manager.Canvas, win))

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
