package main

import (
	"fmt"
)

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func printDetails(pers Person) {
	fmt.Printf("Name: %v\n", pers.name)
	fmt.Printf("Age: %v\n", pers.age)
	fmt.Printf("Job: %v\n", pers.job)
	fmt.Printf("Salary: %v\n\n", pers.salary)
}

func main() {
	var person1 Person
	var person2 Person

	person1.name = "John Doe"
	person1.age = 42
	person1.job = "Sales Executive"
	person1.salary = 240000

	person2.name = "Jane Doe"
	person2.age = 32
	person2.job = "Marketing Executive"
	person2.salary = 320000

	fmt.Println("Details of person1")
	fmt.Printf("Name: %v\n", person1.name)
	fmt.Printf("Age: %v\n", person1.age)
	fmt.Printf("Job: %v\n", person1.job)
	fmt.Printf("Salary: %v\n\n", person1.salary)

	fmt.Println("Details of person2")
	fmt.Printf("Name: %v\n", person2.name)
	fmt.Printf("Age: %v\n", person2.age)
	fmt.Printf("Job: %v\n", person2.job)
	fmt.Printf("Salary: %v\n\n", person2.salary)

	printDetails(person1)
	printDetails(person2)

}
