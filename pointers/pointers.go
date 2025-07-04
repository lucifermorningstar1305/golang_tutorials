package main

import "fmt"

type Person struct {
	name string
	age  int
}

func swap(a *int, b *int) {
	*a = *a ^ *b
	*b = *a ^ *b
	*a = *a ^ *b

}

func main() {
	a := 10
	var ip *int
	fmt.Println("Address of variable a:", &a)

	ip = &a // Pointer to variable a

	fmt.Println("Address stored in ip variable:", ip)
	fmt.Println("Value of *ip variable: ", *ip) // De-referencing the pointer

	fmt.Println("----------*--------------")
	b := 20
	ip2 := &b

	// Manipulating value by de-referencing.
	fmt.Println("Initial value of b:", b)
	fmt.Println("Memory address of b:", &b)
	fmt.Println("Memory address stored in ip2:", ip2)
	fmt.Println("Value pointed at by ip2:", *ip2)

	*ip2 = 30

	fmt.Println("Final value of b:", b)
	fmt.Println("Memory address of b:", &b)
	fmt.Println("Memory address stored in ip2:", ip2)

	// De-referencing struct pointers
	perpt := &Person{"John Doe", 132}
	fmt.Println("Person struct pointer:", perpt)

	fmt.Println("Name=", perpt.name)
	fmt.Println("Age=", (*perpt).age) // Can do this or `perpt.age`

	/* Array of pointers */
	arr := []int{10, 20, 30}

	const MAX int = 3

	var ptrArr [MAX]*int

	for i := 0; i < MAX; i++ {
		fmt.Printf("Value at arr[%d]=%v, address of arr[%d]=%v\n", i, arr[i], i, &arr[i])
	}

	fmt.Println()

	for i := 0; i < MAX; i++ {
		ptrArr[i] = &arr[i]
	}

	for i := 0; i < MAX; i++ {
		fmt.Printf("Value at ptrArr[%d]=%v, value pointed *ptrArr[%d]=%v\n", i, ptrArr[i], i, *ptrArr[i])
	}

	fmt.Println()

	/* Pointer to Pointer */
	x := 300
	ptr_x := &x
	ptr_ptr_x := &ptr_x // Pointer to Pointer

	fmt.Printf("Value of x=%v, Memory Address of x=%v\n", x, &x)
	fmt.Printf("Value contained in ptr_x=%v, Memory address of ptr_x=%v, Value pointed by ptr_x=%v\n", ptr_x, &ptr_x, *ptr_x)
	fmt.Printf("Value contained in ptr_ptr_x=%v, Value of ptr_ptr_x=%v, Value pointed by ptr_ptr_x=%v\n", ptr_ptr_x, *ptr_ptr_x, **ptr_ptr_x)

	fmt.Println()

	/* Passing pointers to function */
	val1, val2 := 5, 10
	fmt.Println("Before swapping")
	fmt.Printf("val1=%d, val2=%d\n\n", val1, val2)

	swap(&val1, &val2)
	fmt.Println("After swapping")
	fmt.Printf("val1=%d, val2=%d\n", val1, val2)
}
