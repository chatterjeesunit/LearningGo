package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

func main() {
	setMaxProcs()

	maxRuns := 9999999

	var wg sync.WaitGroup

	output := make(chan bool)
	stop := make(chan bool)
	go finishCounter(output, stop)

	wg.Add(maxRuns)
	for l := 1; l <= maxRuns; l++ {
		go func() {
			defer wg.Done()
			rand.Seed(time.Now().UnixNano())
			factorial(uint64(10))
			output <- true
		}()
	}
	wg.Wait()
	stop <- true
	time.Sleep(60*time.Second)
}

func setMaxProcs() {
	fmt.Println("Max CPUs = ", runtime.NumCPU())
	maxProcs := runtime.NumCPU()
	if len(os.Args) >= 2 {
		maxProcs, _ = strconv.Atoi(os.Args[1])
	}
	runtime.GOMAXPROCS(maxProcs)
	fmt.Println("Max Procs set to ", maxProcs)
}

func factorial(n uint64) uint64 {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)

}

func finishCounter(out chan bool, stop chan bool) {
	counter := 0
	startTime := time.Now().UnixNano()
	for {
		select {
		case <- out:
			counter++
			if counter%100 == 0 {
				currentTime := time.Now().UnixNano()
				fmt.Printf("%d, \tTime = %d\n", counter, (currentTime - startTime)/(1000*1000*1000) )
			}
		case <- stop:
			fmt.Println("Program finished. Counter = ", counter)
			return
		default:

		}
	}
}
