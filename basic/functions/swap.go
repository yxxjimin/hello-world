package functions

import "fmt"

func init() {
	fmt.Println("Init function for swap")
}

func Swap(x int, y int) (int, int) {
	currentTime()
	return y, x
}
