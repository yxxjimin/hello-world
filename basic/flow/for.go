package flow

import (
	"math"
)

// Basic for loop
func For() int {
	var sum int
	for i := 0; i < 10; i++ {
		sum += i
	}
	return sum
}

// For loop as a while loop
func Newton(x float64) float64 {
	z := 1.0
	for math.Abs(z-math.Sqrt(x)) > 0.0001 {
		z -= (z*z - x) / (2 * z)
	}
	return z
}
