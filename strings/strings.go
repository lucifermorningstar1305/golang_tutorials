package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {

	str := "Hello World" // Defines a string;

	fmt.Println("str=", str)

	// Concatenate two string
	fmt.Println("A" + "B")
	// Concatenate slice of strings
	greetings := []string{"Hello", "World"}
	fmt.Println(strings.Join(greetings, " "))

	/* String length */
	fmt.Println("length of str=Hello World is:", utf8.RuneCountInString(str))

	/*Trim spaces*/
	str2 := " Hello World "
	fmt.Printf("Initial str2=%v\nFinal str2=%v\n", str2, strings.TrimSpace(str2))

	/* Split string */
	str3 := "Hi there"
	fmt.Println(strings.Split(str3, " "))

}
