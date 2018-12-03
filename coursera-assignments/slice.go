package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {

	var input string

	// create an integer slice of length 0 and capacity 3
	slice := make([]int, 0, 3)

	// infinite loop
mainForLoop:
	for {
		fmt.Println("Please enter a number (or enter X to finish and exit): ")

		// read a new input
		_, err := fmt.Scan(&input)

		switch {

		// If error in reading input, then print the error and try reading again.
		case err != nil:
			fmt.Println("Error occurred : " + err.Error())

		// finish the program
		case strings.ToUpper(input) == "X":
			break mainForLoop

		// convert to number and add to slice and sort
		default:
			slice, err = addToSliceAndSort(slice, input)
		}

		if err == nil {
			fmt.Printf("Slice : %v \n", slice)
		}
	}
}

func addToSliceAndSort(slice []int, input string) ([]int, error) {
	// convert the input to a number
	num, err := strconv.Atoi(strings.TrimSpace(input))

	if err != nil {
		fmt.Println("Not a valid number. Please try again.")
		return slice, err
	}

	//append to slice
	slice = append(slice, num)

	//sort the slice
	sort.Ints(slice)

	return slice, nil
}
