package main

import "fmt"

type gasEngine struct {
	mpl    int
	litres int
}

type electricEngine struct {
	mpkwh int
	kwh   int
}

type Engine interface {
	milesLeft() int
}

/* Methods - These are functions that are directly tied to a struct */
func (e gasEngine) milesLeft() int {
	return e.mpl * e.litres
}

/* Methods - These are functions that are directly tied to a struct */
func (e electricEngine) milesLeft() int {
	return e.mpkwh * e.kwh
}

func canMakeIt(e Engine, miles int) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it there!")
	} else {
		fmt.Println("Need to fuel up first!")
	}
}

func main() {
	engine1 := gasEngine{mpl: 4, litres: 100}
	engine2 := electricEngine{mpkwh: 30, kwh: 60}

	fmt.Println("Engine1: ", engine1)
	fmt.Println("Engine2: ", engine2)

	fmt.Printf("Gas engine Miles left : %v\n", engine1.milesLeft())
	fmt.Printf("Electric engine Miles left : %v\n", engine2.milesLeft())

	canMakeIt(engine1, 500)
	canMakeIt(engine2, 50)
}
