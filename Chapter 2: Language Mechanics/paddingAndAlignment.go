package main

type memoryBytes struct {
	flag    bool
	counter int16
	pi      float32
}

type badMemoryBytes struct {
	flag    bool
	counter int16
	flag2   bool
	pi      float32
}

type goodMemoryBytes struct {
	pi      float32
	counter int16
	flag    bool
	flag2   bool
}

func main() {
	// - Working out the memory usage of each type
	var memBytes memoryBytes

	memBytes.flag = true // 1 byte
	memBytes.counter = 2 // 2 bytes
	memBytes.pi = 3.1415 // 4 bytes

	// That gives us 7 bytes, but in reality there are 8. There is a
	// padding byte sitting between the flag & counter fields for reasons
	// of alignment.

	// Alignment is important as it allows the hardware to read memory
	// more efficiently by placing it on specific boundaries.

	// Go determines the padding needed base on field and its
	// placement in the struct.

	// - Being efficient with memory
	// badMemoryBytes results in a total of 12 byte due to the additional
	// padding required bc of the order within the struct. Fields should be
	// laid out from the highest allocation of memory to the lowest allocation
	// of memory. goodMemoryBytes is an example of this.
}
