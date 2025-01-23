package containers

import (
	"fmt"
	"strings"
)

type vertex struct {
	lat, lng float64
}

// Keys cannot be added to empty nil maps
// Maps from `make` are also empty, but not nil and can add keys
func ZeroValueMap() {
	var m map[string]int
	fmt.Printf("Zero value of a map is nil: %t, %v\n", m == nil, m)

	mm := make(map[string]int)
	fmt.Printf("Default map generated from make is nil: %t, %v\n", mm == nil, mm)
}

// If top-level type is just a type name, literal element's type name
// can be omitted
func MapLiteral() {
	m := map[string]vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(m)
}

func MakeMapAndCheck() {
	m := make(map[string]vertex)
	m["Bell Labs"] = vertex{40.68433, -74.39967}
	m["Google"] = vertex{37.42202, -122.08408}
	fmt.Println(m)

	checkMapElem(m, "Google")
	checkMapElem(m, "Apple")
}

func checkMapElem(m map[string]vertex, key string) {
	if elem, ok := m[key]; ok {
		fmt.Printf("Element with key \"%s\" exists: %v\n", key, elem)
	} else {
		fmt.Printf("Element with key \"%s\" does not exist\n", key)
	}
}

func WordCount(s string) map[string]int {
	wordMap := make(map[string]int)
	for _, str := range strings.Fields(s) {
		wordMap[str]++
	}
	return wordMap
}
