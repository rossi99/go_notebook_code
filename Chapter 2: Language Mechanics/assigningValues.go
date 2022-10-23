package main

import "fmt"

type example1 struct {
	word   string
	number int
}

type example2 struct {
	word   string
	number int
}

func main() {
	var ex1 example1
	var ex2 example2

	// ex1 cannot equal ex2 as they are of different types
	// ex1 = ex2 <- gives a compiler error

	// If we want to do this is must be set like:
	ex1 = example1(ex2) // <- a conversion syntax

	fmt.Println(ex1) // Removes err lines
}
