package main

import "fmt"

func main() {
	num1 := 20
	num2 := 30

	if num1 > num2 {
		fmt.Println("num1 is large")
	} else {
		fmt.Println("num2 is large")
	}

	slice1 := make([]int, 5, 8)
	slice2 := make([]int, 3, 10)

	if cap(slice2) > cap(slice1) {
		fmt.Printf("Capacity of slice2=%v, Capacity of slice1=%v\n\n", cap(slice2), cap(slice1))
	} else {
		fmt.Printf("Capacity of slice1=%v, Capacity of slice2=%v\n\n", cap(slice1), cap(slice2))
	}
}
