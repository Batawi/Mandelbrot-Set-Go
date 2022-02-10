package fractal_manager

import (
	"fmt"
	"math/cmplx"

	"github.com/Batawi/Mandelbrot-Set-Go/utils"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Camera struct {
	fractalBounds pixel.Rect // This is camera view boundries inside complex plane. Position vector is not needed
	moveSpeed     float64    // It is expressed as a percentage (from 0 to 1) of the window size
	zoom          float64
	camZoomSpeed  float64
}

// --- GLOBALS ---
var (
	mainCam = Camera{
		fractalBounds: pixel.R(-2, -2, 2, 2),
		moveSpeed:     0.0001,
		zoom:          1,
		camZoomSpeed:  10,
	}

	canvas       *pixelgl.Canvas
	deltaTime    *float64
	windowBounds pixel.Rect // This is size of application window
	// fractalResolution float64    = 0.1
)

func Init(bounds pixel.Rect, deltaTime *float64) *pixelgl.Canvas {
	canvas = pixelgl.NewCanvas(bounds)
	deltaTime = deltaTime
	windowBounds = bounds

	// Set fractal view bounds according to application window bounds
	mainCam.fractalBounds = utils.ScaleRectToRect(windowBounds, mainCam.fractalBounds)

	// canvas.SetBounds(pixel.R(0, 0, canvas.Bounds().Max.X*fractalResolution, canvas.Bounds().Max.Y*fractalResolution))

	return canvas
}

func Update() {

	pixels := canvas.Pixels()

	for i := 0; i < int(canvas.Bounds().H()); i++ {
		for j := 0; j < int(canvas.Bounds().W()); j++ {

			point := complex(
				utils.MapValueToRange(float64(j), 0, canvas.Bounds().W(), mainCam.fractalBounds.Min.X, mainCam.fractalBounds.Max.X),
				utils.MapValueToRange(float64(i), 0, canvas.Bounds().H(), mainCam.fractalBounds.Min.Y, mainCam.fractalBounds.Max.Y))

			escapingPoint := point

			c := 0
			for c = 0; c < 20; c++ {
				if cmplx.Abs(escapingPoint) >= 4 { //4

					pixels[i*4*int(canvas.Bounds().W())+j*4] = 5
					pixels[i*4*int(canvas.Bounds().W())+j*4+1] = 5
					pixels[i*4*int(canvas.Bounds().W())+j*4+2] = 5
					pixels[i*4*int(canvas.Bounds().W())+j*4+3] = 255
					break
				}
				// escapingPoint = cmplx.Pow(escapingPoint, complex(2, 0)) + point
				escapingPoint = escapingPoint*escapingPoint + point
			}

			if c == 20 {
				pixels[i*4*int(canvas.Bounds().W())+j*4] = 200
				pixels[i*4*int(canvas.Bounds().W())+j*4+1] = 200
				pixels[i*4*int(canvas.Bounds().W())+j*4+2] = 200
				pixels[i*4*int(canvas.Bounds().W())+j*4+3] = 255
			}

		}
	}
	canvas.SetPixels(pixels)
}

func CameraMoveUp() {
	mainCam.fractalBounds = mainCam.fractalBounds.Moved(pixel.V(0, windowBounds.H()*mainCam.moveSpeed))
}

func CameraMoveDown() {
	mainCam.fractalBounds = mainCam.fractalBounds.Moved(pixel.V(0, windowBounds.H()*mainCam.moveSpeed*-1))
}

func CameraMoveRight() {
	mainCam.fractalBounds = mainCam.fractalBounds.Moved(pixel.V(windowBounds.H()*mainCam.moveSpeed, 0))
}

func CameraMoveLeft() {
	mainCam.fractalBounds = mainCam.fractalBounds.Moved(pixel.V(windowBounds.H()*mainCam.moveSpeed*-1, 0))
}

func CameraMove(v pixel.Vec) {
	fmt.Println("before maping: ", v)

	v.X = utils.MapValueToRange(v.X, 0, windowBounds.W(), 0, mainCam.fractalBounds.W())
	v.Y = utils.MapValueToRange(v.Y, 0, windowBounds.H(), 0, mainCam.fractalBounds.H())

	fmt.Println("after maping: ", v)
	mainCam.fractalBounds = mainCam.fractalBounds.Moved(v)
}
