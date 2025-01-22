package containers

import "fmt"

// Slices do not store any data; they just describe a section of an array.
// Changes made to a slice's element will modify corresponding element of its
// underlying array.
func ArrayAndSlice() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	a := names[0:2]
	b := names[1:3]

	b[0] = "XXX"

	fmt.Println(a, b)  // [John XXX] [XXX George]
	fmt.Println(names) // [John XXX George Ringo]
}

// Declaring a slice literal will first create an array, then build a slice
// that references it
func SliceLiteral() {
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}

	fmt.Println(s)
	fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))
}

// Create dynamically sized array through `make`
// `append` will reslice if capacity is sufficient. If not, will allocate a new
// array
func MakeAndAllocate(size, cap int) {
	s := make([]int, size, cap)
	printSlice(s)

	// cap -> 2*cap -> 4*cap -> 8*cap ->...
	for i := 0; i < 10; i++ {
		s = append(s, i+1)
		printSlice(s)
	}
}

func printSlice(s []int) {
	// cap is the length of underlying array
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
