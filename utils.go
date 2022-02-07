package main

import "github.com/faiface/pixel"

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
