package utils

import (
	"fmt"

	"github.com/faiface/pixel"
)

// Function to convert value `x` from range `ab` to range `cd`
func MapValueToRange(x, a, b, c, d float64) float64 {
	return (x-a)*(d-c)/(b-a) + c
}

// This function returns affine matrix which will stretch Picture `a` (in both axis) to fit Picture `b`
// and also centers `a` in `b`
func StretchToFit(a, b pixel.Picture) pixel.Matrix {

	return pixel.IM.
		ScaledXY(pixel.ZV, pixel.V(b.Bounds().W()/a.Bounds().W(), b.Bounds().H()/a.Bounds().H())).
		Moved(b.Bounds().Center())
}

// Scale rectangle B to have same proportions as rect A while keeping
// longer edge untouched and center of B in same place
func ScaleRectToRect(A pixel.Rect, B pixel.Rect) pixel.Rect {
	fmt.Println(A)
	fmt.Println(B)

	centerOfB := B.Center() // Save center of B for later

	A = A.Moved(A.Min.Scaled(-1)) // Move rectangle to (0,0) for easier calculations
	B = B.Moved(B.Min.Scaled(-1)) // Move rectangle to (0,0) for easier calculations

	if A.W() >= A.H() {
		// Width of A is bigger than height so we need to modifie Y in output rectangle
		B.Max.Y = B.W() * A.H() / A.W()

	} else {
		B.Max.X = B.H() * A.W() / A.H()
	}

	B = B.Moved(B.Center().Sub(centerOfB).Scaled(-1))

	fmt.Println(B)
	fmt.Println(A.W() / A.H())
	fmt.Println(B.W() / B.H())

	return B
}
