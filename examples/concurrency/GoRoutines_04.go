package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// The main go routines will now wait for other go routines to finish.. using sync.waitGroup


var (
	wg sync.WaitGroup
)


func main() {

	fmt.Println("Start of Program")

	wg.Add(2)
	go printMessage("Hello World.")
	go printMessage("Welcome to Go!")

	wg.Wait()

	fmt.Println("End of Program")

	
}

func printMessage(message string) {
	defer wg.Done()

	for i:=1; i <=5; i++{
		sleepTime := getRandomSleepTime()

		fmt.Printf("%d. %s \t(now sleeping for %v)\n",i, message, sleepTime)

		time.Sleep(sleepTime)
	}

}

func getRandomSleepTime() time.Duration {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Int63n(100)
	sleepTime := time.Duration(randomNumber) * time.Millisecond

	return sleepTime
}
