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

// Scale rectangle b to have same proportions as rect a while keeping
// longer edge untouched and center of b in same place
func ScaleRectToRect(a pixel.Rect, b pixel.Rect) pixel.Rect {
	centerOfB := b.Center() // Save center of b for later

	a = a.Moved(a.Min.Scaled(-1)) // Move rectangle to (0,0) for easier calculations
	b = b.Moved(b.Min.Scaled(-1)) // Move rectangle to (0,0) for easier calculations

	if a.W() >= a.H() {
		// Width of a is bigger than height so we need to modifie Y in output rectangle
		b.Max.Y = b.W() * a.H() / a.W()

	} else {
		b.Max.X = b.H() * a.W() / a.H()
	}

	return b.Moved(b.Center().Sub(centerOfB).Scaled(-1))
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

// Scale rectangle by given scalar, around its center
func ScaleRect(r pixel.Rect, scale float64) pixel.Rect {
	center := r.Center()          // Save center for later
	r = r.Moved(r.Min.Scaled(-1)) // Move rectangle to (0,0) for easier calculations
	r = pixel.R(0, 0, r.Max.X*scale, r.Max.Y*scale)

	return r.Moved(r.Center().Sub(center).Scaled(-1))
}
