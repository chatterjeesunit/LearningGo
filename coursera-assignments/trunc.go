package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	var input string
	fmt.Println("Please enter a floating point number")

	_, err := fmt.Scan(&input)

	if err != nil {
		fmt.Printf("Error occured : %s. Exiting the program\n", err.Error)
		os.Exit(1)
	}

	indexOfDecimal := strings.Index(input, ".")

	switch indexOfDecimal {
	case -1:
		fmt.Println(input)
	default:
		fmt.Println(input[:indexOfDecimal])
	}
}
