package main

import "fmt"

func main() {
	arr1 := [3]int{1, 2, 3}              // array defined with length
	arr2 := [...]int{10, 20, 30, 40, 50} // length of array is inferred
	prices := [...]int{30, 100, 350}

	fmt.Println("Array1: ", arr1)
	fmt.Println("Array2: ", arr2)

	fmt.Println("First element of arr1: ", arr1[0])
	fmt.Println("Second element of arr2: ", arr2[1])

	fmt.Println("Prices array: ", prices)
	prices[2] = 250 // Changing elemnent of an array
	fmt.Println("New Prices array: ", prices)

	/*Array Initialization*/
	arr3 := [5]int{}              // not initialized
	arr4 := [5]int{1, 2}          // partially initialized
	arr5 := [5]int{1, 2, 3, 4, 5} // fully initialized

	fmt.Println("Array3: ", arr3) // arrays which are not initialized takes the default value of their type. For int it is 0.
	fmt.Println("Array4: ", arr4)
	fmt.Println("Array5: ", arr5)

	/*Array initialize specific element*/
	arr6 := [5]int{1: 10, 2: 40} // sets 10 at index 1 and 40 at index 2
	fmt.Println("Array6: ", arr6)

	/*Length of an array*/
	fmt.Println("Length of array1: ", len(arr1))
	fmt.Println("Length of prices: ", len(prices))
	fmt.Println("Length of arr3: ", len(arr3))

}
