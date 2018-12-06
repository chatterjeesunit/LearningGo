package main

import (
	"fmt"
	"sync"
	"time"
)


// The program will generate nth fibonnaci numbers (with and without go routines)
// and we will demonstrate how race conditions occurs with go routines
// Example of Fib Numbers are - 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765,
// 10946, 17711, 28657, 46368, 75025, 121393, 196418, 317811


//So for e.g
// 3rd Fib Number = 1
// 25th Fib Number = 46368


// We will write two functions
// - one function to generate 25th fibonacci numbers using Go routines and race conditions
// - other function which does not uses go routines and has no race condition

// Run the program 10-20 times.

// The function without race condition and go routines will always print the result as 46368

// The function with race condition will print the result different every time


// Reasons for this race condition

// There are two shared variables number1 and number two, that different go routines are trying to modify.
// Different go routines are reading and writing to the same shared variables and hence the results cannot be deterministic
// Results will depend on which thread updated which value in which order.

//

func main() {

	generateNthFibonacciNumberWithoutRaceCondition(25)
	generateNthFibonacciNumberWithRaceCondition(25)

}

func generateNthFibonacciNumberWithRaceCondition(n int) {
	number1 := 0
	number2 := 1
	var wg sync.WaitGroup
	for i := 3; i <= n; i++ {
		wg.Add(1)
		go func() {
			generateNextSequenceOfNumbers(&number1, &number2)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("%dth Fibonacci Number (with Race condition)  = %d \n",n, number2)
}


func generateNthFibonacciNumberWithoutRaceCondition(n int) {
	number1 := 0
	number2 := 1
	for i := 3; i <= n; i++ {
		func() {
			generateNextSequenceOfNumbers(&number1, &number2)
		}()
	}
	fmt.Printf("%dth Fibonacci Number (without Race condition) = %d \n",n, number2)
}


// This functions generates the next sequence of fib numbers
func generateNextSequenceOfNumbers(number1 *int, number2 *int) {
	time.Sleep(time.Millisecond * 10)
	temp := *number2
	*number2 = *number2 + *number1
	*number1 = temp

}



