package fractal_manager

import "math/big"

// Structure for arbitrary big vector
type BigVec struct {
	X, Y big.Float
}

// Structures for arbitrary big rectangle
type BigRect struct {
	Min, Max BigVec
}

// This function returns a new BigRect with given the Min and Max coordinates.
// Note that the returned rectangle is not automatically normalized.
func BigR(minX, minY, maxX, maxY big.Float) BigRect {
	return BigRect{
		Min: BigVec{minX, minY},
		Max: BigVec{maxX, maxY},
	}
}
