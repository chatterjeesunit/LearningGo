package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Why making the main function to sleep and wait for go routines to finish is a bad idea

func main() {

	fmt.Println("Start of Program")

	go printMessage("Hello World.")
	go printMessage("Welcome to Go!")

	time.Sleep(50 * time.Millisecond)

	fmt.Println("End of Program")

}

func printMessage(message string) {
	for i := 1; i <= 5; i++ {
		sleepTime := getRandomSleepTime()

		fmt.Printf("%d. %s \t(now sleeping for %v)\n", i, message, sleepTime)

		time.Sleep(sleepTime)
	}
}

func getRandomSleepTime() time.Duration {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Int63n(100)
	sleepTime := time.Duration(randomNumber) * time.Millisecond

	return sleepTime
}
