package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)


/*
Race condition fixed using Mutex
 */


const (
	MAX = 100000
)

var (
	num int64
)

func main() {

	atomic.StoreInt64(&num, 100)

	var wg sync.WaitGroup

	wg.Add(2 * MAX)

	startTime := time.Now()
	for i:=0; i<MAX; i++ {


		go func() {
			increment(&num, 1)
			wg.Done()
		}()


		go func() {
			increment(&num, -1)
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Total time taken : ", time.Since(startTime))

	fmt.Println("Expected value = 100, Actual Value = ", num)

}


func increment(number *int64, delta int64) {
	time.Sleep(time.Microsecond * 1)
	atomic.AddInt64(number, delta)
}



