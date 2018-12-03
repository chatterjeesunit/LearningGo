package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	numbers := readInputNumbersIntoSlice()

	BubbleSort(numbers)

	fmt.Printf("Sorted Array of Numbers are : %v", numbers)
}

func readInputNumbersIntoSlice() []int {
	numbers := make([]int, 0, 10)

	fmt.Println("Please enter upto 10 integer numbers (type X to stop) :")

	for i := 0; i < 10; i ++ {
		var input string

		fmt.Scan(&input)

		if strings.ToUpper(input) == "X" {
			break;
		}

		number, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Invalid input. Exiting the program")
			os.Exit(1)
		}

		numbers = append(numbers, number)
	}

	return numbers
}

func BubbleSort(numbers []int) {

	size := len(numbers)

	for i:=0; i<size; i++ {

		for j:=0; j<size-i-1; j++ {

			if numbers[j] > numbers[j+1] {
				Swap(numbers, j)
			}
		}
	}
}

func Swap(numbers []int, currentIndex int) {

	temp := numbers[currentIndex]

	numbers[currentIndex] = numbers[currentIndex+1]
	numbers[currentIndex+1] = temp
}