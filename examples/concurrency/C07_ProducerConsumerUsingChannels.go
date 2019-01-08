package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Producer consumer problem

1. Multiple producers are producing some data (e.g some random numbers)
2. Multiple consumers are consuming this data (e.g reading one number at a time and printing it to console)

3. Producers are writing to some queue, and consumers are reading from the queue

4. Queue is of some fixed size

5. Producers should wait if queue is full (till consumers read some data from queue and free up space in queue)
6. Consumers should wait if queue is empty(till producers produce some more data)

There can be many variant of producer consumer problem
 - Fast Producer  / Slow consumer
 - Slow Producer / Fast consumer

In this example we will simulate 10 producers and 5 consumers.
Producer sleep for a random time after producing.
Consumer sleep for a random time after consuming.

Also producers will produce finite number of items and then quit.
Consumers will also quit, once they have consumed all the data from the producers
*/

var (
	items                = make(chan int, 10)
	producerNotification = make(chan string, 3)
	consumerNotification = make(chan string, 1)
)

const (
	TOTAL_PRODUCERS = 10
	TOTAL_CONSUMERS = 5
)

func main() {

	for i := 0; i < TOTAL_PRODUCERS; i++ {

		producerName := fmt.Sprintf("P%d", i)

		go produce(producerName, 5)
	}

	for i := 0; i < TOTAL_CONSUMERS; i++ {

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
	for i := 0; i < n; i++ {
		number := rand.Intn(500)
		items <- number
		fmt.Printf("%s : %d, [Queue Size = %d]\n", name, number, len(items))
		//fmt.Printf("%s : %d\n", name, number)
		sleepForRandomTime()
	}

	producerNotification <- name
}

func consume(name string) {
	fmt.Printf("\t\t\t\t\t\t\t\t\tConsumer %s : Started\n", name)
	rand.Seed(time.Now().UnixNano())
	for item := range items {
		fmt.Printf("\t\t\t\t\t\t\t\t\t%s: %d, [Queue Size = %d]\n", name, item, len(items))
		//fmt.Printf("\t\t\t\t\t\t\t\t\t%s: %d\n", name, item)
		sleepForRandomTime()
	}

	consumerNotification <- name
}

//Sleep for random time upto 10 ms
func sleepForRandomTime() {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(10)))
}
