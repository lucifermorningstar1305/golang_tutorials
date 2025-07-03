package main

import "fmt"

func zeroByVal(ival int) {
	ival = 0
}

func zeroByPtr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1

	fmt.Printf("Initial value of = %v\n", i)

	zeroByVal(i)
	fmt.Printf("zeroByVal = %v\n", i)

	zeroByPtr(&i)
	fmt.Printf("zeroByPtr = %v\n", i)

	fmt.Println("Pointer: ", &i)
}
