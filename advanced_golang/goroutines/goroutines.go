package main

import (
	"fmt"
	"sync"
	"time"
)

var m = sync.RWMutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func dbCall(i int) {
	delay := 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	// fmt.Println("The result from the database is: ", dbData[i])
	save(dbData[i])
	log()
	wg.Done()
}

func save(data string) {
	m.Lock()
	results = append(results, data)
	m.Unlock()
}

func log() {
	m.RLock()
	fmt.Println("\nResults:", results)
	m.RUnlock()
}

func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Println("\nTotal execution time:", time.Since(t0))
	fmt.Println("\nResults:", results)
}
