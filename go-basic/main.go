package main

import (
	"fmt"
)

func main() {
	a := 2
	b := a
	a = 10
	fmt.Println(&a, &b)
	// access to memory address (&a is giving memory address)
	// *c is look the value of memory (pointer)

	c := &a
	a = 20
	fmt.Println(a, *c)

	*c = 50
	fmt.Println(a, *c)
}