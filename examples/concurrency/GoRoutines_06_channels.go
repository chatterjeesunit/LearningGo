package main

import (
	"fmt"
)

// using channels to sum first n numbers

func main() {
	n := 1000000

	channel := make(chan int)

	for i := 1; i <= n; i = i + 2 {
		go add(i, i+1, channel)
	}

	sum := 0
	for i := 0; i < n/2; i++ {
		//time.Sleep(time.Second * 5)

		sum += <-channel
		//fmt.Println("\tRecieved from channel")
	}

	fmt.Printf("\n\nSum of %d numbers = %d\n\n", n, sum)

}

func add(x, y int, c chan int) {
	c <- x + y
	//fmt.Println("Sent to channel")
}