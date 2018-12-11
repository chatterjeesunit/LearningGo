package main

import (
	"fmt"
	sort2 "sort"
	"strconv"
)

/*
Sample Input : 1 2 3 4 5 5 2 1 6 7 3 2 77 2 55 1 54 21 77 11 0 -1 5 -4 -45 23 x
 */
func main() {


	numbers := readInput()
	length := len(numbers)

	n := 4
	numChannel := make(chan []int, n)

	sortSubArrays(length, n, numbers, numChannel)

	for {

		array1 := <- numChannel

		if n > 1 {
			array2 := <- numChannel
			mergeTwoSortedArrays(array1, array2, numChannel)

			n = n - 1
		}else {
			fmt.Printf("\n\nSorted Array = %v \n", array1)
			break;
		}
	}
}

func mergeTwoSortedArrays(a1 []int, a2 []int, numChannel chan []int) {

	//fmt.Printf("\tMerging arrays : %v , %v\n", a1, a2)
	lengthA1 := len(a1)
	lengthA2 := len(a2)
	sortedArray := make([]int, 0, lengthA1 + lengthA2)

	i, j := 0, 0
	for i < lengthA1 && j < lengthA2 {

		if a1[i] < a2[j] {
			sortedArray = append(sortedArray, a1[i])
			i++
		} else {
			sortedArray = append(sortedArray, a2[j])
			j++
		}
	}

	if i < lengthA1 {
		sortedArray = append(sortedArray, a1[i:lengthA1]...)
	}else if j < lengthA2 {
		sortedArray = append(sortedArray, a2[j:lengthA2]...)
	}


	numChannel <- sortedArray
}


func sortSubArrays(length int, n int, numbers []int, numChannel chan []int) {
	maxSize := length/n
	if length % n > 0 {
		maxSize++
	}

	for i := 0; i < length; i = i + maxSize {

		end := i + maxSize
		if end > length {
			end = length
		}

		go sort(numbers[i:end], numChannel)
	}
}

func readInput() []int {
	fmt.Println("Sorting Demo")
	fmt.Println("Please enter numbers separated by space (enter any other character to exit) :")
	numbers := make([]int, 0)
	for {
		var input string
		fmt.Scan(&input)

		num, err := strconv.Atoi(input)

		if err != nil {
			break;
		}

		numbers = append(numbers, num)
	}
	return numbers
}


func sort(numArray []int, numChannel chan []int) {

	fmt.Printf("\tGo Routine: Sorting Sub-Array : %v \n", numArray)

	sort2.Ints(numArray)

	numChannel <- numArray
}