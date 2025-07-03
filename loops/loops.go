package main

import "fmt"

func main() {

	/* General for loop */
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println()

	/*General nested loop*/
	for i := 0; i < 5; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Printf("%v \t", j+1)
		}
		fmt.Println()
	}

	fmt.Println()

	/* Continue statement & Break statement */

	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		} else if i == 4 {
			fmt.Println()
			break
		} else {
			fmt.Printf("%v\t", i)
		}

	}

	fmt.Println()

	/* Range keyword */

	fruits := [3]string{"apple", "banana", "orange"}
	for idx, val := range fruits {
		if idx == len(fruits)-1 {

			fmt.Printf("%d \t %v\n\n", idx, val)
		} else {
			fmt.Printf("%d \t %v\n", idx, val)
		}
	}

}
