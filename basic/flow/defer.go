package flow

import "fmt"

func eval() string {
	fmt.Println("internal evaluation logic...") // first
	return "world"
}

// Immediate evaluation, deferred execution
func Defer() {
	defer fmt.Println(eval()) // third
	fmt.Println("hello")      // second
}

// Deferred statements are stacked in a LIFO manner
func StackDefer() {
	fmt.Println("counting...")
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done!")
}
