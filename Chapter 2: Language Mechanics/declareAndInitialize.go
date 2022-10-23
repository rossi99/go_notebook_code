package main

import "fmt"

func main() {
	// var can be used to construct values to their zero value for all types
	var a int
	var b string
	var c float64
	var d bool

	fmt.Printf("var a int \t %T [%v] \n", a, a)
	fmt.Printf("var b string \t %T [%v] \n", b, b)
	fmt.Printf("var c float64 \t %T [%v] \n", c, c)
	fmt.Printf("var d bool \t %T [%v] \n", d, d)

	// the shorthand ':=' can also be used to declare vars
	aa := 10
	bb := "hello"
	cc := 3.1234
	dd := true

	fmt.Printf("aa := 10 \t %T [%v] \n", aa, aa)
	fmt.Printf("bb := \"hello\" \t %T [%v] \n", bb, bb)
	fmt.Printf("cc := 3.1234 \t %T [%v] \n", cc, cc)
	fmt.Printf("dd := true \t %T [%v] \n", dd, dd)
}
