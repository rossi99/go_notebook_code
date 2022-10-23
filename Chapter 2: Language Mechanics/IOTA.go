package main

import "fmt"

// IOTA <- setting successive int constants. It works with const block and starts with the value of 0 <- each
// successive constant declared increments by 1.
func iota1() {
	const (
		A1 = iota // 0 : Start at 0
		B1 = iota // 1 : Inc by 1
		C1 = iota // 2 : Inc by 2
	)
	fmt.Println(A1, B1, C1)
}

// The 'iota' keyword doesn't need to be repeated <- successive nature is assumed once applied
func iota2() {
	const (
		A1 = iota // 0 : Start at 0
		B1        // 1 : Inc by 1
		C1        // 1 : Inc by 1
	)
	fmt.Println(A1, B1, C1)
}

// If I need a different operation than increment by one, it can be added to the iota keyword
func iota3() {
	const (
		A1 = iota + 2 // 2 : 0 + 2
		B1            // 3 : 1 + 2
		C1            // 4 : 2 + 2
	)
	fmt.Println(A1, B1, C1)
}

// This can be used like the Log package sets flag <- bit operations can be applied with the increasing value of iota
// to calculate flag values
func iota4() {
	const (
		lDate         = 1 << iota // 1: shift 1 to the left 0 = 0000 0001
		lTime                     // 2: shift 1 to the left 1 = 0000 0010
		lMicroseconds             // 4: shift 1 to the left 2 = 0000 0100
		lLongFile                 // 8: shift 1 to the left 3 = 0000 1000
		lShortFile                // 16: shift 1 to the left 4 = 0001 0000
		lUTC                      // 32: shift 1 to the left 5 = 0010 0000
	)
	fmt.Println(lDate, lTime, lMicroseconds, lLongFile, lShortFile, lUTC)
}

func main() {
	iota1()
	iota2()
	iota3()
	iota4()
}
