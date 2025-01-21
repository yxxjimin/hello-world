package types

import "fmt"

func ReadPointer() {
	i, j := 10, 9999
	p := &i
	fmt.Printf("*p = %d\n", *p)

	p = &j
	*p += 1 // changes made to `j`
	// p += 1 // pointer arithmetics are not available
	fmt.Printf("*p = %d, j = %d\n", *p, j)
}

// Struct fields can be accessed directly through pointers
// Use `p.Field` instead of `(*p).Field` (but both are allowed)
func ReadStructPointer(p *Vertex) {
	fmt.Printf(
		"PublicX = %d, PublicY = %d, privateName = %s\n",
		p.PublicX,
		p.PublicY,
		p.privateName,
	)
}
