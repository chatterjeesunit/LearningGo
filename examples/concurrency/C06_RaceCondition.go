package main

import (
	"fmt"
	"sync"
	"time"
)

// We have a initial number set to value 100

// Now we run a loop 10000 times, and call two go routines...
// One go routine increments the value
// Other go routine decrements the value

// This leads to a race condition since both increment and decrement go routine are trying to read and write to
// the same shared variable "num"

// Race condition occurs because every go routine tries to compete with each other and tries to complete first.
// Due to this the final value is NON Deterministic, as it will depend on the order in which the shared variable is
// updated by different go routins.

// After the program completes, the expected value is 100 (the original value),
// but everytime the actual value will be different

const (
	MAX = 100000
)

func main() {

	num := 100

	var wg sync.WaitGroup

	wg.Add(2 * MAX)

	startTime := time.Now()

	for i := 0; i < MAX; i++ {

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

func increment(number *int, delta int) {
	time.Sleep(time.Microsecond * 1)
	*number = *number + delta
}
