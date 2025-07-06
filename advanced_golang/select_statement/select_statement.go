package main

import (
	"fmt"
	"math/rand"
	"time"
)

const MAX_PEN_PRICE = 3.0
const MAX_NOTEBOOK_PRICE = 8.0

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

func checkNotebookPrice(website string, nbchan chan string) {
	for {
		time.Sleep(time.Second * 1)
		nbPrice := rand.Float32() * 20
		if nbPrice <= MAX_PEN_PRICE {
			nbchan <- website
			break
		}
	}
}

func sendMessage(penchan chan string, nbchan chan string) {
	select {
	case website := <-penchan:
		fmt.Printf("Text Sent: Found Deal at %s\n", website)
	case website := <-nbchan:
		fmt.Printf("Email Sent: Found Deal at %s\n", website)
	}
}

func main() {
	penChannel := make(chan string)
	nbChannel := make(chan string)
	websites := []string{"aliexpress.com", "temu.com", "amazon.com"}
	for i := range websites {
		go checkPenDeals(websites[i], penChannel)
		go checkNotebookPrice(websites[i], nbChannel)
	}
	sendMessage(penChannel, nbChannel)
}
