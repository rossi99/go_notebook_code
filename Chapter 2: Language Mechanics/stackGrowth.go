package main

import "fmt"

// Size of each frame is calculated at compile time <- if compile doesn't know size, value is constructed on heap. An
// example of this is using the built-in func `make` to create a slice that's size is based on a var:
func main() {
	sliceSize := 1024
	b := make([]byte, sliceSize)

	fmt.Println(b)
}

// Go uses contiguous stack implementation (alternative to segmented stack implementation) <- determines stack growth
// & shrink.

// A func call has a preamble that asks if stack has enough room for frame. If yes, frame is taken & initialized. If
// no, a new stack is constructed & memory on existing stack is copied over <- requires changes in pointers that ref
// memory on stack. Benefits of contiguous memory & linear traversal is the trade-off for the cost of copy.

// Bc of contiguous stacks, no Goroutine can have a pointer to other Goroutine stack <- too much overhead for runtime
// to keep track of pointer to stack and readjust those pointers to new location.