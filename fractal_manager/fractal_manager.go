package fractal_manager

import (
	"math/cmplx"

	"github.com/Batawi/Mandelbrot-Set-Go/utils"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Camera struct {
	pos          pixel.Vec
	moveSpeed    float64 // It is expressed as a percentage of the window size
	zoom         float64
	camZoomSpeed float64
}

// --- GLOBALS ---
// const (
// 	initialFractalBoundsMinX float64 = -2
// 	initialFractalBoundsMinY float64 = -2
// 	initialFractalBoundsMaxX float64 = 2
// 	initialFractalBoundsMaxY float64 = 2
// )

var (
	mainCam = Camera{
		moveSpeed:    20,
		zoom:         1,
		camZoomSpeed: 10,
	}

	canvas        *pixelgl.Canvas
	deltaTime     *float64
	windowBounds  pixel.Rect // This is size of application window
	fractalBounds pixel.Rect // This is camera view boundries inside complex plane
)

func CameraMoveUp() {
	mainCam.pos.Y += 0.1
}

func Init(bounds pixel.Rect, deltaTime *float64) *pixelgl.Canvas {
	canvas = pixelgl.NewCanvas(bounds)
	deltaTime = deltaTime
	windowBounds = bounds
	fractalBounds = pixel.R(-2, -2, 2, 2)

	// Set fractal view bounds according to window bounds

	// if bounds.W() >= bounds.H() {

	// 	// Calculate new fractal bounds height
	// 	newH = fractalBounds.W() * windowBounds.H() / windowBounds.W()

	// 	fractalBounds.Min.Y
	// 	fractalBounds.Max.Y

	// } else {

	// }

	utils.ScaleRectToRect(windowBounds, fractalBounds)

	return canvas
}

func Update() {

	canvas.SetBounds(pixel.R(0, 0, 50, 50))
	pixels := canvas.Pixels()

	for i := 0; i < int(canvas.Bounds().H()); i++ {
		for j := 0; j < int(canvas.Bounds().W()); j++ {

			coord := complex(
				utils.MapValueToRange(float64(j), 0, canvas.Bounds().W(), -2, 2),
				utils.MapValueToRange(float64(i), 0, canvas.Bounds().H(), -2+mainCam.pos.Y, 2+mainCam.pos.Y))

			escaping_point := coord

			c := 0
			for c = 0; c < 20; c++ {
				if cmplx.Abs(escaping_point) >= 4 { //4

					pixels[i*4*int(canvas.Bounds().W())+j*4] = 5
					pixels[i*4*int(canvas.Bounds().W())+j*4+1] = 5
					pixels[i*4*int(canvas.Bounds().W())+j*4+2] = 5
					pixels[i*4*int(canvas.Bounds().W())+j*4+3] = 255
					break
				}
				// escaping_point = cmplx.Pow(escaping_point, complex(2, 0)) + coord
				escaping_point = escaping_point*escaping_point + coord
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
