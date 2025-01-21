package main

import (
	"example/basic/flow"
	"example/basic/functions"
	"example/basic/types"
	"fmt"
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
}
