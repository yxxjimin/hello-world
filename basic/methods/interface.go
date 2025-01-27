package methods

import "fmt"

type Abser interface {
	Abs() float64
}

type I interface {
	M()
}

type T struct {
	S string
}

// type T implements interface I implicitly
func (t T) M() {
	fmt.Println(t.S)
}

// type *T implements interface I
// T and *T cannot be implemented at the same time
// func (t *T) M() {
// 	fmt.Println(t.S)
// }
