package flow

import (
	"runtime"
	"time"
)

// Single case selected exclusively
// No break statements needed
func Switch() string {
	str := "Go running on "
	switch os := runtime.GOOS; os {
	case "darwin":
		return str + "macOS"
	case "linux":
		return str + "Linux"
	default:
		return str + os
	}
}

// Switch without a condition is equivalent to `switch true` pattern
func Greetings() string {
	t := time.Now().Hour()
	switch {
	case t < 12:
		return "Good Morning"
	case t < 17:
		return "Good Afternoon"
	default:
		return "Good Evening"
	}
}
