package main

func scale_value_to_range(x, a, b, c, d float64) float64 {
	return (x-a)*(d-c)/(b-a) + c
}
