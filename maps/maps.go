package main

import "fmt"

func main() {
	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1984"} // Creates a map
	b := map[string]int{"apple": 1, "banana": 2, "watermelon": 3, "orange": 4}     // Creates a map
	c := make(map[string]string)                                                   // Creates a map
	c["A"] = "1"
	c["B"] = "2"

	fmt.Printf("Map a: %v\n", a)
	fmt.Printf("Map b: %v \n", b)
	fmt.Printf("Map c: %v \n", c)

	d := make(map[string]string) // Creates an empty map
	fmt.Printf("Is map d empty =%v\n\n", len(d) == 0)

	/* Update elements of a map */
	a["color"] = "red"
	a["year"] = "1989"

	fmt.Printf("Updated map a: %v\n\n", a)

	/* Delete element from a map */
	delete(a, "year")
	fmt.Printf("Updated map a: %v\n\n", a)

	/* Check for specific element in map */
	val1, ok1 := a["brand"]
	val2, ok2 := a["year"]
	_, ok3 := a["color"]

	fmt.Println(val1, ok1)
	fmt.Println(val2, ok2)
	fmt.Println(ok3)

	/* Maps are references */
	var k = map[string]string{"brand": "Ford"}
	l := k

	fmt.Printf("Map k: %v\n", k)
	fmt.Printf("Map l: %v\n", l)

	l["brand"] = "GMC"
	fmt.Println("---After changing in l---")
	fmt.Printf("Map k: %v\n", k)
	fmt.Printf("Map l: %v\n", l)

	/* Iterate over maps */
	var numbers = map[string]int{"zero": 0, "one": 1, "two": 2, "three": 3}
	for k, v := range numbers {
		fmt.Println(k, v)
	}

	/* Iterate over maps in specific order */

	var num_order []string
	num_order = append(num_order, "zero", "one", "two", "three")

	for _, element := range num_order {
		fmt.Printf("%v : %v, ", element, numbers[element])
	}

}
