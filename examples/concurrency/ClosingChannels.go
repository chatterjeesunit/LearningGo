package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	datachannel1 := make(chan int)
	datachannel2 := make(chan int)
	datachannel3 := make(chan int)

	stopChan := make(chan struct{}, 1)

	wg.Add(2)
	go test("A", datachannel1, stopChan,  &wg)
	go test("B", datachannel1, stopChan,  &wg)
	go test("C", datachannel1, stopChan,  &wg)

	for i:=0; i < 5 ; i++ {
		datachannel1 <- i
		datachannel2 <- i
		datachannel3 <- i
	}
	close(stopChan)
	wg.Wait()

}


func test(id string, datachan chan int, stopchan chan struct{}, wg *sync.WaitGroup) {
	defer func() {
		fmt.Println("Closing : ", id)
		wg.Done()
	}()

	fmt.Println("Testing ", id)
	for {
		select {
			case <- stopchan:
				fmt.Println("Stoppping : ", id)
				return
			case data := <- datachan:
				fmt.Println(id , ": ", data)
		}

	}

}