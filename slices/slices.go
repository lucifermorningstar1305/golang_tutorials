package main

import "fmt"

func main() {
	arr1 := [6]int{1, 2, 3, 4, 5, 6}
	slice1 := []int{1, 2, 3} // Creates a slice
	slice2 := arr1[2:4]      // Creates a slice from an array

	fmt.Printf("Slice1=%v\n", slice1)
	fmt.Printf("Length of slice1=%d\n", len(slice1))
	fmt.Printf("Capacity of slice1=%d\n", cap(slice1))

	fmt.Printf("Slice2=%v\n", slice2)
	fmt.Printf("Length of slice2=%d\n", len(slice2))
	fmt.Printf("Capacity of slice2=%d\n", cap(slice2))

	slice3 := make([]int, 5, 10) // Creates a slice using the `make` function
	fmt.Printf("Slice3=%v\n", slice3)
	fmt.Printf("Length of slice3=%d\n", len(slice3))
	fmt.Printf("Capacity of slice3=%d\n", cap(slice3))

	prices := []int{10, 20, 30}
	fmt.Printf("Prices=%v\n", prices)
	fmt.Printf("First element of prices=%v\n", prices[0])            // Accessing elements of slice
	fmt.Printf("Last element of prices=%v\n", prices[len(prices)-1]) // Accessing the last element of slice

	prices = append(prices, 20) // Appends elements to the `prices` slice.
	fmt.Printf("New prices=%v\n", prices)

	prices2 := []int{100, 200}

	prices3 := append(prices, prices2...) // Appends one slice into another

	fmt.Printf("Prices3 = %v\n\n", prices3)

	fmt.Printf("Removing last 2 element of Prices3 = %v\n\n", prices3[:len(prices3)-2]) // Removes the last 2 elements of prices3

	/* Memory Efficiency using copy() */

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	fmt.Printf("numbers=%v\n", numbers)
	fmt.Printf("length of numbers=%v\n", len(numbers))
	fmt.Printf("capacity of numbers=%v\n\n", cap(numbers))

	neededNumbers := numbers[:len(numbers)-10]
	numbersCopy := make([]int, len(neededNumbers))
	copy(numbersCopy, neededNumbers)

	fmt.Printf("numbersCopy=%v\n", numbersCopy)
	fmt.Printf("length of numbersCopy=%v\n", len(numbersCopy))
	fmt.Printf("capacity of numbersCopy=%v\n", cap(numbersCopy))

	/* Multi-dimensional slice */
	multi_dim := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
	}

	fmt.Println(multi_dim)

}
