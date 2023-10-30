package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")

	fmt.Println("Go" + "Is Awesome")

	fmt.Println(1 + 2 + 3)

	fmt.Println(true && false)

	fmt.Println(true || false)

	fmt.Println(7 / 3.1)

	fmt.Print("Enter a number: ")

	fmt.Print(2222i)

	var a = 3
	fmt.Println(a)

	var b int
	fmt.Println(b)

	var c string
	fmt.Println(c)

	f := true
	fmt.Println(f)

	fmt.Println(int64(256))

	fmt.Println(len("Hello, World!"))
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
