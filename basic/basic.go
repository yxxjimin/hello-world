package main

import (
	"example/basic/containers"
	"example/basic/flow"
	"example/basic/functions"
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

}
