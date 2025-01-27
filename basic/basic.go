package main

import (
	"example/basic/containers"
	"example/basic/flow"
	"example/basic/functions"
	"example/basic/methods"
	"example/basic/types"
	"fmt"
	"math"
)

func main() {
	fmt.Println(functions.Add(1, 2))
	fmt.Println(functions.Swap(3, 4))

	fmt.Println(flow.For())
	fmt.Println(flow.Newton(2))
	fmt.Println(flow.Power(2, 10, 1000))
	fmt.Println(flow.Switch())
	fmt.Println(flow.Greetings())

	flow.Defer()
	flow.StackDefer()

	types.ReadPointer()
	// myVertex := types.Vertex{1, 2}             // too few values in struct literal of type types.Vertex
	// myVertex := types.Vertex{1, 2, "myVertex"} // implicit assignment to unexported field privateName in struct literal of type types.Vertex
	myVertex := types.Vertex{
		PublicX: 1,
		PublicY: 2,
		// privateName: "myVertex", // cannot initialize here
	}
	types.ReadStructPointer(&myVertex)

	fn := math.Pow
	fmt.Printf("Original function is %v\n", &fn)
	types.Compute(fn, 3, 4)
	types.TestClosure()
	types.Fibonacci(10)

	containers.ArrayAndSlice()
	containers.SliceLiteral()
	containers.MakeAndAllocate(3, 3)
	containers.ZeroValueMap()
	containers.MapLiteral()
	containers.MakeMapAndCheck()
	fmt.Println(containers.WordCount("I ate a donut. Then I ate another donut."))

	var num methods.MyFloat = 3
	fmt.Println(num.Add(4))

	v := methods.Vertex{X: 3, Y: 4}
	fmt.Printf("Original value at: %p\n", &v)
	fmt.Println("Original Abs: ", v.Abs())
	v.ScaleFake(10)
	fmt.Println("Abs after ScaleFake: ", v.Abs())
	v.Scale(10) // Automatically interpreted as `(&v).Scale(10)`
	fmt.Println("Abs after Scale: ", v.Abs())

	p := &v
	p.Abs()
	v.Abs() // Equivalent
	p.Scale(1)
	v.Scale(1) // Equivalent

	// But not for interface type values
	var a methods.Abser = num // can hold any value that implements Abser methods
	fmt.Println(a.Abs())

	a = &v
	fmt.Println(a.Abs())

	// cannot use v (variable of type methods.Vertex) as methods.Abser value in assignment:
	// methods.Vertex does not implement methods.Abser (method Abs has pointer receiver)
	// a = v;

	t := methods.T{S: "Hello!"}
	var i methods.I = t
	i.M()

	// `pi` is of type *methods.I, not *methods.T, therefore `pi.M` is undefined
	// pi := &t
	// pi.M()
	pt := &t
	pt.M() // But this is OK; `pt` is of type *methods.T

	var nilI methods.I
	fmt.Printf("(%v, %T)\n", nilI, nilI)
	// nilI.M() // This will panic since there is no type (no concrete method)
}
