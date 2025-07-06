package main

import (
	"fmt"
	"math/rand"
	"time"
)

const MAX_PEN_PRICE = 3.0

func checkPenDeals(website string, penchan chan string) {
	for {
		time.Sleep(time.Second * 1)
		penPrice := rand.Float32() * 20
		if penPrice <= MAX_PEN_PRICE {
			penchan <- website
			break
		}
	}
}

func sendMessage(penchan chan string) {
	fmt.Println("Found deal at:", <-penchan)
}

func main() {
	penChannel := make(chan string)
	websites := []string{"aliexpress.com", "temu.com", "amazon.com"}
	for i := range websites {
		go checkPenDeals(websites[i], penChannel)
	}
	sendMessage(penChannel)
}
