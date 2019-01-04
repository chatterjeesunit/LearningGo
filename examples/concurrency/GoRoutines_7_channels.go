package main

import (
	"fmt"
	"sync"
)

// using channels to sum first n numbers

func main() {
	n := 10

	channel := make(chan int)

	var wg sync.WaitGroup

	go channelProducer(n, channel)

	wg.Add(1)
	go channelConsumer(channel, n, &wg)

	wg.Wait()
}

func channelProducer(n int, channel chan int) {
	for i := 1; i <= n; i = i + 2 {
		if i == n {
			go add(i, 0, channel)
		} else {
			go add(i, i+1, channel)
		}
	}
}

func channelConsumer(channel chan int, n int, wg *sync.WaitGroup) {
	sum := 0
	counter := 0
	for i := range channel {
		//time.Sleep(time.Second * 5)

		sum += i
		//fmt.Printf("\tRecieved from channel : %d \n", i)
		counter += 2
		if counter >= n {
			close(channel)  //TODO: Wrong.. Receiver should not close the channel.
		}
	}
	fmt.Printf("Sum of first %d integers = %d \n", n, sum)
	wg.Done()

}

func add(x, y int, c chan int) {
	c <- x + y
	//fmt.Printf("Sent to channel : %d, %d\n", x, y)
}