package main

import (
	"fmt"
)

func helloWorld() {
	fmt.Println("Hello World")
}

func myName(firstName string, lastName string) {
	fullName := firstName + " " + lastName
	fmt.Printf("Your name is %v\n\n", fullName)
}

func addNumbers(num1 int, num2 int, message string) {
	sum := num1 + num2
	fmt.Printf("%v %d\n\n", message, sum)
}

func multiNumbers(num1 int, num2 int) int {
	return num1 * num2
}

/* Named return values */
func myFunc(x int, y int) (result int) {
	result = x*y - y
	return
}

func factorial(num float64) (result float64) { // Recursive function
	if num != 0 {
		result = num * factorial(num-1)
	} else {
		result = 1
	}
	return
}

func sumVariadic(nums ...int) int { // Variadic function
	total := 0
	fmt.Printf("nums: %v\n", nums)
	for _, val := range nums {
		total += val
	}

	return total
}

func intSeq() func() int { // Anonymous Function / Closure
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	helloWorld()
	myName("John", "Doe")
	addNumbers(2, 3, "2 + 3 =")
	mul := multiNumbers(4, 8)
	fmt.Printf("4 * 8 = %d\n\n", mul)

	var1 := myFunc(23, 56)
	fmt.Println(var1)

	fact_4 := factorial(4)
	fmt.Printf("4! = %.2f\n\n", fact_4)

	sumVariadic(1, 2)
	sumVariadic(3, 4)
	nums := []int{1, 2, 3, 4, 5}
	sumV := sumVariadic(nums...)

	fmt.Printf("sum of nums: %v\n\n", sumV)

	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
}
