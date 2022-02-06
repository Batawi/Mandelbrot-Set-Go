package main

func mapValueToRange(x, a, b, c, d float64) float64 {
	return (x-a)*(d-c)/(b-a) + c
}
