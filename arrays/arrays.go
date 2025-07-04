package main

import "fmt"

func getAverage(arr [3]int) float64 {
	var average float64 = 0

	for _, x := range arr {
		average += float64(x) // type-casting
	}

	return average / float64(len(arr))
}

func getAverageMulti(arr [3][4]int, args ...int) []float64 {
	var dim int

	r, c := len(arr), len(arr[0])

	if len(args) == 0 {
		dim = 0
	} else {
		if len(args) > 1 {
			panic("Too many arguments")
		} else {
			if args[0] == -1 {
				dim = 1
			} else {
				dim = args[0]
			}
		}
	}

	var avgLen int
	if dim == 0 {
		avgLen = c
	} else {
		avgLen = r
	}
	average := make([]float64, avgLen)

	if dim == 0 {
		for j := 0; j < c; j++ {
			for i := 0; i < r; i++ {
				average[j] += float64(arr[i][j])
			}
		}
	} else {
		for i := 0; i < r; i++ {
			for j := 0; j < c; j++ {
				average[i] += float64(arr[i][j])
			}
		}
	}

	if dim == 0 {

		for i := 0; i < len(average); i++ {
			average[i] /= float64(r)
		}
	} else {
		for i := 0; i < len(average); i++ {
			average[i] /= float64(c)
		}
	}

	return average
}

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

	/* Multi-dimensional array */
	arr7 := [3][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			fmt.Printf("%v\t", arr7[i][j])
		}
		fmt.Println()
	}
	fmt.Println()

	fmt.Println("Value at arr7[1][2]=", arr7[1][2])

	avgPrice := getAverage(prices) // Passing array to a function
	fmt.Println("Average price=", avgPrice)

	fmt.Println("Multi-dimensional row-wise average:", getAverageMulti(arr7))
	fmt.Println("Multi-dimensional column-wise average:", getAverageMulti(arr7, -1))

}
