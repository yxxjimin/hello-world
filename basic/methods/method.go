package methods

import (
	"fmt"
	"math"
)

// Methods can only be declared for types defined in the same package
type MyFloat float64

// func (x int) Add(y int) int; cannot define new methods on non-local type int
func (x MyFloat) Add(y float64) float64 {
	return float64(x) + y
}

func (x MyFloat) Abs() float64 {
	if x < 0 {
		return float64(-x)
	}
	return float64(x)
}

type Vertex struct {
	X float64
	Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Value receiver; passed by (copied) value, has no effect
func (v Vertex) ScaleFake(f float64) {
	fmt.Printf("Called on %p\n", &v)
	v.X = v.X * f
	v.Y = v.Y * f
}

// Pointer receiver; modifies the value
func (v *Vertex) Scale(f float64) {
	fmt.Printf("Called on %p\n", v)
	v.X = v.X * f
	v.Y = v.Y * f
}
