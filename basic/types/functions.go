package types

import "fmt"

// Functions can be passed as values; not references
func Compute(fn func(float64, float64) float64, x, y float64) float64 {
	fmt.Printf("Using function %v with %v, %v\n", &fn, x, y)
	return fn(x, y)
}

// This function returns a closure `func(x int) int` that is bound to `sum`
func closure() func(int) int {
	sum := 0
	// this is a cloure
	return func(x int) int {
		sum += x
		return sum
	}
}

// Each closure is bound to its own `sum` variable
func TestClosure() {
	pos, neg := closure(), closure()
	for i := 0; i < 10; i++ {
		fmt.Printf("%d pos=%d neg=%d\n", i, pos(i), neg(-2*i))
	}
}

func fibonacciClosure() func() int {
	prev, curr := 0, 0
	next := 1
	return func() int {
		prev = curr
		curr = next
		next = prev + curr
		return prev
	}
}

func Fibonacci(x int) {
	f := fibonacciClosure()
	for i := 0; i < x; i++ {
		fmt.Printf("fibo %d: %d\n", i, f())
	}
}
