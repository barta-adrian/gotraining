package main

import "fmt"

func wrapper() func() int {
	x := 0
	fmt.Println("in wrapper", x)
	return func() int {
		fmt.Println("in closure", x)
		x++
		return x
	}
}

func main() {
	inter := wrapper()

	fmt.Println(inter())
	fmt.Println(inter())
	fmt.Println(inter())
}
