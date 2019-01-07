package main

import (
	"fmt"
	"math/rand"
	"time"
)

// another example of synchronization using channels

func main() {

	c := make(chan string)

	fmt.Println("Start of Program")

	go printMessage("Hello World." , c)
	go printMessage("Welcome to Go!" , c)


	for i:=0; i<2; i++ {
		msg := <- c
		fmt.Println(msg)
	}

	fmt.Println("End of Program")

	
}

func printMessage(message string, channel chan string) {
	for i:=1; i <=5; i++{
		sleepTime := getRandomSleepTime()

		fmt.Printf("%d. %s \t(now sleeping for %v)\n",i, message, sleepTime)

		time.Sleep(sleepTime)
	}

	channel <- "Completed " + message + "!!!"
}

func getRandomSleepTime() time.Duration {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Int63n(100)
	sleepTime := time.Duration(randomNumber) * time.Millisecond

	return sleepTime
}
