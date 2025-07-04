package main

import (
	"fmt"
)

type Rectange struct {
	width, height float64
}

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

func intSeq() func() int {
	i := 0

	// Closure Creation
	return func() int {
		i++
		return i
	}
}

/* Function with struct type receiver */
func (rect Rectange) area() float64 {
	return rect.width * rect.height
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

	nextInt := intSeq() // Closure used
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println()

	rectObj := Rectange{width: 3.4, height: 2.5}
	fmt.Println("Area of the rectangle:", rectObj.area())

	/* Anonymous Function */
	add := func(a int, b int) int {
		return a + b
	}

	fmt.Println("3+5=", add(3, 5))

	/* Anonymous Function with closure */
	seqCounter := func() func() int {
		cntr := 0
		return func() int {
			cntr++
			return cntr
		}
	}

	s := seqCounter()
	fmt.Println(s())
	fmt.Println(s())

}
