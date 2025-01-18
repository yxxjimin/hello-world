package main

import (
	"example/basic/flow"
	"example/basic/functions"
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
}
