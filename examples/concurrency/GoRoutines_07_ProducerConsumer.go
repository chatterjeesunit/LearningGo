package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Producers and consumers using channels

var (
	items = make(chan int, 10)
	producerNotification = make(chan string, 3)
	consumerNotification = make(chan string, 1)
)

const (
	TOTAL_PRODUCERS = 10
	TOTAL_CONSUMERS = 5
)


func main() {


	for i:=0; i< TOTAL_PRODUCERS; i++ {

		producerName := fmt.Sprintf("P%d", i)

		go produce(producerName, 5)
	}

	for i:=0; i< TOTAL_CONSUMERS; i++ {

		consumerName := fmt.Sprintf("C%d", i)

		go consume(consumerName)
	}

	activeProducersLeft, activeConsumersLeft := TOTAL_PRODUCERS, TOTAL_CONSUMERS

	for {
		select {
		case msg := <-producerNotification:
			activeProducersLeft--
			fmt.Printf("Producer %s : Finished\n", msg)
			if activeProducersLeft == 0 {
				close(items)
			}

		case msg := <-consumerNotification:
			activeConsumersLeft--
			fmt.Printf("\t\t\t\t\t\t\t\t\tConsumer %s : Finished\n", msg)

		default:
			if activeConsumersLeft == 0 {
				return
			}
		}
	}

	fmt.Println("Exiting Program")

}

func produce(name string, n int) {
	fmt.Printf("Producer %s : Started\n", name)
	rand.Seed(time.Now().UnixNano())
	for i:=0; i<n; i++ {
		number := rand.Intn(500)
		items <- number
		fmt.Printf("%s : %d, [Queue Size = %d]\n", name, number, len(items))
		sleepForRandomTime()
	}

	producerNotification <- name
}

func consume(name string) {
	fmt.Printf("\t\t\t\t\t\t\t\t\tConsumer %s : Started\n", name)
	rand.Seed(time.Now().UnixNano())
	for item := range items {
		fmt.Printf("\t\t\t\t\t\t\t\t\t%s: %d, [Queue Size = %d]\n", name, item, len(items))
		sleepForRandomTime()
	}

	consumerNotification <- name
}

//Sleep for random time upto 10 ms
func sleepForRandomTime(){
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(10)))
}
