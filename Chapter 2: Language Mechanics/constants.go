package main

import "fmt"

// Go uniquely implements constants <- the rules for constants are specific to Go, provides the flexibility needed to
// make code readable & intuitive while maintaining type safety.

// Constants can be types/untyped. Untyped is considered to of a kind, these can be converted by the compiler <- happens
// at compile time.

// Untyped constants have a precision of 256 bits as stated by the spec. They are based on a kind.

// Typed constants still use constants type system, but their precision is restricted. (L:29)

// Constant arithmetic supports the use if different kinds of constants <- handled by Kind promotion. All happens
// implicitly. The 'answer' var is type float64 & represent 0.999 at a precision of 64 bits. (L:31)

// The 'third' constant will be of kind float and represent 1/3 at a precision of 256 bits. (L:32)

// The 'zero' constant will be of kind int and set to 0 since int division has no remainder. (L:36)

// L:40-41 is an example of constant arithmetic between types & untyped constants (int8 & set to 2). A constant of a
// type promotes over a constant of a kind.

// 'maxInt' is the maximum value for a 64-bit integer. 'bigger' is much larger than a 64-bit integer, but can be stored
// in a constant of kind int <- these aren't limited to 64 bits.

func main() {
	// UnTyped
	const ui = 12345    // type: int
	const uf = 3.141592 // type: floating-point

	// Typed
	const ti int = 12345        // type: int
	const tf float64 = 3.141592 // type: floating-point

	// const myUint8 uint8 = 1000 <- compiler error, 1000 overflows uint8s

	answer := 3 * 0.333   // float64 = kindFloat(3) * KindFloat(0.333)
	const third = 1 / 3.0 // float64 = kindFloat(1) * KindInt(3)

	const zero = 1 / 3 // KindInt = kindInt(1) / KindInt(3)

	const one int8 = 1
	const two = 2 * one // int8(2) * int8(1)

	const maxInt = 9223372036854775807
	const bigger = 9223372036854775808543522345

	const biggerErr int64 = 9223372036854775808543522345 // Compiler err: overflows int64

	// REMOVE ERR LINES
	fmt.Println(answer)
}
