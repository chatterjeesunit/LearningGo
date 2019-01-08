package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Race condition fixed using Mutex
*/

type SafeNum struct {
	sync.Mutex
	num int
}

func (safeNum *SafeNum) increment(delta int) {
	safeNum.Lock()
	defer safeNum.Unlock()

	time.Sleep(time.Microsecond * 1)
	safeNum.num = safeNum.num + delta
}

const (
	MAX = 100000
)

func main() {

	safeNumber := SafeNum{num: 100}

	var wg sync.WaitGroup

	wg.Add(2 * MAX)

	startTime := time.Now()
	for i := 0; i < MAX; i++ {

		go func() {
			safeNumber.increment(1)
			wg.Done()
		}()

		go func() {
			safeNumber.increment(-1)
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Total time taken : ", time.Since(startTime))

	fmt.Println("Expected value = 100, Actual Value = ", safeNumber.num)

}
