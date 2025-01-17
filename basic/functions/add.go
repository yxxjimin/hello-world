package functions

import (
	"fmt"
	"time"
)

// Called once on package import
func init() {
	fmt.Println("Init function for add")
}

// Non-exported names;
// cannot be used outside "functions" package
func currentTime() {
	fmt.Println("Current time:", time.Now())
}

func Add(x int, y int) int {
	currentTime()
	return x + y
}
