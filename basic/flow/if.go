package flow

import (
	"fmt"
	"math"
)

// `if` with a short statement
// Variable only available within the scope of `if` block
func Power(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("Answer %g is larger than %g\n", v, lim)
	}
	return lim
}
