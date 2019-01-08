package main

import (
	"fmt"
	"math/rand"
	"time"
)

// using channels to sum first n numbers

func main() {

	stringChannel1 := fetchMessage("Test")
	stringChannel2 := fetchMessage("Dummy")
	stringChannel3 := fetchMessage("Dummy")

	timeLimit := time.After(1 * time.Second)

	for {
		select {
		case msg := <-stringChannel1:
			fmt.Println(msg)
		case msg := <-stringChannel2:
			fmt.Println(msg)
		case msg := <-stringChannel3:
			fmt.Println(msg)
		case <-timeLimit:
			fmt.Println("Timeout achieved. Exiting the loop")
			return
		default:
			fmt.Println("waiting for message")
			sleepForRandomTime()
		}
	}
}

func fetchMessage(msgPrefix string) (stringChannel chan string) {
	stringChannel = make(chan string, 5)
	go func() {
		rand.Seed(time.Now().UnixNano())
		for {
			stringChannel <- fmt.Sprintf("%s-%06d", msgPrefix, rand.Intn(9999))
			sleepForRandomTime()
		}
	}()
	return
}

func sleepForRandomTime() {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(500)
	sleepTime := time.Duration(randomNumber) * time.Millisecond

	time.Sleep(sleepTime)
}
