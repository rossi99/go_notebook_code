package main

import "fmt"

// User-defined type, composite of different fields & types.
type example struct {
	flag    bool
	counter int
	pi      float32
}

func main() {
	var el example

	// print with zero states
	fmt.Printf("%v \n", el)

	// set value using literal construction syntax
	el2 := example{
		flag:    true,
		counter: 1,
		pi:      3.14,
	}

	fmt.Println("Flag: ", el2.flag)
	fmt.Println("Counter: ", el2.counter)
	fmt.Println("Pi: ", el2.pi)
}
