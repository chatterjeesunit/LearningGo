package main

import "fmt"

func main() {

	sum3Odd := sum(1, 3, 5)

	fmt.Println("Sum of 3 Odd numbers = ", sum3Odd)

	sum10 := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	fmt.Println("Sum of 10 numbers = ", sum10)


	numbers := []int {1, 2, 3, 4, 5}

	sum5Num := sum(numbers...)

	fmt.Println("Sum of 5 numbers = ", sum5Num)
	
}


func sum(numbers ...int) (sum int) {

	fmt.Printf("Data type of argument is %T\n", numbers)

	for num := range numbers {
		sum += num
	}

	return
}