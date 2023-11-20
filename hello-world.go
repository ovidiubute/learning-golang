package main

import (
	"fmt"
)

func main() {
	theArrays()
	theMaps()
	theRange()
}

func isOdd(n int) bool {
	if (n % 2) == 0 {
		return false
	}

	return true
}

func len(s string) int {
	fmt.Println("You can shadow built-in functions ?!")
	return len(s)
}

func theArrays() {
	var a [5]int
	a[0] = 3

	fmt.Println(a)

	var b = make([]int, 5)
	b[0] = 3
	fmt.Println(b)
}

func theRange() {
	var a = []int{1, 4, 5, 9}

	for i, x := range a {
		fmt.Println(i, x)
	}

	var m = make(map[string]string)
	m["a"] = "alpha"
	m["b"] = "beta"

	for k, v := range m {
		fmt.Println(k, v)
	}
}

func theMaps() {
	var m = make(map[string]int)

	m["a"] = 23
	m["b"] = 42
	fmt.Println(m)
}
