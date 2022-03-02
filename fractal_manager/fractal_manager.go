package fractal_manager

import (
	"fmt"
	"math"
	"math/cmplx"

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
var (
	Canvas       *pixelgl.Canvas
	windowBounds pixel.Rect // This is size of application window
	// fractalResolution float64    = 0.5

	fractalBounds           = pixel.R(-2, -2, 2, 2)
	moveSpeed               = 0.0001
	camZoomSpeed    float64 = 0.5
	iterationsLimit uint64  = 20
	iterationsJump  uint64  = 20
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
	calculateMandelbrot(windowBounds, fractalBounds, iterCounter, pixels)

	Canvas.SetPixels(pixels)
}

// inR concerns rect of application window
// outR concerns rect in complex plane
func calculateMandelbrot(inR, outR pixel.Rect, iterCounter [][]uint64, pixels []uint8) {

	for i := 0; i < int(inR.H()); i++ {
		for j := 0; j < int(inR.W()); j++ {

			point := complex(
				utils.MapValueToRange(float64(j), 0, inR.W(), outR.Min.X, outR.Max.X),
				utils.MapValueToRange(float64(i), 0, inR.H(), outR.Min.Y, outR.Max.Y))

			escapingPoint := point

			var color pixel.RGBA
			var iterations uint64 = 0
			for iterations = 0; iterations < iterationsLimit; iterations++ {

				if cmplx.Abs(escapingPoint) >= 4 { //4
					break
				}
				// escapingPoint = cmplx.Pow(escapingPoint, complex(2, 0)) + point
				escapingPoint = escapingPoint*escapingPoint + point
			}

			iterCounter[i][j] = iterations

			color = calculateColor(iterations, iterationsLimit, escapingPoint)
			colorToPixels(uint64(i), uint64(j), uint64(inR.W()), pixels, color)
		}
	}

}
func colorToPixels(i, j, width uint64, pixels []uint8, color pixel.RGBA) {
	pixels[4*(i*width+j)] = uint8(color.R)
	pixels[4*(i*width+j)+1] = uint8(color.G)
	pixels[4*(i*width+j)+2] = uint8(color.B)
	pixels[4*(i*width+j)+3] = uint8(color.A)
}

func calculateColor(iterations uint64, iterlimit uint64, escapedPoint complex128) pixel.RGBA {
	if iterations == iterlimit {
		return pixel.RGBA{0, 0, 0, 255}
	}

	var s float64 //smooth coef

	s = float64(iterations) + 1.0 - math.Log(math.Log(math.Sqrt(real(escapedPoint)*real(escapedPoint)+imag(escapedPoint)*imag(escapedPoint))))/math.Log(2)

	s /= (float64(iterlimit))

	return pixel.RGBA{s * 190, 10, 10, 255}
}

func CameraMoveUp() {
	fractalBounds = fractalBounds.Moved(pixel.V(0, windowBounds.H()*moveSpeed))
}

func CameraMoveDown() {
	fractalBounds = fractalBounds.Moved(pixel.V(0, windowBounds.H()*moveSpeed*-1))
}

func CameraMoveRight() {
	fractalBounds = fractalBounds.Moved(pixel.V(windowBounds.H()*moveSpeed, 0))
}

func CameraMoveLeft() {
	fractalBounds = fractalBounds.Moved(pixel.V(windowBounds.H()*moveSpeed*-1, 0))
}

func CameraMove(v pixel.Vec) {
	v.X = utils.MapValueToRange(v.X, 0, windowBounds.W(), 0, fractalBounds.W())
	v.Y = utils.MapValueToRange(v.Y, 0, windowBounds.H(), 0, fractalBounds.H())

	fractalBounds = fractalBounds.Moved(v)
}

func CameraMoveCenter(v pixel.Vec) {
	v.X = utils.MapValueToRange(v.X, 0, windowBounds.W(), fractalBounds.Min.X, fractalBounds.Max.X)
	v.Y = utils.MapValueToRange(v.Y, 0, windowBounds.H(), fractalBounds.Min.Y, fractalBounds.Max.Y)

	fractalBounds = fractalBounds.Moved(v.Sub(fractalBounds.Center()))
}

// to jest do naprawy, zoom zawsze przyciaga się do punktu 0,0 Potrzebny debug
func CameraZoom(zoomCounts float64) {

	minX := fractalBounds.Min.X * math.Pow(camZoomSpeed, zoomCounts)
	minY := fractalBounds.Min.Y * math.Pow(camZoomSpeed, zoomCounts)
	maxX := fractalBounds.Max.X * math.Pow(camZoomSpeed, zoomCounts)
	maxY := fractalBounds.Max.Y * math.Pow(camZoomSpeed, zoomCounts)

	fmt.Println(fractalBounds)
	fractalBounds = pixel.R(minX, minY, maxX, maxY)
	fmt.Println(fractalBounds)
}

func IterationsUp() {
	iterationsLimit += iterationsJump
}

func IterationsDown() {
	iterationsLimit -= iterationsJump
}

func UpdateWinBounds(r pixel.Rect) {
	windowBounds = r
	Canvas.SetBounds(r)
}
