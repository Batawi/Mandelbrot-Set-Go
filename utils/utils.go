package utils

import (
	"math"

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
	return B
}

// Chop given rectangle into uniform horizontal slices
// If last slice is requested than given block is increased by remainder of the division
// noSlices - number of Slices
// sliceNumber - requested slice number indexed from 0
func ChopHor(r pixel.Rect, noSlices, sliceNumber int32) pixel.Rect {

	res := math.Mod(r.H(), float64(noSlices))
	rH := r.H() - res

	minY := float64(sliceNumber)*rH/float64(noSlices) + r.Min.Y
	maxY := minY + rH/float64(noSlices)

	// Check if last slice is requested
	if sliceNumber+1 == noSlices {
		maxY += res
	}

	return pixel.R(r.Min.X, float64(minY), r.Max.X, float64(maxY))
}
