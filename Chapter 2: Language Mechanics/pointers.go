package main

import "fmt"

// Pointers

// When a go program starts is greats a 'Goroutine'. These are lightweight app level threads with many of the same
// semantics as OS threads. They manage physical execution of instructions. Every program has at least 1 (the main
// goroutine).

// Each routine given its own block of memory: a stack. Each stack starts at 2048 bytes allocation - v small but
// can grow overtime.

// Every time a function is called, a block is added to the stack to aid the routine execute. These blocks are
// called a frame. The size of the frame is calculated at runtime.

// No value can be constructed on the stack unless the compiler knows the size of that value at compile time. If it
// doesn't know this then value has to be constructed on the heap.

// Stacks are self-cleaning & 0 value helps with initialization of the stack. Every time a function is called & a
// frame of memory is blocked out, that frame is initialized <- self-cleaning.

// On a func return, the memory block is left alone as it is unknown whether it will be needed again and it would
// be inefficient to initialized memory in returns.

func main() {
	fmt.Println("Pointer info")
}
