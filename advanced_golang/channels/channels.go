package main

import "fmt"

func process(c chan int) {
	c <- 123
}

func process2(c chan int) {
	defer close(c)
	for i := 0; i < 5; i++ {
		c <- i
	}
	fmt.Println("Exiting process")
}

func main() {
	c := make(chan int)
	buffered_chan := make(chan int, 5)

	go process(c)
	fmt.Println(<-c)

	fmt.Println("--------*---------")
	go process2(c)
	for i := range c {
		fmt.Println(i)
	}

	fmt.Println("--------*---------")

	go process2(buffered_chan)
	for i := range buffered_chan {
		fmt.Println(i)
	}

}
