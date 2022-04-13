package fractal_manager

import (
	"fmt"
	"sync"
	"time"

	"github.com/Batawi/Mandelbrot-Set-Go/utils"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

// type fractal_t struct {
// 	inputRange  pixel.Rect
// 	outputRange pixel.Rect
// 	iterCount   []uint64
// 	colorRGBA   []uint8
// }

// --- GLOBALS ---
type states uint8

const (
	wait states = iota
	update
	updateColors
)

var (
	Canvas       *pixelgl.Canvas
	windowBounds pixel.Rect // This is size of application window
	// fractalResolution float64    = 0.5

	fractalBounds           = pixel.R(-2, -2, 2, 2)
	moveSpeed               = 0.0001
	camZoomSpeed    float64 = 0.8
	iterationsLimit uint64  = 100
	iterationsJump  uint64  = 20
	machineState    states  = update
	bailoutRange    float64 = 8  // By definition should be 2 (or 4 if we dont sqrt() both sides) but higher values don't creates color bands
	maxGoroutines   uint32  = 25 //8ms
	// maxGoroutines uint32 = 1
)

// --- FUNCTIONS ---
func Init(bounds pixel.Rect) {
	Canvas = pixelgl.NewCanvas(bounds)
	windowBounds = bounds

	// Set fractal view bounds according to application window bounds
	fractalBounds = utils.ScaleRectToRect(windowBounds, fractalBounds)
}

func Update() {

	// This slice stores number of iterations for given point,
	// purpose of this is to recolor fractal wihtout calculating it again
	iterCounter := make([][]uint64, int(windowBounds.H()))
	for i := range iterCounter {
		iterCounter[i] = make([]uint64, int(windowBounds.W()))
	}

	pixels := make([]uint8, int(windowBounds.Area())*4)

	if machineState == update {
		start := time.Now()
		workDistributor(windowBounds, fractalBounds, iterCounter, pixels)
		fmt.Println(time.Since(start))

		Canvas.SetPixels(pixels)
		machineState = wait
	}
}

/*
Update() -> workDistributor() -> calculateMandelbrot()
*/

// This functions splits up given area and distribute chunks to separate goroutines
func workDistributor(inR, outR pixel.Rect, iterCounter [][]uint64, pixels []uint8) {

	noWorkers := maxGoroutines
	// Check if windows is not too small
	if maxGoroutines > uint32(inR.H()) {
		noWorkers = uint32(inR.H())
	}

	// Setup wait group
	var wg sync.WaitGroup
	wg.Add(int(noWorkers))

	for i := 0; i < int(noWorkers); i++ {

		i := i
		go func() {
			defer wg.Done()

			inRSlice := utils.ChopHor(inR, int32(noWorkers), int32(i))
			calculateMandelbrot(inRSlice, outR, iterCounter, pixels)

		}()
	}
	wg.Wait()
}

// inR concerns rect of application window
// outR concerns rect in complex plane
func calculateMandelbrot(inR, outR pixel.Rect, iterCounter [][]uint64, pixels []uint8) {

	// Slopes for later values mapping
	slopeX := (outR.Max.X - outR.Min.X) / (windowBounds.Max.X - windowBounds.Min.X)
	slopeY := (outR.Max.Y - outR.Min.Y) / (windowBounds.Max.Y - windowBounds.Min.Y)

	for i := int(inR.Min.Y); i < int(inR.Max.Y); i++ {
		for j := int(inR.Min.X); j < int(inR.Max.X); j++ {

			pointX := (float64(j)-windowBounds.Min.X)*slopeX + outR.Min.X
			pointY := (float64(i)-windowBounds.Min.Y)*slopeY + outR.Min.Y

			// Escaping point
			x := pointX
			y := pointY

			var iterations uint64 = 0

			if testBulb(pointX, pointY) || testCardioid(pointX, pointY) {
				iterations = iterationsLimit

			} else {

				for iterations = 0; x*x+y*y <= bailoutRange && iterations < iterationsLimit; iterations++ {
					xTemp := x*x - y*y + pointX
					y = 2*x*y + pointY
					x = xTemp
				}
			}

			iterCounter[i][j] = iterations

			color := calculateColor(iterations, iterationsLimit, x, y)
			colorToPixels(uint64(i), uint64(j), uint64(inR.W()), pixels, color)
		}
	}
}

// Check wiki:
// https://en.wikipedia.org/wiki/Plotting_algorithms_for_the_Mandelbrot_set#Cardioid_/_bulb_checking
func testBulb(x, y float64) bool {
	a := x + 1
	return a*a+y*y <= 0.0625
}

// Check wiki:
// https://en.wikipedia.org/wiki/Plotting_algorithms_for_the_Mandelbrot_set#Cardioid_/_bulb_checking
func testCardioid(x, y float64) bool {
	q := (x-0.25)*(x-0.25) + y*y
	return q*(q+(x-0.25)) <= 0.25*y*y
}

func colorToPixels(i, j, width uint64, pixels []uint8, color pixel.RGBA) {
	pixels[4*(i*width+j)] = uint8(color.R)
	pixels[4*(i*width+j)+1] = uint8(color.G)
	pixels[4*(i*width+j)+2] = uint8(color.B)
	pixels[4*(i*width+j)+3] = uint8(color.A)
}
