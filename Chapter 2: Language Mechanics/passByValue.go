package main

import "fmt"

// All data is moved around program by value. As data is passed around, each func & Goroutine gets its own copy of
// the data.

// Two types of data that I work with in Go: the value or the value's address. Addresses are data that need to be
// copied and stored program boundaries

// Code below outlines this movement:
func main() {
	// Declare var(int) with value of 10
	count := 10

	// Get address of value, using &
	fmt.Println("count:\t Value =", count, "\t Address =", &count)

	// Pass value of count to increment1 function
	increment1(count)

	// Print value and address of count var. This value will NOT have changed after increment function call.
	fmt.Println("count:\t Value =", count, "\t Address =", &count)

	// Pass copy of address of count to increment2 function. Still considered a pass by value and not by reference, as
	// addresses are values
	increment2(&count)

	// Print value and address of count. Value WILL have changed after increment function call
	fmt.Println("count:\t Value =", count, "\t Address =", &count)
}

// increment1 declares the func to accept its own copy of the int value
func increment1(inc int) {
	// Increments local copy of caller's int value.
	inc++
	fmt.Println("inc1:\t Value =", inc, "\t Address =", &inc)
}

// increment2 declares the func to accept its own copy of an address that points to an int value.
// Pointer vars are literal types and are declared using *
func increment2(inc *int) {
	// Increments the caller's int value through the pointer
	*inc++
	fmt.Println("inc2:\t Value =", inc, "\t Address =", &inc, "\t Points to =", *inc)
}
