package main

import "fmt"

func main() {
	// Go doesn't have casting, bytes need to be copied to new mem location for new representation.
	// Go has a package called 'unsafe' if a casting op needs to be preformed - this should never be used.
	aaa := int32(10)
	fmt.Printf("aaa := int32(10) \t %T [%v]", aaa, aaa)
}
