package main

import (
	"fmt"
	"sync"
)

// The main go routines will now wait for other go routines to finish.. using sync.waitGroup

func main() {

	var wg sync.WaitGroup

	fmt.Println("Start of Program")

	wg.Add(1)
	go dummyPrintAgain(&wg)

	wg.Wait()

	fmt.Println("End of Program")

	
}

func dummyPrintAgain(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello World")
}